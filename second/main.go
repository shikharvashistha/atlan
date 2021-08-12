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

func handler(w *http.ResponseWriter, r *http.Request) {
	var user User
	rules := govalidator.MapData{
		user.Name: []string{"required", "between:3,8"},
		user.Age: []string{"required", "between:18,30"},
		user.Mobile: []string{"required", "digits:11"},
	}
	opts := govalidator.Options{
		Request:         r,
		Rules:           rules,
		RequiredDefault: true,
	}
	v:= govalidator.New(opts)
	e:= v.Validate()
	err := map[string]interface{}{"validationError": e}
	if err != nil {
	    panic(err)
	}
}

func main(){
	db, err := sql.Open("mysql", "user:password@tcp(127.0.2.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err:=db.Query("INSERT INTO users VALUES('Name',19,'M', 'Basketball, Cricket', '98xxxxxx', 'Delhi')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	results, err := db.Query("SELECT Name, Age, Mobile FROM testdb")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	for results.Next() {
		handler(nil, nil)
	    fmt.Println("Listening on port: 8080")
	    http.ListenAndServe(":8080", nil)
	}
}