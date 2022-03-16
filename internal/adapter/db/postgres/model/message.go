package model

type Message struct {
	Id   string `gorm:"primary_key;not null;column:id"`
	Code string `gorm:"column:code"`
	Text string `gorm:"column:text"`
}
