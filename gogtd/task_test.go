package gogtd

import (
	"testing"
	"time"
)

func getTestTask() *Task {
	return &Task{
		Description: "Hello",
		Done:        false,
		Due:         time.Now(),
	}
}

func getTestList() *TaskList {
	newTl := NewTaskList()
	newTl.AddTask(getTestTask())
	return newTl
}

func TestTick(t *testing.T) {
	testTask := getTestTask()
	testTask.Tick()
	if !testTask.Done {
		t.Error("Task was not ticked.")
	}
}

func TestDueToday(t *testing.T) {
	testTask := getTestTask()
	if !testTask.DueToday() {
		t.Error("Task should be due today.")
	}
}

func TestGetPending(t *testing.T) {
	tasks := getTestList()
	if len(tasks.GetPending().Tasks) == 0 {
		t.Error("Should be at least 1 pending task!")
	}
}

func TestTodays(t *testing.T) {
	tasks := getTestList()
	if len(tasks.Todays().Tasks) == 0 {
		t.Error("Should be at least 1 task today!")
	}
}
