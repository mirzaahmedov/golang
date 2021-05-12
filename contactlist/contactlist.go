package main

import "fmt"

type Contact struct {
	first_name   string
	last_name    string
	profile_img  string
	phone_number uint
}

type Contactlist struct {
	contacts []Contact
}

func (c *Contactlist) create(data Contact) {
	c.contacts = append(c.contacts, data)
}

func (c *Contactlist) delete(data Contact) {
	for i, contact := range c.contacts {
		if data.phone_number == contact.phone_number {
			c.contacts[i] = c.contacts[len(c.contacts)-1]
			c.contacts = c.contacts[:len(c.contacts)-1]
			break
		}
	}
}

func (c *Contactlist) edit(data Contact, modifications Contact) {
	for i, contact := range c.contacts {
		if data.phone_number == contact.phone_number {
			c.contacts[i] = modifications
		}
	}
}

func (c Contactlist) read() {
	var output string
	for _, contact := range c.contacts {
		output += fmt.Sprintf("\nFirst name: %s\nLast name: %s\nProfile picture: %s\nPhone number: %d\n", contact.first_name, contact.last_name, contact.profile_img, contact.phone_number)
	}
	fmt.Println("------- Contacts -------\n", output)
}

func main() {
	var contacts Contactlist
	contacts.create(Contact{first_name: "Bekzod", last_name: "Mirzaahmedov", phone_number: 990303245})
	contacts.create(Contact{first_name: "John", last_name: "Adams", profile_img: "https://picsum.photos/200/200", phone_number: 123555544})
	contacts.create(Contact{first_name: "Emily", last_name: "Taylor", profile_img: "https://picsum.photos/200/200", phone_number: 123456879})
	contacts.create(Contact{first_name: "Johongir", last_name: "Foziljonov", profile_img: "https://picsum.photos/200/200", phone_number: 1111111111})
	contacts.create(Contact{first_name: "Akmal", last_name: "Mirzayev", profile_img: "https://picsum.photos/200/200", phone_number: 997777777})
	contacts.read()
	contacts.edit(Contact{phone_number: 990303245}, Contact{phone_number: 990303245, first_name: "Michael", last_name: "Atkinson"})
	contacts.read()
	contacts.delete(Contact{phone_number: 1111111111})
	contacts.read()
}
