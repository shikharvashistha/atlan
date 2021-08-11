package main

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/thedevsaddam/govalidator"
	_"github.com/go-sql-driver/mysql"
)

type User struct {
	Name string //`json:"name"`
	Age  int   //`json:"age"`	
	Gender char //`json:"gender"`
	Hobbies []string //`json:"hobbies"`
	Mobile string //`json:"mobile"`
	Location string //`json:"location"`
}

func handler(w *http.ResponseWriter, r *http.Request) {
	rules := govalidator.MapData{
		"name": &govalidator.StringRule{Min: 1, Max: 10, ErrorMsg: "name must be between 1 and 10 characters long"},
		"age": &govalidator.IntRule{Min: 1, Max: 100, ErrorMsg: "age must be between 1 and 100"},
		"mobile": &govalidator.IntRule{Min: 1, Max: 10, ErrorMsg: "mobile must be between 1 and 10 digits long"},
	}
	v:= govalidator.New(rules)
	e:= v.Validate()
	err := map[string]interface{}{"validationError": e}
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
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
	defer results.Close()
}