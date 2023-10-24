package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/C0C045/Prueba-Tecnica/db"
	"github.com/C0C045/Prueba-Tecnica/models"
	"github.com/gorilla/mux"
)

func GetDetail(w http.ResponseWriter, r *http.Request){
    var details []models.Detail
    db.DB.Find(&details)
    log.Println(details)
    json.NewEncoder(w).Encode(&details)
}

func PostDetail(w http.ResponseWriter, r *http.Request){
    var detail models.Detail
    var user models.User
    pr := mux.Vars(r)
    
    json.NewDecoder(r.Body).Decode(&detail)
    db.DB.First(&user, pr["id"])
    
    if user.UserId==0{
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("User doesnt exist"))
        return
    }
    detail.Owner = int(user.UserId)

    /*Guardar en BD*/
    userCreated := db.DB.Create(&detail)
    err := userCreated.Error

    if err != nil{
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
    }

    json.NewEncoder(w).Encode(&detail)
}
