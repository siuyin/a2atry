package main

import (
	"context"
	"fmt"
	"log"

	"github.com/siuyin/a2atry/ptr"
	"github.com/siuyin/dflt"
	spec "trpc.group/trpc-go/trpc-a2a-go/protocol"
	"trpc.group/trpc-go/trpc-a2a-go/server"
	tm "trpc.group/trpc-go/trpc-a2a-go/taskmanager"
)

type timeAgent struct{}

func (t *timeAgent) ProcessMessage(ctx context.Context, msg spec.Message, opts tm.ProcessOptions, handler tm.TaskHandler) (*tm.MessageProcessingResult, error) {
	return &tm.MessageProcessingResult{}, nil
}

func main() {
	port := dflt.EnvString("PORT", "8080")
	log.Printf("PORT=%s", port)

	svr, err := server.NewA2AServer(myAgentCard(port), myTaskManager(&timeAgent{}))
	if err != nil {
		log.Fatal("new server:", err)
	}
	log.Fatal(svr.Start(":" + port))
}

func myAgentCard(port string) server.AgentCard {
	return server.AgentCard{
		Name:        "timeAgent",
		Description: "An A2A agent that tells the current time.",
		URL:         fmt.Sprintf("http://localhost:%s/", port),
		Version:     "1.0.0",
		Capabilities: server.AgentCapabilities{
			Streaming:              ptr.Bool(true),
			PushNotifications:      ptr.Bool(false),
			StateTransitionHistory: ptr.Bool(true),
		},
		Skills: []server.AgentSkill{{
			ID:          "tell_time",
			Name:        "Tells the time",
			Description: ptr.String("Tells the time."),
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
