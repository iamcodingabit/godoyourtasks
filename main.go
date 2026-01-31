package main

import (
	"fmt"
	//"maps"
)

func main() {
	fmt.Println("***** Another Obligatory To Do List App *****")
	
	var task = &Task{}
	task.Add("Clean up room", "doing")
	task.Add("Code hash maps", "backlog")
	task.Add("Code array sorting algo", "backlog")
	task.Finish(1, "finished")
	task.Remove(1)

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
}