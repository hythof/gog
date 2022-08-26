package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type Gog struct {
	db *sql.DB
}

func Open(driverName, dataSourceName string) (*Gog, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &Gog{db}, nil
}

func (g *Gog) Close() error {
	return g.db.Close()
}

func First[T any](g *Gog) (*T, error) {
	row := new(T)
	table:=strings.ToLower(fmt.Sprintf("%T",*row)[5:])
	query:= fmt.Sprintf("SELECT * FROM %s LIMIT 1", table)
	println(query)
	rows, err := g.db.Query(query)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		rows.Scan(&row)
		return row, nil
	} else {
		return nil, fmt.Errorf("No record matched")
	}
}

func main() {
	gog, err := Open("postgres", "host=localhost port=5432 user=hiroshi password=postgres dbname=testdb sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer gog.Close()
	person, err := First[Person](gog)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", person)
}

type Person struct {
	Name string
	Age  int
}
