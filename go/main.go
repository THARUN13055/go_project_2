package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

const (
	apiPaths = "/database"
)

type Book struct {
	Id, Name, Isbn string
}

type library struct {
	dbHost, dbPass, dbName , dbUser string
}

type book struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	ISBN string `json:"isbn"`
}

func main() {
	// DB_HOST is of the form host:port

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "root"
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost:3306"
	}
	// DB_PASS
	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		dbPass = "password"
	}

	// apiPath

	apiPath := os.Getenv("API_PATH")
	if apiPath == "" {
		apiPath = apiPaths
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "library"
	}
	l := library{
		dbUser: dbUser,
		dbHost: dbHost,
		dbPass: dbPass,
		dbName: dbName,
		
	}

	r := mux.NewRouter()
	r.HandleFunc(apiPath, l.getBooks).Methods("GET")
	r.HandleFunc(apiPath, l.createBook).Methods("POST")
	http.ListenAndServe(":8080", r)
}

// POST method
func (l library) createBook(w http.ResponseWriter, r *http.Request) {
	book := Book{}

	json.NewDecoder(r.Body).Decode(&book)

	db := l.openConnection()
	defer l.closeConnection(db)

	insertQuery,err := db.Prepare("insert into book values(? ? ? )")

	if err != nil {
		log.Fatalf("preparing the db qurey %s \n", err.Error())
	}

	tx, err := db.Begin()

	if err != nil {
		log.Fatalf("while begining the transaction %s\n", err.Error())
	}

	_, err = tx.Stmt(insertQuery).Exec(book.Id,book.Name,book.Isbn)

	if err != nil{
		log.Fatalf("execing the insert commond %s\n",err.Error())
	}
	err = tx.Commit()
	if err != nil{
		log.Fatalf("while commint the transaction %s\n",err.Error())
	}

}

// GET Method
func (l library) getBooks(w http.ResponseWriter, r *http.Request) {
	db := l.openConnection()
	defer l.closeConnection(db)
    
	rows, err := db.Query("SELECT * FROM book")
	if err != nil {
		log.Fatalf("Error querying the database: %s\n", err.Error())
	}

	defer rows.Close()

	books := []book{}

	for rows.Next() {
		var id, name, isbn string
		err := rows.Scan(&id, &name, &isbn)
		if err != nil {
			log.Fatalf("Error while scanning the row: %s\n", err.Error())
		}
		aBook := book{
			ID:   id,
			Name: name,
			ISBN: isbn,
		}
		books = append(books, aBook)
	}

	json.NewEncoder(w).Encode(books)
}



// Connection of db
func (l library) openConnection() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", l.dbUser, l.dbPass, l.dbHost, l.dbName))
	if err != nil {
		log.Fatalf("Error opening the database: %s\n", err.Error())
	}
	return db
}

// close of db
func (l library) closeConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("Error closing the database connection: %s\n", err.Error())
	}
}
