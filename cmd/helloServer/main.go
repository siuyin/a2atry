package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/a2aproject/a2a-go/a2a"
	"github.com/siuyin/dflt"
)

// TaskHandler is a function type that handles task processing
type TaskHandler func(task *a2a.Task, message *a2a.Message) (*a2a.Task, error)

// A2AServer represents an A2A server instance
type A2AServer struct {
	agentCard   a2a.AgentCard
	handler     TaskHandler
	port        string
	basePath    string
	taskStore   map[string]*a2a.Task
	taskHistory map[string][]*a2a.Message
	mu          sync.RWMutex
}

// Start starts the A2A server
func (s *A2AServer) Start() error {
	mux := http.NewServeMux()
	mux.Handle(s.basePath, s)
	return http.ListenAndServe(":"+s.port, mux)
}

// ServeHTTP implements the http.Handler interface
func (s *A2AServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("%s: Method not allowed", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	port := dflt.EnvString("PORT", "8080")
	log.Printf("PORT=%s", port)

	yes := true
	no := false
	svr := A2AServer{
		agentCard: a2a.AgentCard{
			Name:         "Time Agent",
			Description:  "An agent that can tell you UTC time",
			URL:          "http://localhost:" + port,
			Version:      "1.0.0",
			Capabilities: a2a.AgentCapabilities{Streaming: &yes, PushNotifications: &no, StateTransitionHistory: &yes},
			Skills: []a2a.AgentSkill{
				a2a.AgentSkill{ID: "utctime-skill", Name: "utctime-skill",
					Description: "a skill that gets the current time in UTC"},
			},
		},
		port:     port,
		basePath: "/",
		handler:  utcTime,
	}

	log.Fatal(svr.Start())
}

func utcTime(task *a2a.Task, msg *a2a.Message) (*a2a.Task, error) {
	log.Println("utctime called")
	return &a2a.Task{}, nil
}
