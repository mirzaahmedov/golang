package main

import "testing"

var tests = []Contact{
	{
		first_name:   "first name 1",
		last_name:    "last name 1",
		phone_number: 1234567891,
		profile_img:  "profile img 1",
	},
	{
		first_name:   "first name 2",
		last_name:    "last name 2",
		phone_number: 1234567892,
		profile_img:  "profile img 2",
	},
	{
		first_name:   "first name 3",
		last_name:    "last name 3",
		phone_number: 1234567893,
		profile_img:  "profile img 3",
	},
}

var mock = Contactlist{}

func TestCreate(t *testing.T) {

	for index, elem := range tests {
		mock.create(elem)
		if mock.contacts[index] != elem {
			t.Fatal("testing create method failed")
		}
	}
}

func TestDelete(t *testing.T) {
	mock.delete(tests[0])
	if len(mock.contacts) != len(tests)-1 {
		t.Fatal("testing delete method failed")
	}
}

func TestEdit(t *testing.T) {
	for index, elem := range mock.contacts {
		mock.edit(elem, Contact{
			first_name:   "after testing",
			last_name:    "after testing",
			profile_img:  "after testing",
			phone_number: 7777777,
		})
		if !(mock.contacts[index].first_name == "after testing" && mock.contacts[index].last_name == "after testing" && mock.contacts[index].profile_img == "after testing" && mock.contacts[index].phone_number == 7777777) {
			t.Fatal("testing edit method failed")
		}
	}
}
