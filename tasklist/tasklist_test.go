package main

import "testing"

var tests = []Task{
	{
		title: "test title 1",
		done:  true,
	},
	{
		title: "test title 2",
		done:  false,
	},
	{
		title: "test title 3",
		done:  false,
	},
}

var mock = Tasklist{}

func TestCreate(t *testing.T) {

	for index, elem := range tests {
		mock.create(elem)
		if mock.tasks[index].title != elem.title {
			t.Fatal("testing create method failed")
		}
	}
}

func TestDelete(t *testing.T) {
	mock.delete(tests[0])
	if len(mock.tasks) != len(tests)-1 {
		t.Fatal("testing delete method failed")
	}
}

func TestToggle(t *testing.T) {
	for index, elem := range mock.tasks {
		value := elem.done
		mock.toggle(elem)
		if mock.tasks[index].done == value {
			t.Fatal("testing toggle method failed")
		}
	}
}
