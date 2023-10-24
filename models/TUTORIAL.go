package models

type Tutorial struct {

    TutoId   uint   `gorm:"primary_key"`
    Title string
    Description string
    State bool
    TutoDetail int
}
