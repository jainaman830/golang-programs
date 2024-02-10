package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}

var db *gorm.DB

func init() {
	// Define connection
	var err error
	db, err = gorm.Open(mysql.Open("username:password@tcp(localhost:3306)/test"), &gorm.Config{})
	if err != nil {
		log.Fatalf("error in connection %f", err)
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("error in AutoMigrate %f", err)
	}
}
func Createuser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode("Error in reading body : " + err.Error())
		return
	}
	defer r.Body.Close()
	type InputBody struct {
		Name string
	}
	var input InputBody
	err = json.Unmarshal(body, &input)
	if err != nil {
		json.NewEncoder(w).Encode("Error in input unmarshal : " + err.Error())
		return
	}
	// Create a new record
	insertedData := db.Create(&User{Name: input.Name})
	if insertedData.Error != nil {
		json.NewEncoder(w).Encode("Error in create : " + err.Error())
		return
	}
	json.NewEncoder(w).Encode("User created")
	return
}
func Getuser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode("Error in reading body : " + err.Error())
		return
	}
	defer r.Body.Close()
	type InputBody struct {
		ID int
	}
	var input InputBody
	err = json.Unmarshal(body, &input)
	if err != nil {
		json.NewEncoder(w).Encode("Error in input unmarshal:" + err.Error())
		return
	}
	// Query records
	var user User
	fetchedRecord := db.First(&user, input.ID)
	if fetchedRecord.Error != nil {
		json.NewEncoder(w).Encode("Error in fetch : " + fetchedRecord.Error.Error())
		return
	}
	json.NewEncoder(w).Encode(user)
	return
}
func Updateuser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode("Error in reading body : " + err.Error())
		return
	}
	defer r.Body.Close()
	type InputBody struct {
		ID   int
		Name string
	}
	var input InputBody
	err = json.Unmarshal(body, &input)
	if err != nil {
		json.NewEncoder(w).Encode("Error in input unmarshal:" + err.Error())
		return
	}
	var user User
	fetchedRecord := db.First(&user, input.ID)
	if fetchedRecord.Error != nil {
		json.NewEncoder(w).Encode("Error in fetch : " + fetchedRecord.Error.Error())
		return
	}
	// Update record
	updatedRecord := db.Model(&user).Update("Name", input.Name)
	if updatedRecord.Error != nil {
		json.NewEncoder(w).Encode("Error in update : " + updatedRecord.Error.Error())
		return
	}
	json.NewEncoder(w).Encode("User Updated")
	return
}
func Deleteuser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode("Error in reading body : " + err.Error())
		return
	}
	defer r.Body.Close()
	type InputBody struct {
		ID int
	}
	var input InputBody
	err = json.Unmarshal(body, &input)
	if err != nil {
		json.NewEncoder(w).Encode("Error in input unmarshal:" + err.Error())
		return
	}
	// Delete record
	result := db.Delete(&User{}, input.ID)
	if result.Error != nil {
		json.NewEncoder(w).Encode("Failed to delete record:" + result.Error.Error())
		return
	}
	json.NewEncoder(w).Encode("User Deleted")
	return
}
func StartServer() {
	server := mux.NewRouter().StrictSlash(true)
	server.HandleFunc("/createuser", Createuser).Methods("POST")
	server.HandleFunc("/getuser", Getuser)
	server.HandleFunc("/updateuser", Updateuser).Methods("PUT")
	server.HandleFunc("/deleteuser", Deleteuser).Methods("DELETE")
	http.ListenAndServe(":8080", server)
}

func main() {
	StartServer()
}
