package gogtd

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

/*
STOP #12 The backend.
You can think of this as the backend. Control is passed here from the main
package so commands can be executed. It acts like a kind of query API.

Think of it like SQL. The types of operations are relatively few (adding,
selecting, updating), but most operations accept a function that specifies
the details of how that operation should be carried out.
*/

func CmdAdd(cmd string) string {
	task := NewTask()
	cmd = strings.Trim(cmd, "\n")
	cmdParts := strings.Split(cmd, "\t")

	if len(cmdParts) != 3 {
		return "Error reading input"
	}

	task.Description = cmdParts[1]
	dueDate, err := time.Parse("2006-01-02", cmdParts[2])
	if err != nil {
		return "Error reading date"
	}
	task.Due = dueDate

	taskList := GetTasksFromFile()
	taskList.AddTask(task)
	SaveTaskList(taskList)

	return "Done."
}

func CmdQuery(input string) string {
	t, err := findTaskById(input)
	if err != nil {
		return err.Error()
	}

	return t.String()
}

func CmdQueryMany(f func(*TaskList) *TaskList) string {
	tasks := GetTasksFromFile()
	resultSet := f(tasks)
	return resultSet.String()
}

func CmdUpdate(input string, f func(*Task)) string {
	id, err := parseTaskNumFromInput(input)
	if err != nil {
		return err.Error()
	}

	t, err := findTaskById(input)
	if err != nil {
		return err.Error()
	}

	f(t)

	taskList := GetTasksFromFile()
	taskList.UpdateTask(id, t)
	SaveTaskList(taskList)

	return "Done."
}

func findTaskById(input string) (*Task, error) {
	id, err := parseTaskNumFromInput(input)
	task := NewTask()
	if err != nil {
		return task, err
	}

	task, err = GetTasksFromFile().GetTaskById(id)
	if err != nil {
		return task, err
	}

	return task, nil
}

func parseTaskNumFromInput(input string) (int, error) {
	input = strings.Trim(input, "\n")
	inputParts := strings.Split(input, "\t")

	if len(inputParts) != 2 {
		return -1, errors.New("Error reading input")
	}

	i, err := strconv.Atoi(inputParts[1])
	if err != nil {
		return -1, errors.New("Task num not recognised")
	}
	return i, nil
}
