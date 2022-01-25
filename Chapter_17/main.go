package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// map this type to the record in the table
type Course struct {
	ID     string
	Detail string
}

func getResults(db *sql.DB) {
	results, err := db.Query("SELECT * FROM Course")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var course Course

		err = results.Scan(&course.ID, &course.Detail)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(course.ID, "-", course.Detail)
	}
}

func main() {
	// use mysql as driverName and a valid DSN
	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/CoursesDB")
	// handle error
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database object created")
		getResults(db)
	}
	// defer the close till after the main function has finished executing
	defer db.Close()
}
