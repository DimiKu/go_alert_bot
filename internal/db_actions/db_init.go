package db_actions

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

// TODO не забыть потом изменть. Плохая практирка
const (
	host     = "localhost"
	port     = 5436
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
	l          *zap.Logger
}

func NewDBAdminManage(l *zap.Logger) *DBAdminManage {
	return &DBAdminManage{
		user:     user,
		password: password,
		address:  host,
		port:     port,
		l:        l,
	}
}

// Drop database if there is to use every time new db
func (a DBAdminManage) DBCreate(DBName string) string {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, user, password)
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		a.l.Error("can't connect to db", zap.Error(err))
	}

	_, err = conn.Exec(`DROP DATABASE IF EXISTS ` + DBName)
	if err != nil {
		a.l.Error("Failed to drop db", zap.Error(err))
	}

	_, err = conn.Exec("create database " + DBName)
	if err != nil {
		a.l.Error("failed to create db", zap.Error(err), zap.String("DB name", DBName))
	}

	a.ConnString = fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",
		host, port, user, password, DBName)

	return a.ConnString
}

type Storage struct {
	conn   *sql.DB
	l      *zap.Logger
	DBName string
}

func NewStorage(connString string, l *zap.Logger) *Storage {
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		l.Error("Failed to create connection to db", zap.Error(err))
	}
	return &Storage{
		conn: conn,
		l:    l,
	}
}

func (s *Storage) CreateBasicTables() error {
	if err := s.CreateUserTable(); err != nil {
		s.l.Error("can't create user table", zap.Error(err))
		return err
	}

	if err := s.CreateTelegramChatTable(); err != nil {
		s.l.Error("can't create tg chat table", zap.Error(err))
		return err
	}

	if err := s.CreateStdoutChatTable(); err != nil {
		s.l.Error("can't create stdout chat table", zap.Error(err))

		return err
	}

	if err := s.CreateChannelTable(); err != nil {
		s.l.Error("can't create channel table", zap.Error(err))
		return err
	}

	return nil
}

func (s *Storage) CreateUserTable() error {
	_, err := s.conn.Exec(createUsersTable)
	if err != nil {
		return fmt.Errorf("error create table %w", err)
	}

	return nil
}

func (s *Storage) CreateChannelTable() error {
	_, err := s.conn.Exec(createChannelsTable)
	if err != nil {
		return fmt.Errorf("error create table, %w", err)
	}

	return nil
}

func (s *Storage) CreateTelegramChatTable() error {
	_, err := s.conn.Exec(createTelegramChatsTable)
	if err != nil {
		return fmt.Errorf("error create table, %w", err)
	}

	return err
}

func (s *Storage) CreateStdoutChatTable() error {
	_, err := s.conn.Exec(createStdoutChatsTable)
	if err != nil {
		return fmt.Errorf("error create table, %w", err)
	}

	return nil
}
