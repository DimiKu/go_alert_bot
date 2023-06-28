package db_operations

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5434
	user     = "postgres"
	password = "postgres"
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

	_, err = conn.Exec(`DROP DATABASE IF EXISTS ` + DBName)
	if err != nil {
		fmt.Errorf("Failed to drop db")
	}

	_, err = conn.Exec(`CREATE DATABASE ` + DBName + ` ;`)
	fmt.Print(DBName)
	if err != nil {
		fmt.Errorf("failed to create db")
	}

	a.ConnString = fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",
		host, port, user, password, DBName)

	fmt.Println(a.ConnString)
	return a.ConnString
}

type Storage struct {
	conn   *sql.DB
	DBName string
}

func NewStorage(connString string) *Storage {
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		fmt.Errorf("Failed to craate connection to db")
	}
	return &Storage{
		conn: conn,
	}
}

func (s *Storage) CreateDatabase() sql.Result {
	result, err := s.conn.Exec(`CREATE DATABASE ` + s.DBName + ` ;`)
	if err != nil {
		fmt.Errorf("failed to create db")
	}
	s.CreateUserTable()
	return result
}

func (s *Storage) CreateUserTable() sql.Result {
	resp, err := s.conn.Exec(`CREATE TABLE users (id integer PRIMARY KEY, chat_id integer)`)
	if err != nil {
		fmt.Print("Error create table %s", err)
	}

	fmt.Printf("response %s", resp)
	return resp
}

func (s Storage) CreateNewUser(ChatId, UserId int) sql.Result {
	q := `INSERT INTO users (chat_id, user_id) values ($1, $2)`
	resp, err := s.conn.Exec(q, ChatId, UserId)
	if err != nil {
		fmt.Errorf("Failed add new user")
	}
	fmt.Print(resp)
	return resp
}
