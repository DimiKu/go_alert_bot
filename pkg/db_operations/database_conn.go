package db_operations

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

type DBAdminManage struct {
	user       string
	password   string
	address    string
	port       int
	DBName     string
	ConnString string
}

func NewDBAdminManage() *DBAdminManage {
	return &DBAdminManage{
		user:     user,
		password: password,
		address:  host,
		port:     port,
	}
}

// Drop database if there is to use every time new db

func (a DBAdminManage) DBCreate(DBName string) string {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, user, password)
	conn, err := sql.Open("postgres", connString)
	_, err = conn.Exec(`DROP DATABASE IF EXISTS ` + a.DBName)
	if err != nil {
		fmt.Errorf("Failed to drop db")
	}
	if err != nil {
		fmt.Errorf("failed to create db")
	}
	_, err = conn.Exec(`CREATE DATABASE ` + DBName + ` ;`)
	if err != nil {
		fmt.Errorf("failed to create db")
	}
	a.ConnString = fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",
		host, port, user, password, a.DBName)
	return a.ConnString
}

type Storage struct {
	connString string
	DBName     string
}

func NewStorage(connString string) *Storage {
	return &Storage{
		connString: connString,
	}
}

func (s *Storage) CreateDatabase() sql.Result {
	conn, err := sql.Open("postgres", s.connString)
	if err != nil {
		fmt.Errorf("failed to create db")
	}
	result, err := conn.Exec(`CREATE DATABASE ` + s.DBName + ` ;`)
	if err != nil {
		fmt.Errorf("failed to create db")
	}
	s.CreateUserTable()
	return result
}

func (s *Storage) CreateUserTable() sql.Result {
	conn, err := sql.Open("postgres", s.connString)
	resp, err := conn.Exec(`CREATE TABLE users (id integer PRIMARY KEY, chat_id integer)`)
	if err != nil {
		fmt.Print("Error create table %s", err)
	}

	fmt.Printf("response %s", resp)
	return resp
}
