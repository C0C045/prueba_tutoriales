package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/C0C045/Prueba-Tecnica/db"
	"github.com/C0C045/Prueba-Tecnica/models"
	"github.com/gorilla/mux"
)

func GetTutorials(w http.ResponseWriter, r *http.Request){
    var tutorials []models.Tutorial
    db.DB.Find(&tutorials)
    log.Println(tutorials)
    json.NewEncoder(w).Encode(&tutorials)
}

func DeleteTutorial(w http.ResponseWriter, r *http.Request){
    var tutorial models.Tutorial
    pr := mux.Vars(r)
    db.DB.First(&tutorial, pr["id"])
    
    if tutorial.TutoId == 0{
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("Tutorial doesnt exist"))
        return
    }

    db.DB.Unscoped().Delete(&tutorial)
}

func PostTutorial(w http.ResponseWriter, r *http.Request){
    var tutorial models.Tutorial
    var detail models.Detail
    pr := mux.Vars(r)
    
    json.NewDecoder(r.Body).Decode(&tutorial)
    db.DB.First(&detail, pr["id"])
    
    if detail.DetailsId == 0{
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("Detail doesnt exist"))
        return
    }
    tutorial.TutoDetail = int(detail.DetailsId)

    /*Guardar en BD*/
    userCreated := db.DB.Create(&tutorial)
    err := userCreated.Error

    if err != nil{
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
    }

    json.NewEncoder(w).Encode(&tutorial)
}

func UpdateTutorial(w http.ResponseWriter, r *http.Request){
    var tutorial models.Tutorial
    pr := mux.Vars(r)
    
    db.DB.First(&tutorial, pr["id"])
    
    if tutorial.TutoId == 0{
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("Tutorial doesnt exist"))
        return
    }
    
    newtuto := tutorial 

    json.NewDecoder(r.Body).Decode(&newtuto)
    tutorial.Title = newtuto.Title
    tutorial.Description = newtuto.Description
    tutorial.State = newtuto.State

    /*Guardar en BD*/
    userUpdated := db.DB.Save(&tutorial)
    err := userUpdated.Error

    if err != nil{
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
    }

    json.NewEncoder(w).Encode(&newtuto)
}
