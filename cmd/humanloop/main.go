package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"trpc.group/trpc-go/trpc-a2a-go/log"

	"github.com/siuyin/a2atry/msg"
	"github.com/siuyin/a2atry/ptr"
	"github.com/siuyin/dflt"
	spec "trpc.group/trpc-go/trpc-a2a-go/protocol"
	"trpc.group/trpc-go/trpc-a2a-go/server"
	tm "trpc.group/trpc-go/trpc-a2a-go/taskmanager"
)

// Human in the loop interactions are tracked with multiTurnSession.
type multiTurnSession struct {
	stage    int
	text     string
	mode     string
	complete bool
}

// humanLoopAgent satisfies the tm.MessageProcessor interface.
type humanLoopAgent struct {
	sess map[string]multiTurnSession // keyed by contextID
}

func (h *humanLoopAgent) ProcessMessage(ctx context.Context, m spec.Message, opts tm.ProcessOptions, handler tm.TaskHandler) (*tm.MessageProcessingResult, error) {
	h.init()

	txt := msg.Text(m)
	log.Info("received input: ", txt)

	cID := handler.GetContextID()
	if cID == "" {
		resp := compose("hello world response: no context ID")
		return &tm.MessageProcessingResult{Result: resp}, nil
	}

	resp := compose("second hello world response: with context ID: " + cID)
	return &tm.MessageProcessingResult{Result: resp}, nil
}

func (h *humanLoopAgent) init() {
	if h.sess == nil {
		h.sess = make(map[string]multiTurnSession)
	}
}

func compose(s string) *spec.Message {
	resp := spec.NewMessage(
		spec.MessageRoleAgent,
		[]spec.Part{spec.NewTextPart(s)},
	)
	log.Info("sending output: ", s)
	return &resp
}

func main() {
	port := dflt.EnvString("PORT", "8080")
	log.Infof("PORT=%s", port)
	log.Infof("curl http://localhost:%s/.well-known/agent.json for agent card", port)

	svr, err := server.NewA2AServer(myAgentCard(port), myTaskManager(&humanLoopAgent{}))
	if err != nil {
		log.Fatal("new server:", err)
	}

	log.Fatal(svr.Start(":" + port))
}

func myAgentCard(port string) server.AgentCard {
	return server.AgentCard{
		Name:        "humanLoopAgent",
		Description: "An A2A agent that incorporates human decision making in the loop.",
		URL:         fmt.Sprintf("http://localhost:%s/", port),
		Version:     "1.0.0",
		Capabilities: server.AgentCapabilities{
			Streaming:              ptr.Bool(true),
			PushNotifications:      ptr.Bool(true),
			StateTransitionHistory: ptr.Bool(true),
		},
		Skills: []server.AgentSkill{{
			ID:          "humanloop",
			Name:        "Have a human in the loop make the final decision.",
			Description: ptr.String("Have a human in the loop make the final decision."),
		}},
	}
}

func myTaskManager(mp tm.MessageProcessor) tm.TaskManager {
	mgr, err := tm.NewMemoryTaskManager(mp)
	if err != nil {
		log.Fatal("new task manager: ", err)
	}

	return mgr
}

func timeFor(loc string) string {
	tz, ok := tzOf(loc)
	if !ok {
		return "unsupported location"
	}
	l, err := time.LoadLocation(tz)
	if err != nil {
		log.Error("timeFor: ", err, tz)
		return ""
	}
	return time.Now().In(l).Format("15:04:05.000")
}

func tzOf(loc string) (string, bool) {
	tz := make(map[string]string)
	tz["singapore"] = "Asia/Singapore"
	tz["sgp"] = "Asia/Singapore"
	tz["new york"] = "America/New_York"
	tz["los angeles"] = "America/Los_Angeles"

	keys := keysOfMap(tz)
	l, ok := supportedLocation(loc, keys)
	if !ok {
		return "UTC", false
	}
	return tz[l], true
}

func keysOfMap[T any](m map[string]T) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
func supportedLocation(loc string, db []string) (string, bool) {
	loc = strings.ToLower(loc)
	for _, v := range db {
		if strings.Contains(loc, v) {
			return v, true
		}
	}
	return "", false
}
