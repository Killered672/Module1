package agent

import (
	"log"
	"time"

	"calc_service/internal/models"
	"calc_service/internal/orchestrator"
)

type Agent struct {
	orchestrator *orchestrator.Orchestrator
	power        int
}

func NewAgent(orchestrator *orchestrator.Orchestrator, power int) *Agent {
	return &Agent{
		orchestrator: orchestrator,
		power:        power,
	}
}

func (a *Agent) Start() {
	for i := 0; i < a.power; i++ {
		go a.worker()
	}
}

func (a *Agent) worker() {
	for {
		task, err := a.orchestrator.GetTask()
		if err != nil {
			log.Println("No tasks available, waiting...")
			time.Sleep(1 * time.Second)
			continue
		}

		result := a.executeTask(task)
		if err := a.orchestrator.SubmitTaskResult(task.ID, result); err != nil {
			log.Printf("Failed to submit task result: %v", err)
		}
	}
}

func (a *Agent) executeTask(task *models.Task) float64 {
	time.Sleep(time.Duration(task.OperationTime) * time.Millisecond)

	switch task.Operation {
	case "+":
		return task.Arg1 + task.Arg2
	case "-":
		return task.Arg1 - task.Arg2
	case "*":
		return task.Arg1 * task.Arg2
	case "/":
		return task.Arg1 / task.Arg2
	default:
		return 0
	}
}
