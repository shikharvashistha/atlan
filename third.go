/*
In such cases, each response to the form becomes a row in
the sheet, and questions in the form become columns.
*/
package main

import (
	"github.com/plandem/xlsx"
	"github.com/plandem/xlsx/format/styles"
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
		xl := xlsx.New()
		defer xl.Close()
		sheet := xl.AddSheet("Initial Sheet")
		cell := sheet.CellByRef("A1")
		cell.SetStyles(styles.New(
			styles.Font.Bold,
			styles.Font.Color("#ff0000"),
			styles.Fill.Type(styles.PatternTypeSolid),
			styles.Fill.Color("#ffff00"),
			styles.Border.Color("#009000"),
			styles.Border.Type(styles.BorderStyleMedium),
		))
		for iRow := 1; iRow < 7; iRow++ {
			cell := sheet.Cell(1, iRow)
			cell.SetValue(iRow)//store responses
		}
		for iCol := 2; iCol < 7; iCol++ {
			cell := sheet.Cell(iCol, 1)
			cell.SetValue(iCol)//store questions
		}
	}
}