/*
One of our clients wanted to search for slangs (in local language) for an answer to a
text question on the basis of cities (which was the answer to a different MCQ
question)
*/

package main

import (
	"fmt"
	"os"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	g "github.com/serpapi/google-search-results-golang"
)

type User struct {
	Name string //`json:"name"`
	Age  int   //`json:"age"`
	Gender byte //`json:"gender"`
	Hobbies []string //`json:"hobbies"`
	Mobile string //`json:"mobile"`
	Location string //`json:"location"`
}

func main(){
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/testdb")//Open connection to database
	if err != nil {// handle error
		panic(err.Error())
	}
	defer db.Close()//defer closing connection to database
	insert, err:=db.Query("INSERT INTO users VALUES('Name',19,'M', 'Basketball, Cricket', '98xxxxxx', 'Delhi')")//inserting values into database
	if err != nil {// handle error
		panic(err.Error())//panic if error occurs
	}
	defer insert.Close()//defer closing insert query
	results, err := db.Query("SELECT location FROM testdb")//selecting values from database
	if err != nil {// handle error
		panic(err.Error())
	}
	defer results.Close()//defer closing results query
	for results.Next() {//iterating through results
		var user User//creating user struct
		err := results.Scan(&user.Location)//scanning values into user struct
		if err != nil {// handle error
			panic(err.Error())//panic if error occurs
		}
		parameter :=map[string]string{//creating parameter map
			"q": "Slangs",//search query
			"location": user.Location,//location
		}
		query := g.NewGoogleSearch(parameter, os.Getenv("API_KEY"))//creating google search object
		rsp, err := query.GetJSON()//getting json response
		if err != nil {// handle error
			panic(err.Error())//panic if error occurs
		}
		results := rsp["Slangs"].([]interface{})//creating results array
		first_result := results[0].(map[string]interface{})//creating first result map
		fmt.Println(first_result["text"])//printing first result
		fmt.Println(rsp["title"].(string))//printing title
	}
}