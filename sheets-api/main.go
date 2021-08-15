/*
In such cases, each response to the form becomes a row in
the sheet, and questions in the form become columns.
*/
package main

import (
	"github.com/plandem/xlsx"// Import the package
	"github.com/plandem/xlsx/format/styles"// Import the package
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`//Name of the user
	Age  string   `json:"age"`//Age of the user
	Gender byte `json:"gender"`//Gender of user
	Hobbies []string `json:"hobbies"`//hobbies of user
	Mobile string `json:"mobile"`//mobile of user
	Location string `json:"location"`//User location
}


func main(){
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/testdb")//Connect to the database
	if err != nil {//If there is an error
		panic(err.Error())
	}
	defer db.Close()//defer closing the connection
	insert, err:=db.Query("INSERT INTO users VALUES('Collect',19,'M', 'Basketball, Cricket', '98xxxxxx', 'Delhi')")//Insert a row in the database
	if err != nil {//If there is an error
		panic(err.Error())
	}
	defer insert.Close()//defer closing the connection
	results, err := db.Query("SELECT * FROM testdb")//Select all the rows from the database
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()//defer closing the connection
	for results.Next() {//Iterate through the rows
		xl := xlsx.New()//Create a new workbook
		defer xl.Close()//defer closing the workbook
		sheet := xl.AddSheet("Initial Sheet")//Add a sheet to the workbook
		cell := sheet.CellByRef("A1")//Add a cell to the sheet
		cell.SetStyles(styles.New(//Set the styles
			styles.Font.Bold,//Bold
			styles.Font.Color("#ff0000"),//Red
			styles.Fill.Type(styles.PatternTypeSolid),//Solid fill
			styles.Fill.Color("#ffff00"),//Yellow
			styles.Border.Color("#009000"),//Green
			styles.Border.Type(styles.BorderStyleMedium),//Medium border
		))
		for iRow := 1; iRow < 7; iRow++ {//Iterate through the rows
			cell := sheet.Cell(1, iRow)//Add a cell to the sheet
			cell.SetValue(iRow)//Set the value of the cell

		}
		for iCol := 2; iCol < 7; iCol++ {//Iterate through the columns
			cell := sheet.Cell(iCol, 1)//Add a cell to the sheet
			cell.SetValue(iCol)//Set the value of the cell
		}
	}
}