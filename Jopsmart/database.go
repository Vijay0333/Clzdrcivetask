package main

import (
	"database/sql"
	"fmt"
	"log" // lowercase 'log'

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open connection
	con, err := sql.Open("mysql", "root@localhost:9682Vijay@s(127.0.0.1:3306)/201b375_vijay") // removed spaces in connection string
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	stmt, err := con.Prepare("SELECT  idemployee, employeename, employeesallery, employeeaddress, employeeworktime, employeeweight FROM employees WHERE  idemployee=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var employeename, employeesallery, employeeaddress, employeeworktime, employeeweight string
	//var empName, email string

	// Execute query/stmt
	err = stmt.QueryRow(1).Scan(&employeename, &employeesallery, &employeeaddress, &employeeworktime, &employeeweight) // Passing parameter '1' for placeholder
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("employeename: %s, employeesallery: %s, employeeaddress: %s, employeeworktime: %s, employeeweight: %s", employeename, employeesallery, employeeaddress, employeeworktime, employeeweight)
}
