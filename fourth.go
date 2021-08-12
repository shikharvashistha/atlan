package main

import (
	"fmt"
	"github.com/leozz37/hare"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Age  string   `json:"age"`
	Gender string `json:"gender"`
	Hobbies string `json:"hobbies"`
	Mobile string `json:"mobile"`
	Location string `json:"location"`
}


func main(){
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err:=db.Query("INSERT INTO users VALUES('Collect',19,'M', 'Basketball, Cricket', '98xxxxxx', 'Delhi')")
	results, err := db.Query("SELECT * FROM testdb")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	for results.Next() {
		var user User
		err := hare.Send("8080", user.Name)
		hare.Send("8080", user.Age)
		hare.Send("8080", user.Gender)
	 	hare.Send("8080", user.Hobbies)
		if err != nil {
			panic(err)
		}
		 r, _ := hare.Listen("8080")

		for {
        	if r.HasNewMessages() {
            	fmt.Println(r.GetMessage())
        	}
    	}
	}
}