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
		log.Printf("Started worker %d", i+1)
	}
}

func (a *Agent) worker() {
	for task := range a.orchestrator.GetTask() {
		log.Printf("Processing task %s: %f %s %f", task.ID, task.Arg1, task.Operation, task.Arg2)
		result := a.executeTask(task)
		log.Printf("Task %s result: %f", task.ID, result)

		a.orchestrator.SubmitTaskResult(&models.TaskResult{
			ID:     task.ID,
			Result: result,
		})
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
		if task.Arg2 == 0 {
			log.Println("Division by zero detected")
			return 0
		}
		return task.Arg1 / task.Arg2
	default:
		log.Println("Invalid operation")
		return 0
	}
}
