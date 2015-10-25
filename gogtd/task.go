/*
STOP #1 Initial stuff

Let's start with some simple stuff. The most fundamental part
of our task manager is the definition of the task itself.

In Go, every file must begin with a package statement which states to which
package the file belongs. Packages are a way of structuring your code in
a Go project.

If you are reusing code from other packages, you also need to include an import
statement which lists the packages you want access to. In this case, we want
code for outputting to the console ("fmt") and processing dates and times
("time").

More: https://golang.org/doc/code.html#PackagePaths
*/

package gogtd

import (
	"fmt"
	"time"
)

/*
STOP #2 Structs
Structs are very similar to their namesake in ancestor languages (like C) and
records from Haskell, in that they group together fields into a new custom type.
However, Go adds a little more into the mix...
*/
type Task struct {
	Description string
	Done        bool
	Due         time.Time
}

/*
STOP #3 Methods
Go allows you to attach functions to particular types - such functions are
called methods. Go is not strictly object-oriented, although methods behave a
bit like methods from proper OO languages.

You can tell a method apart from a function like so:

A function looks something like this: "func MyFunc()"

Whereas a method looks like this: "func (t *MyType) MyMethod()"

Take the following two methods, Tick and Untick. They are methods on the Task
type we created above. Therefore, when we create an instance of a task later,
we can call its methods like this: "myTask.Tick()"
*/
func (t *Task) Tick() {
	t.Done = true
}

func (t *Task) Untick() {
	t.Done = false
}

/*
Something to note: Go has a very simple way to specify whether a type or
function is visible outside the package its defined in. If the first letter
of the name is upper-case, it's visible; otherwise it's not visible.

In our case, Tick and Untick are visible outside this package. If they where
named tick and untick, then other packages wouldn't be able to access them.
*/

/*
STOP #4 Return types
Here's another simple function that behaves kind of like a constructor,
because it returns a newly-instantiated Task. You can tell it does this from
two things.

Firstly, the return type (in Go, the return type is declared immediately
following the function's parameter list). In this case, the return type
is a task - or rather, a pointer to a Task...
*/
func NewTask() *Task {
	/* ... Secondly, the return statement includes the construction of a
	new Task instance with (some of) its fields ready-instantiated. The
	ampersand denotes that we're getting the address of this new instance
	since we're returning a Task poiner.

	All-in-all the pointer stuff is very similar to what you see in C/C++
	although made somewhat safer (e.g. pointer arithmetic is not allowed).
	*/
	return &Task{
		Due: time.Now(),
	}
}

/*
STOP #5 Multiple return types
Another unusual feature of Go is that functions can return multiple values.
In the following example, we use the time.Date() method, which returns the
year, month and day values... all at once! In Haskell, you can achieve the
same thing using a tuple.

It's commonly used in Go to write functions that might produce errors.
You'll see functions that are defined something like this:

func openDataFile(path string) (fileHandle, error)

In this case, the return type is a tuple, meaning it returns multiple values.
Using this function would therefore look something like this:

fh, err := openDataFile(path)
if err != nil {
	fmt.Printf("Error when opening file")
}

You'll see this pattern quite a lot in Go code.
*/
func DueToday(t *Task) bool {
	yearNow, monthNow, dayNow := time.Now().Date()
	yearTask, monthTask, dayTask := t.Due.Date()

	return yearNow == yearTask &&
		monthNow == monthTask &&
		dayNow == dayTask
}

func IsPending(t *Task) bool {
	return !t.Done
}

func (t *Task) String() string {
	var done string
	if t.Done {
		done = "(DONE)"
	} else {
		done = ""
	}

	return fmt.Sprintf("'%s' due %v %v %v %v",
		t.Description,
		t.Due.Day(),
		t.Due.Month(),
		t.Due.Year(),
		done,
	)
}
