package main

import (
	"fmt"
	"net/http"
	"database/sql"
	"github.com/thedevsaddam/govalidator"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Age  string   `json:"age"`
	Gender byte `json:"gender"`
	Hobbies []string `json:"hobbies"`
	Mobile string `json:"mobile"`
	Location string `json:"location"`
}

func handler(w *http.ResponseWriter, r *http.Request) {// Handler for the HTTP request
	var user User// User struct
	rules := govalidator.MapData{// Define the rules for the struct
		user.Name: []string{"required", "between:3,8"},// Name is required and must be between 3 and 8 characters long.
		user.Age: []string{"required", "between:18,30"},// Age is required and must be between 18 and 30.
		user.Mobile: []string{"required", "digits:11"},// Mobile is required and must be 11 digits long.
	}
	opts := govalidator.Options{// Define the options for the struct
		Request:         r,// Use the HTTP request as the source of data
		Rules:           rules,// Use the rules defined above
		RequiredDefault: true,// If a field is not provided in the request, set it to the default value
	}
	v:= govalidator.New(opts)// Create a new validator
	e:= v.Validate()// Validate the struct
	err := map[string]interface{}{"validationError": e}// If there are errors, return them in a map
	if err != nil {// If there are errors, return them in a map
	    panic(err)
	}
}

func main(){// Main function
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/testdb")// Connect to the database
	if err != nil {// If there is an error, return it
		panic(err.Error())// Return the error
	}
	defer db.Close()//defer the database connection
	insert, err:=db.Query("INSERT INTO users VALUES('Name',19,'M', 'Basketball, Cricket', '98xxxxxx', 'Delhi')")// Insert a new user
	if err != nil {// If there is an error, return it
		panic(err.Error())
	}
	defer insert.Close()//defer the database connection
	results, err := db.Query("SELECT Name, Age, Mobile FROM testdb")// Select all users
	if err != nil {// If there is an error, return it
		panic(err.Error())
	}
	defer results.Close()//defer the database connection
	for results.Next() {// Iterate through the results
		handler(nil, nil)// Call the handler function
	    fmt.Println("Listening on port: 8080")// Print the port
	    http.ListenAndServe(":8080", nil)// Start the HTTP server
	}
}