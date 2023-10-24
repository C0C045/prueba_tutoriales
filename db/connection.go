package db

import (
    "log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbconfig = "host=localhost user=postgres password=C0C045 dbname=bd_tutos port=5432"
var DB *gorm.DB

func DataBaseConnection(){
    var err error
    DB, err = gorm.Open(postgres.Open(dbconfig), &gorm.Config{})    
    if err != nil{
        log.Fatal(err)
    }else{
        log.Println("Data Base Connected")
    }
}
