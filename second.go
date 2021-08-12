package main

import (
	"fmt"
	"net/http"
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
	return
}

func main(){
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err:=db.Query("INSERT INTO users VALUES('Collect',19,'M', 'Basketball, Cricket', '98xxxxxx', 'Delhi')")
	results, err := db.Query("SELECT Name, Age, Mobile FROM user")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	for results.Next() {
		http.HandleFunc("/", handler)
	fmt.Println("Listening on port: 8080")
	http.ListenAndServe(":8080", nil)
	}
}