package models

type Users struct {
	Userid         string `gorm:"PRIMARY_KEY"`
	Name          string
}