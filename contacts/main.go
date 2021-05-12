package main

import (
	"fmt"
	"log"

	"github.com/mirzaahmedov/goroadmap/contacts/db"
)

func main() {
  var output string
  client, err := db.NewDBClient()
  if err != nil {
    log.Fatal(err)
  }
  defer client.Close()
  output, err = client.Get(3)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(output)
  output, err = client.Update(4, db.Model{
    FirstName: "Asror",
    LastName: "Jamolov",
    PhoneNumber: 111111111,
  })
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(output)
}
