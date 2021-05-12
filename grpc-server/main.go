package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/mirzaahmedov/golang/api"
	"google.golang.org/grpc"
)

type Server struct {
	api.UnimplementedContactProviderServer
}

func (s *Server) GetAll(ctx context.Context, empty *api.Test) (*api.Contacts, error) {
	var (
		result              []*api.Contact
		id, phoneNumber     int64
		firstName, lastName string
	)
	db, err := sql.Open("postgres", " host=localhost port=5432 database=test sslmode=disable ")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM contacts;")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		if err := rows.Scan(&id, &firstName, &lastName, &phoneNumber); err != nil {
			log.Fatal(err)
		}
		result = append(result, &api.Contact{
			Id:          id,
			FirstName:   firstName,
			LastName:    lastName,
			PhoneNumber: phoneNumber,
		})
	}
	return &api.Contacts{
		Length:   int64(len(result)),
		Contacts: result,
	}, nil
}

func (s *Server) Add(ctx context.Context, contact *api.Contact) (*api.Id, error) {
	var id int64
	db, err := sql.Open("postgres", " host=localhost port=5432 database=test sslmode=disable ")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	row := db.QueryRow("INSERT INTO contacts ( first_name, last_name, phone_number ) VALUES ( $1, $2, $3 ) RETURNING id ;", contact.FirstName, contact.LastName, contact.PhoneNumber)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&id)
	return &api.Id{
		Id: id,
	}, nil
}

func (s *Server) Remove(_ context.Context, target *api.Id) (*api.Contact, error) {
	var (
		id, phone_number      int64
		first_name, last_name string
	)
	db, err := sql.Open("postgres", " host=localhost port=5432 database=test sslmode=disable ")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	row := db.QueryRow("DELETE FROM contacts WHERE id = $1 RETURNING id , first_name , last_name, phone_number ;", target.Id)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&id, &first_name, &last_name, &phone_number)
	return &api.Contact{
		Id:          id,
		FirstName:   first_name,
		LastName:    last_name,
		PhoneNumber: phone_number,
	}, nil
}

func (s *Server) Update(_ context.Context, contact *api.ContactUpdate) (*api.Contact, error) {
	var (
		id, phone_number      int64
		first_name, last_name string
	)
	db, err := sql.Open("postgres", " host=localhost port=5432 database=test sslmode=disable ")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	row := db.QueryRow("UPDATE contacts SET first_name = $2, last_name = $3 , phone_number = $4 WHERE id = $1 RETURNING id , first_name , last_name, phone_number ;", contact.Id, contact.Contact.FirstName, contact.Contact.LastName, contact.Contact.PhoneNumber)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&id, &first_name, &last_name, &phone_number)
	return &api.Contact{
		Id:          id,
		FirstName:   first_name,
		LastName:    last_name,
		PhoneNumber: phone_number,
	}, nil
}

func (s *Server) Get(_ context.Context, target *api.Id) (*api.Contact, error) {
	var (
		id, phone_number      int64
		first_name, last_name string
	)
	db, err := sql.Open("postgres", " host=localhost port=5432 database=test sslmode=disable ")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	row := db.QueryRow("SELECT * FROM contacts WHERE id = $1;", target.Id)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&id, &first_name, &last_name, &phone_number)
	return &api.Contact{
		Id:          id,
		FirstName:   first_name,
		LastName:    last_name,
		PhoneNumber: phone_number,
	}, nil
}

func main() {
	srv := &Server{}
	grs := grpc.NewServer()
	api.RegisterContactProviderServer(grs, srv)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := grs.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
