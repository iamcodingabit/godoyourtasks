package main

import (
	"fmt"
	//"maps"
)

func main() {
	fmt.Println("***** Another Obligatory To Do List App *****")
	/*
	keyword := map[string][]any{
		"help": {
		},
		"add": {
			"task",
			"function",
		},
		"remove": {
			"task",
		},
		"update": {
			"task",
		},
	}
		*/
	//Check if array is empty
	//currentTask := ""
	var task = &Task{}
	task.Add("Clean up room", "doing")
	task.Add("Code hash maps", "backlog")
}