package routes

import (
	"encoding/json"
	"net/http"

	"github.com/C0C045/Prueba-Tecnica/db"
	"github.com/C0C045/Prueba-Tecnica/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request){
    var users []models.User
    db.DB.Find(&users)
    json.NewEncoder(w).Encode(&users)
}

func PostUser(w http.ResponseWriter, r *http.Request){
    var user models.User
    
    json.NewDecoder(r.Body).Decode(&user)
    /*Guardar en BD*/
    userCreated := db.DB.Create(&user)
    err := userCreated.Error

    if err != nil{
        w.WriteHeader(http.StatusBadRequest) //400
        w.Write([]byte(err.Error()))
    }

    json.NewEncoder(w).Encode(&user)
}
