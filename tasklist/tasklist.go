package main

import (
	"fmt"
	"time"
)

type Task struct {
	title string
	done bool
	date string
}

type Tasklist struct {
	tasks []Task
}

func (t *Tasklist) read() {
	var output string
	for _, task := range t.tasks {
		output += fmt.Sprintf("\nTask: %s\nDone: %v\nDate: %s\n", task.title, task.done, task.date)
	}
	fmt.Println("----------- Tasks ------------\n", output) 
}

func (t *Tasklist) create(data Task) {
	data.date = time.Now().String()
	t.tasks = append(t.tasks, data)
}

func (t *Tasklist) toggle(data Task) {
	for i, task := range t.tasks {
		if data.title == task.title {
			t.tasks[i].done = !t.tasks[i].done
			break
		}
	} 
}

func (t *Tasklist) delete(data Task) {
	for i, task := range t.tasks {
		if data.title == task.title {
			t.tasks[i] = t.tasks[len(t.tasks)-1]
			t.tasks = t.tasks[:len(t.tasks)-1]
			break
		}
	} 
}

func main() {
	var tasklist Tasklist
	tasklist.create(Task{ title:  "Learn Go Language",done: false})
	tasklist.create(Task{ title:  "Learn Unit testing",done: false})
	tasklist.create(Task{ title:  "Learn English",done: false})
	tasklist.create(Task{ title:  "Make new Projects",done: false})
	tasklist.read()
	tasklist.toggle(Task{ title: "Learn Go Language" })
	tasklist.toggle(Task{ title: "Learn English" })
	tasklist.read()
	tasklist.toggle(Task{ title: "Learn Go Language" })
	tasklist.delete(Task{ title: "Learn Unit testing" })
	tasklist.read()
}