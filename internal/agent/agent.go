package agent

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"calc_service/internal/evaluator"
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
	for {
		pendingTasks := a.orchestrator.GetPendingTasks()

		if len(pendingTasks) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		task := pendingTasks[0]
		log.Printf("Processing task %s: %f %s %f", task.ID, task.Arg1, task.Operation, task.Arg2)

		result := a.executeTask(task)
		log.Printf("Task %s result: %f", task.ID, result)

		a.orchestrator.UpdateTaskResult(task.ID, result)
	}
}

func (a *Agent) executeTask(task *models.Task) float64 {
	expression := fmt.Sprintf("%f %s %f", task.Arg1, task.Operation, task.Arg2)

	resultStr, err := evaluator.EvaluateExpression(expression)
	if err != nil {
		log.Printf("Failed to evaluate task %s: %v", task.ID, err)
		return 0
	}

	result, err := strconv.ParseFloat(resultStr, 64)
	if err != nil {
		log.Printf("Failed to parse result for task %s: %v", task.ID, err)
		return 0
	}

	return result
}
