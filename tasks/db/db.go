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
    Title string`json:"title"`
    Done bool`json:"done"`
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
  var ( 
    id int
    title string
    done bool
  )
  for rows.Next() {
    rows.Scan(&id,&title,&done)
    rowsArray = append(rowsArray, Model{ Id: id, Title: title, Done: done, })
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
  return db.exec("SELECT * FROM tasks;")
}

func (db dbClient) Get(id int) (string, error){
  var query string = fmt.Sprintf("SELECT * FROM tasks WHERE id = %v ;", id )
  return db.exec(query)
}

func (db dbClient) Delete(id int) (string, error){
  var query string = fmt.Sprintf("DELETE FROM tasks WHERE id = %v ;", id )
  return db.exec(query)
}

func (db dbClient) Update(id int) (string, error){
  var query string = fmt.Sprintf("UPDATE tasks SET done = NOT done WHERE id = %v ;", id )
  return db.exec(query)
}

func (db dbClient) Add( data Model ) (string, error) {
  var query string = fmt.Sprintf("INSERT INTO tasks ( title , done ) VALUES ( '%v', %v );", data.Title, data.Done )
  return db.exec(query)
}
