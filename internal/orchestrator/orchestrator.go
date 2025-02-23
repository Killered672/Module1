package orchestrator

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"calc_service/internal/models"
)

type Orchestrator struct {
	expressions  map[string]*models.Expression
	tasks        map[string]*models.Task
	pendingTasks chan *models.Task
	resultChan   chan *models.TaskResult
	mu           sync.Mutex
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{
		tasks:        make(map[string]*models.Task),
		expressions:  make(map[string]*models.Expression),
		pendingTasks: make(chan *models.Task, 100),
		mu:           sync.Mutex{},
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

	log.Printf("Added expression %s with %d tasks", id, len(tasks))
	for _, task := range tasks {
		task.Status = "pending"
		o.tasks[task.ID] = task
		o.pendingTasks <- task
		log.Printf("Added task %s for expression %s (status: %s)", task.ID, id, task.Status)
	}

	return id, nil
}

func (o *Orchestrator) GetTask() (*models.Task, error) {
	select {
	case task := <-o.pendingTasks:
		o.mu.Lock()
		defer o.mu.Unlock()

		task.Status = "processing"
		log.Printf("Assigned task %s to agent", task.ID)
		return task, nil
	default:
		return nil, errors.New("no tasks available")
	}
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

	allTasksCompleted := true
	for _, t := range o.tasks {
		if t.ExpressionID == task.ExpressionID && t.Status != "completed" {
			allTasksCompleted = false
			break
		}
	}

	if allTasksCompleted {
		expr := o.expressions[task.ExpressionID]
		expr.Status = "completed"
		expr.Result = fmt.Sprintf("%g", calculateExpressionResult(task.ExpressionID, o.tasks))
		log.Printf("Expression %s completed with result %s", expr.ID, expr.Result)
	}

	return nil
}

func (o *Orchestrator) GetAllExpressions() []*models.Expression {
	o.mu.Lock()
	defer o.mu.Unlock()

	expressions := make([]*models.Expression, 0, len(o.expressions))
	for _, expr := range o.expressions {
		expressions = append(expressions, expr)
	}

	return expressions
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

func calculateExpressionResult(expressionID string, tasks map[string]*models.Task) float64 {
	var result float64
	for _, task := range tasks {
		if task.ExpressionID == expressionID {
			switch task.Operation {
			case "+":
				result += task.Result
			case "-":
				result -= task.Result
			case "*":
				result *= task.Result
			case "/":
				result /= task.Result
			}
		}
	}
	return result
}

func parseExpressionToTasks(expressionID, expr string) ([]*models.Task, error) {
	var tasks []*models.Task

	expr = strings.ReplaceAll(expr, " ", "")

	tokens := strings.FieldsFunc(expr, func(r rune) bool {
		return !isDigit(r) && r != '.'
	})

	operations := make([]string, 0)
	for _, char := range expr {
		if isOperator(char) {
			operations = append(operations, string(char))
		}
	}

	if len(tokens) == 1 && len(operations) == 0 {
		arg1, err := strconv.ParseFloat(tokens[0], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", tokens[0])
		}

		task := &models.Task{
			ID:            generateID(),
			ExpressionID:  expressionID,
			Arg1:          arg1,
			Arg2:          0,
			Operation:     "+",
			OperationTime: 0,
			Status:        "completed",
			Result:        arg1,
		}

		tasks = append(tasks, task)
		return tasks, nil
	}

	if len(tokens) != len(operations)+1 {
		return nil, fmt.Errorf("invalid expression: number of tokens and operations mismatch")
	}

	for i := 0; i < len(operations); i++ {
		arg1, err := strconv.ParseFloat(tokens[i], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", tokens[i])
		}

		arg2, err := strconv.ParseFloat(tokens[i+1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", tokens[i+1])
		}

		task := &models.Task{
			ID:            generateID(),
			ExpressionID:  expressionID,
			Arg1:          arg1,
			Arg2:          arg2,
			Operation:     operations[i],
			OperationTime: 1000,
			Status:        "pending",
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func isOperator(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}
