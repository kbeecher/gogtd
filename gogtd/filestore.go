package gogtd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createTaskFile() *os.File {
	taskFile, err := os.Create("tasks.txt")
	check(err)

	return taskFile
}

func GetTasksFromFile() *TaskList {
	taskFile, err := os.Open("tasks.txt")
	if err != nil {
		taskFile = createTaskFile()
	}
	defer taskFile.Close()

	scanner := bufio.NewScanner(taskFile)
	taskList := NewTaskList()
	var taskLine string

	for scanner.Scan() {
		taskLine = scanner.Text()
		taskList.AddTask(parseLineToTask(taskLine))
	}

	return taskList
}

func SaveTaskList(tl *TaskList) error {
	err := os.Rename("tasks.txt", "tasks.txt.bak")
	check(err)
	taskFile := createTaskFile()
	defer taskFile.Close()
	writer := bufio.NewWriter(taskFile)

	var desc, done, dueDate string
	for _, t := range tl.Tasks {
		desc = t.Description
		if t.Done {
			done = "1"
		} else {
			done = "0"
		}
		dueDate = t.Due.Format("2006-01-02")
		fmt.Fprintf(writer, "%s\t%s\t%s\n", desc, done, dueDate)
	}
	writer.Flush()

	return nil
}

func parseLineToTask(taskLine string) *Task {
	taskParts := strings.Split(taskLine, "\t")
	newTask := NewTask()
	newTask.Description = taskParts[0]
	if taskParts[1] == "0" {
		newTask.Done = false
	} else {
		newTask.Done = true
	}
	// Note specific syntax in format string!
	newTask.Due, _ = time.Parse("2006-01-02", taskParts[2])

	return newTask
}
