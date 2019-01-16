package db

import (
	"fmt"

	"database/sql"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "ContactBook"
)

func InsertDB(mobile, name, email string) (err error) {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	fmt.Println("# Inserting values")

	var lastInsertId int

	err = db.QueryRow("INSERT INTO contact.details(mobile,email,name,create_date,status) VALUES($1,$2,$3,now(),$4) returning uid;", mobile, email, name, "ACTIVE").Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("last inserted id =", lastInsertId)
	return
}

//	_, err := db.Db.Exec("UPDATE contact.details SET name = $1, mobile =$2, email = $3,last_update=now(),status=$4 WHERE id=$5 ", inputName, inputMobile, inputEmail, "ACTIVE", contactId)

func UpdateDB(mobile, name, email, id string) (err error) {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	stmt, err = db.Prepare("UPDATE contact.details SET name = $1, mobile =$2, email = $3,last_update=now(),status=$4 WHERE id=$5 ", name, mobile, email, "ACTIVE", id)
	checkErr(err)
	fmt.Println("last inserted id =", lastInsertId)
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
	return
}
