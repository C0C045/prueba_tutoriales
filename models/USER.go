package models


type User struct {

    UserId   uint   `gorm:"primary_key"`
    FullName string
    Username string
}
