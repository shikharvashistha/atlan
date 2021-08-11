/*
One of our clients wanted to search for slangs (in local language) for an answer to a
text question on the basis of cities (which was the answer to a different MCQ
question)
*/

package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	g "https://github.com/serpapi/google-search-results-golang"
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
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err:=db.Query("INSERT INTO users VALUES('Name',19,'M', 'Basketball, Cricket', '98xxxxxx', 'Delhi')")
	results, err := db.Query("SELECT location FROM user")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var user User
		err := results.Scan(&user.Location)
		if err != nil {
			panic(err.Error())
		}
		parameter :=map[string]string{
			"q": "Slangs",
			"location": user.Location,
		}
		query := NewGoogleSearch(parameter, apikey)//we need to pass the api key
		res, err := query.json()
		if err != nil {
			panic(err.Error())
		}
		results := rsp["Slangs"].([]interface{})
		first_result := results[0].(map[string]interface{})
		fmt.Println(ref["title"].(string))
	}
}