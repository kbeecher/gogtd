package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kbeecher/gogtd/gogtd"
)

/*
STOP #11 Anonymous functions
We've already seen how functions can be parameters in other functions. We'll
see now that functions come in two flavours: named and anonymous (or "lambda
function" in Haskell terms).

The command "backend" (commands.go), which we'll see shortly, is a series of
functions that perform generic operations (i.e. query, add, update), but rely
on the user to fill in the detail in the form of another function.

For example, the CmdQueryMany function will search a TaskList and return some
of the tasks. Using which criteria does CmdQueryMany search? You have to give
it the criteria in the form of a function that takes a TaskList and returns a
(presumably smaller) TaskList.

Let's look at the first one, todo...
*/
func interpret(input string) string {
	if strings.HasPrefix(input, "todo") {
		return gogtd.CmdQueryMany(
			/* This argument creates an anonymous function, i.e. a nameless,
			on-the-fly function. It has just one line in it.

			Interestingly, that line contains a reference to a different,
			named function (IsPending). What we're doing is telling CmdQueryMany
			to execute the TaskList's Filter method. The Filter method, like
			the CmdQueryMany method, is generic, so you can provide your own
			implementation of a filter - so long as it's a function that
			takes a task and returns a bool.

			In the task.go file, we saw a couple of functions just like that:
			IsPending and DueToday. So, if we give one of those to the Filter
			method, we'll end up with a fully-functioning filter.
			*/
			func(tl *gogtd.TaskList) *gogtd.TaskList { return tl.Filter(gogtd.IsPending) },
		)
	} else if strings.HasPrefix(input, "today") {
		return gogtd.CmdQueryMany(
			func(tl *gogtd.TaskList) *gogtd.TaskList { return tl.Filter(gogtd.DueToday) },
		)
	} else if strings.HasPrefix(input, "all") {
		return gogtd.CmdQueryMany(
			func(tl *gogtd.TaskList) *gogtd.TaskList { return tl },
		)
	} else if strings.HasPrefix(input, "show") {
		return gogtd.CmdQuery(input)
	} else if strings.HasPrefix(input, "add") {
		return gogtd.CmdAdd(input)
	} else if strings.HasPrefix(input, "tick") {
		return gogtd.CmdUpdate(
			input, func(t *gogtd.Task) { t.Tick() },
		)
	} else if strings.HasPrefix(input, "untick") {
		return gogtd.CmdUpdate(
			input, func(t *gogtd.Task) { t.Untick() },
		)
	}
	return "Invalid command"
}

/*
STOP #9 Program entry point
Now that we've seen the two main parts of the program (Task and TaskList)
let's see how we construct a program using them.

Basically, this function keeps accepting a line from the user and passing it
to the interpret function.
*/
func main() {
	/* STOP #10 Declaring and initialising variables
	Go is strongly typed and variables must be declared before use. There are
	two ways to do it:

	1) var variable_name variable_type

	2) variable_name := variable_value

	The second one declares and initialises the variable in one step (it
	sliently takes the type of whatever is on the right-hand side of the :=).

	Variables in Go are a little different from variables in Haskell.
	In Haskell, variables (in lines line "let x = y + 1") are variables in the
	mathematical sense -- that is they	are immutable; once initialised they
	don't change in value.

	Variables in Go are mutable, so they're like labels for locations
	in memory.
	*/
	var input, output string
	consolereader := bufio.NewReader(os.Stdin)

	for input != "q\n" && input != "quit\n" {
		fmt.Print("> ")
		input, _ = consolereader.ReadString('\n')

		output = interpret(input)
		fmt.Println(output)
	}
}
