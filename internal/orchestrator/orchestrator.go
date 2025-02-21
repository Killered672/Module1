package orchestrator

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"calc_service/internal/models"
)

type Orchestrator struct {
	expressions map[string]*models.Expression
	tasks       map[string]*models.Task
	mu          sync.Mutex
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{
		expressions: make(map[string]*models.Expression),
		tasks:       make(map[string]*models.Task),
	}
}

func (o *Orchestrator) AddExpression(expr string) (string, error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	id := generateID()
	o.expressions[id] = &models.Expression{
		ID:     id,
		Status: "pending",
	}

	tasks, err := parseExpressionToTasks(id, expr)
	if err != nil {
		return "", err
	}

	for _, task := range tasks {
		o.tasks[task.ID] = task
	}

	return id, nil
}

func (o *Orchestrator) GetExpression(id string) (*models.Expression, error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	expr, exists := o.expressions[id]
	if !exists {
		return nil, errors.New("expression not found")
	}

	return expr, nil
}

func (o *Orchestrator) GetAllExpressions() []models.Expression {
	o.mu.Lock()
	defer o.mu.Unlock()

	var expressions []models.Expression
	for _, expr := range o.expressions {
		expressions = append(expressions, *expr)
	}

	return expressions
}

func (o *Orchestrator) GetTask() (*models.Task, error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	for _, task := range o.tasks {
		if task.Status == "pending" {
			task.Status = "processing"
			return task, nil
		}
	}

	return nil, errors.New("no tasks available")
}

func (o *Orchestrator) SubmitTaskResult(taskID string, result float64) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	task, exists := o.tasks[taskID]
	if !exists {
		return errors.New("task not found")
	}

	task.Result = result
	task.Status = "completed"

	expr := o.expressions[task.ExpressionID]
	expr.Status = "completed"
	expr.Result = fmt.Sprintf("%g", result)

	return nil
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func parseExpressionToTasks(expressionID, expr string) ([]*models.Task, error) {
	return []*models.Task{}, nil
}
