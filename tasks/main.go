package main

import (
	"fmt"
	"log"

	"github.com/mirzaahmedov/goroadmap/tasks/db"
)

func main() {
  var output string
  client, err := db.NewDBClient()
  if err != nil {
    log.Fatal(err)
  }
  defer client.Close()
  output, err = client.Delete(1)
     if err != nil {
    log.Fatal(err)
  }
  fmt.Println(output)
}
