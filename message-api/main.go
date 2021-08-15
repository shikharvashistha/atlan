package main

import (
	"fmt"
	"github.com/leozz37/hare"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {// Initiallize a struct with a json tag
	Name string `json:"name"`// Set the name field to be the json tag
	Age  string   `json:"age"`// Set the age field to be the json tag
	Gender string `json:"gender"`// Set the gender field to be the json tag
	Hobbies string `json:"hobbies"`// Set the hobbies field to be the json tag
	Mobile string `json:"mobile"`// Set the mobile field to be the json tag
	Location string `json:"location"`// Set the location field to be the json tag
}


func main(){
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/testdb")// Open a database connection
	if err != nil {//If there is an error, handle it
		panic(err.Error())
	}
	defer db.Close()// defer closing the connection when done
	insert, err:=db.Query("INSERT INTO users VALUES('Collect',19,'M', 'Basketball, Cricket', '98xxxxxx', 'Delhi')")// Insert a record
	if err != nil {// If there is an error, handle it
		panic(err.Error())
	}
	defer insert.Close()// Close the insert statement
	results, err := db.Query("SELECT * FROM testdb")// Select all records from the table
	if err != nil {// If there is an error, handle it
		panic(err.Error())// Print the error
	}
	defer results.Close()// Close the resultset
	for results.Next() {// Iterate through the resultset
		var user User// Create a new user
		err := hare.Send("8080", user.Name)// Send the user Name to the server
		hare.Send("8080", user.Age)// Send the user Age to the server
		hare.Send("8080", user.Gender)// Send the user Gender to the server
	 	hare.Send("8080", user.Hobbies)// Send the user Hobbies to the server
		if err != nil {// Handle the error if any
			panic(err)
		}
		 r, _ := hare.Listen("8080")// Listen to the port in case of a response

		for {
        	if r.HasNewMessages() {// If there is a new message
            	fmt.Println(r.GetMessage())// Print the message
        	}
    	}
	}
}