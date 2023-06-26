package entities

type User struct {
	id     int `db:"id"`
	chatId int `db:"chat"`
}
