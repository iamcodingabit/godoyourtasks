package main

import "fmt"

//"fmt"

type Task struct {
	TaskName string
	TaskStatus string // backlog, in-progress, review, finished	
}

var Tasks []Task

func (t *Task) list(){
	for i, v := range Tasks{
		fmt.Printf("%d. %s - %s \n", i+1, v.TaskName, v.TaskStatus)
	}
}

func (t *Task) Add(task_name string, task_status string){
		new_task := Task{
			TaskName: task_name,
			TaskStatus: task_status,
		}
		Tasks = append(Tasks, new_task)
		t.list()
}