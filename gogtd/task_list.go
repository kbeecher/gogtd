package gogtd

import (
	"bytes"
	"errors"
	"strconv"
)

/* STOP #6 Using in-built types
We're going to hold our list of tasks in a simple map structure. A map is
an in-built type in Go. "Tasks map[int]*Task" means that we're declaring a
field of TaskList to be Tasks, which maps integer key values to (pointers to)
Tasks.

Pointers are not part of the Haskell language.
*/
type TaskList struct {
	Tasks map[int]*Task
}

func NewTaskList() *TaskList {
	return &TaskList{
		// The built-in make method here initialises an empty map
		Tasks: make(map[int]*Task),
	}
}

func (tl *TaskList) AddTask(t *Task) {
	tl.Tasks[len(tl.Tasks)] = t
}

func (tl *TaskList) UpdateTask(id int, t *Task) {
	tl.Tasks[id] = t
}

// Here's an example of returning multiple values in order to provide
// error information. If all goes well, the error is nil.
func (tl *TaskList) GetTaskById(id int) (*Task, error) {
	newTask := NewTask()
	if id < 0 {
		return newTask, errors.New("ID must not be negative.")
	}
	if id >= len(tl.Tasks) {
		return newTask, errors.New("Task not found.")
	}

	return tl.Tasks[id], nil
}

/*
STOP #7 Functions as first-class objects
Go is not a functional language. This is true in the strict sense: functions
in Go are actually subroutines since they are a sequence of steps that operate
on mutable data (as opposed to the mathematical notion of a function). But it's
also true in the looser sense that it does not promote a functional style of
programming - for instance, the core higher-order functions (map, filter and
reduce) are not included.

However, functions in Go can be treated as first-class objects, created and
passed around between other functions just as integers and strings can be.
This means it is possible to write in a somewhat functional style and roll our
own higher-order functions like, say, a filter function.

Our Filter function works on a TaskList. It takes one argument: a
function, f, that takes a task and returns a bool. It works by applying f to
every element in the original list and adding the element to a new list if
that application returns true.
*/
func (tl *TaskList) Filter(f func(t *Task) bool) *TaskList {
	newTl := NewTaskList()

	/*
		STOP #8 Looping
		Go has only one looping construct: for. You can use it in
		the following ways:

		for init; condition; post { } // classic index-based

		for condition { } // like a while loop (condition can be empty to make it a while true)

		for key, value := range x // iterating over a collection
	*/
	for id, t := range tl.Tasks {
		if f(t) {
			newTl.Tasks[id] = t
		}
	}

	return newTl
}

func (tl *TaskList) String() string {
	var buffer bytes.Buffer

	for id, t := range tl.Tasks {
		buffer.WriteString(strconv.Itoa(id))
		buffer.WriteString(": " + t.String() + "\n")
	}

	return buffer.String()
}
