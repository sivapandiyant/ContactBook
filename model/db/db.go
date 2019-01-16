package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "ContactBook"
)

func InsertDB(mobile, email, name string) (err error) {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	defer db.Close()

	fmt.Println("# Inserting values")

	var lastInsertId int

	err = db.QueryRow("INSERT INTO contact.details(mobile,email,name,create_date,status) VALUES($1,$2,$3,now(),$4) returning id;", mobile, email, name, "ACTIVE").Scan(&lastInsertId)
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	fmt.Println("last inserted id =", lastInsertId)
	return
}

func UpdateDB(mobile, email, name, id string) (err error) {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	defer db.Close()
	status := "ACTIVE"
	_, err = db.Exec("UPDATE contact.details SET name = $1, mobile =$2, email =$3,last_update=now(),status=$4 WHERE id=$5", name, mobile, email, status, id)
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	return
}

func UpdateStatus(id string) (err error) {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	defer db.Close()
	status := "INACTIVE"
	_, err = db.Exec("UPDATE contact.details SET status =$1 WHERE id= $2", status, id)

	return
}

func SelectDB(name, email string) (data [][]string, err error) {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	defer db.Close()

	var row *sql.Rows

	row, err = db.Query(`SELECT id, name,email,mobile,to_char(create_date,'YYYY-MM-DD HH24:MI:SS') FROM contact.details WHERE ($1='' OR LOWER(name) like '%' || LOWER($1) || '%') AND ($2='' OR LOWER(email) like '%' || LOWER($2) || '%') AND status=$3 order by create_date desc `, name, email, "ACTIVE")

	defer Close(row)

	row.Scan()

	fmt.Println("Debug", row)
	_, data, err = Scan(row)
	if err != nil {
		fmt.Println("Error", err)
		err = errors.New("contact search failed")
		return
	}

	return
}

func SelectContact(id string) (data [][]string, err error) {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	defer db.Close()

	var row *sql.Rows

	row, err = db.Query(`SELECT id, name,email,mobile,to_char(create_date,'YYYY-MM-DD HH24:MI:SS') FROM contact.details WHERE id=$1`, id)

	defer Close(row)

	row.Scan()

	fmt.Println("Debug", row)
	_, data, err = Scan(row)
	if err != nil {
		fmt.Println("Error", err)
		err = errors.New("contact search failed")
		return
	}

	return
}

func SelectUser(username string) (data [][]string, err error) {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("error", err.Error())
	}
	defer db.Close()

	var row *sql.Rows

	row, err = db.Query(`SELECT username,password FROM contact.user WHERE username=$1`, username)
	if err != nil {
		fmt.Println("error", err.Error())
	}

	defer Close(row)

	row.Scan()

	fmt.Println("Debug", row)
	_, data, err = Scan(row)
	if err != nil {
		fmt.Println("Error", err)
		err = errors.New("contact search failed")
		return
	}

	return
}

func Scan(row *sql.Rows) (cols []string, data [][]string, err error) {
	cols, err = row.Columns()
	if err != nil {
		err = errors.New("DB get fail")
		return
	}

	tmp_byte := make([][]byte, len(cols))
	tmp := make([]interface{}, len(cols))
	for i, _ := range tmp_byte {
		tmp[i] = &tmp_byte[i]
	}
	for row.Next() {
		err = row.Scan(tmp...)
		if err != nil {
			err = errors.New("DB row fail")
			return
		}
		rawResult := make([]string, len(cols))
		for i, _ := range tmp_byte {
			rawResult[i] = string(tmp_byte[i])
		}
		data = append(data, rawResult)
	}
	return
}

func Close(row *sql.Rows) (err error) {
	if row != nil {
		row.Close()
	}
	return
}
