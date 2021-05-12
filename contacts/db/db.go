package db

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

type (

  Model struct {
    Id int`json:"id"`
    FirstName string`json:"first_name"`
    LastName string`json:"last_name"`
    PhoneNumber int`json:"phone_number"`
  }

  dbClient struct {
    db *sql.DB
  }

  DBClient interface {
    Close()
    List()
    Get()
    Add()
    Update()
    Delete()
  }

)

func (db *dbClient) exec(query string) (string, error) {
  var rowsArray []Model
  rows , err := db.db.Query(query)
  if err != nil {
    return "", err
  }
  defer rows.Close()
  var id , phoneNumber int
  var firstName, lastName string
  for rows.Next() {
    rows.Scan(&id,&firstName,&lastName,&phoneNumber)
    rowsArray = append(rowsArray, Model{ Id: id, FirstName: firstName, LastName: lastName, PhoneNumber: phoneNumber, })
  }
  res, err := json.Marshal(rowsArray)
  if err != nil {
    return "", err
  }
  return string(res), nil
}


// constructor
func NewDBClient() ( dbClient, error ) {
  var client dbClient 
  db, err := sql.Open("postgres", " host=localhost port=5432 database=test sslmode=disable " ) 
  if err != nil {
    return client, err
  }
  client = dbClient{
    db: db,
  }
  return client, nil
}

// close db connection
func (db *dbClient)Close(){
  db.db.Close()
}

// list all
func (db dbClient) List() (string, error){
  return db.exec("SELECT * FROM contacts;")
}

func (db dbClient) Get(id int) (string, error){
  var query string = fmt.Sprintf("SELECT * FROM contacts WHERE id = %v ;", id )
  return db.exec(query)
}

func (db dbClient) Delete(id int) (string, error){
  var query string = fmt.Sprintf("DELETE FROM contacts WHERE id = %v ;", id )
  return db.exec(query)
}

func (db dbClient) Update(id int, data Model) (string, error){
  var query string = fmt.Sprintf("UPDATE contacts SET first_name = '%v', last_name = '%v', phone_number = %v WHERE id = %v ;", data.FirstName, data.LastName, data.PhoneNumber, id )
  return db.exec(query)
}

func (db dbClient) Add( data Model ) (string, error) {
  var query string = fmt.Sprintf("INSERT INTO contacts ( first_name , last_name, phone_number ) VALUES ( '%v', '%v', %v );", data.FirstName, data.LastName, data.PhoneNumber )
  return db.exec(query)
}
