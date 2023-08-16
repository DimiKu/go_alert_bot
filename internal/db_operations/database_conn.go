package db_operations

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// TODO не забыть потом изменть. Плохая практирка
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
	s.CreateChatTable()
	s.CreateChannelTable()
	return result
}

func (s *Storage) CreateUserTable() sql.Result {
	resp, err := s.conn.Exec(`CREATE TABLE users (user_id integer PRIMARY KEY, chat_id integer)`)
	if err != nil {
		fmt.Print("Error create table %s", err)
	}

	return resp
}

func (s *Storage) CreateChannelTable() sql.Result {
	resp, err := s.conn.Exec(`CREATE TABLE channels (user_id integer PRIMARY KEY, chat_id bigint, channel_link bigint)`)
	if err != nil {
		fmt.Print("Error create table %s", err)
	}

	return resp
}

func (s *Storage) CreateChatTable() sql.Result {
	resp, err := s.conn.Exec(`CREATE TABLE chats (user_id integer PRIMARY KEY, chat_id integer)`)
	if err != nil {
		fmt.Print("Error create table %s", err)
	}

	return resp
}
