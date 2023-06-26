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

type DBManage struct {
	user     string
	password string
	address  string
	port     int
	db_name  string
}

func NewDBManage(DBName string) *DBManage {
	return &DBManage{
		user:     user,
		password: password,
		address:  host,
		port:     port,
		db_name:  DBName,
	}
}

func (m *DBManage) connectionStringForCreate() string {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, user, password)
	return connString
}

// TODO как сделать возможные параметры
func (m *DBManage) connectionStringForWork() string {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",
		host, port, user, password, m.db_name)
	return connString
}

func (m *DBManage) CreateDatabase() sql.Result {
	conn, err := sql.Open("postgres", m.connectionStringForCreate())
	defer conn.Close()
	//m.db_name = dbname
	if err != nil {
		fmt.Errorf("failed to create db")
	}
	result, err := conn.Exec(`CREATE DATABASE ` + m.db_name + ` ;`)
	if err != nil {
		fmt.Errorf("failed to create db")
	}
	return result
}

func (m *DBManage) CreateUserTable(tableName string) sql.Result {
	conn, err := sql.Open("postgres", m.connectionStringForWork())
	defer conn.Close()

	if err != nil {
		fmt.Errorf("Failed to create table %s", tableName)
	}
	// resp, err := conn.Exec(`CREATE TABLE "`+tableName+`%s`, fileds)
	resp, err := conn.Exec(`CREATE TABLE "` + tableName + `" (id integer PRIMARY KEY, chat_id integer)`)
	if err != nil {
		fmt.Print("Error create table %s", err)
	}

	fmt.Printf("response %s", resp)
	return resp
}

func (m *DBManage) DropDatabase() {
	conn, err := sql.Open("postgres", m.connectionStringForCreate())
	defer conn.Close()
	if err != nil {
		fmt.Errorf("Failed to drop db")
	}
	_, err = conn.Exec(`DROP DATABASE IF EXISTS ` + m.db_name)
	if err != nil {
		fmt.Errorf("Failed to drop db")
	}
}
