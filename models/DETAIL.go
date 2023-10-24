package models

import "time"


type Detail struct {

    DetailsId   uint   `gorm:"primary_key"`
    Owner int
    DateCreation time.Time
}
