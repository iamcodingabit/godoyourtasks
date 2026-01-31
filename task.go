package main

import (
	"fmt"
	"slices"
)

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
	fmt.Println()
}

func (t *Task) Add(task_name string, task_status string){
		new_task := Task{
			TaskName: task_name,
			TaskStatus: task_status,
		}
		Tasks = append(Tasks, new_task)
		t.list()
}

func (t *Task) Finish(task_number int, new_task_status string){
	index := task_number-1 
	Tasks[index].TaskStatus = new_task_status
	t.list()
}

func (t *Task) Remove(task_number int){
	index := task_number-1 
	Tasks = slices.Delete(Tasks, index, index+1)
	t.list()
}