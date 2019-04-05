package main

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"time"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Username string
	Password string
	EMail    string
}

type UserJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
	EMail    string `json:"email"`
}

func main() {
	log.Println("Start Software")

	//Migrate and Connect
	db = migrateConnect()

	// Start REST API
	buildRestAPI()
}

func buildRestAPI() {
	router := mux.NewRouter()
	log.Println("Serving at localhost:80...")

	// User Authentification Routes
	router.HandleFunc("/api/user", getUser).Methods("GET")
	router.HandleFunc("/api/user", createUser).Methods("POST")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func migrateConnect() *gorm.DB {
	// Wait until database is succesfully up
	time.Sleep(5 * time.Second)

	// Connect to Database
	var globalDB *gorm.DB
	var err error

	//for connected {
	globalDB, err = gorm.Open("mysql", "root:g240iowejgadvsijfgwrgafd@tcp(db:3306)/golangstarter?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Println("failed to connect database ... retry")
		log.Println(err)
	} else {
		log.Println("Successfully connect to database")
		// Run Migrations
		errors := globalDB.AutoMigrate(&User{}).Error
		if errors != nil {
			log.Println("Got errors")
			log.Println(errors)
		}
	}
	return globalDB
}

func getUser(w http.ResponseWriter, r *http.Request) {
	// Build Response
	w.Header().Add("Content-Type", "application/json")

	// Get all users from Database
	var user []User
	errDB := db.Find(&user).Error
	if errDB != nil {
		err := json.NewEncoder(w).Encode("Database Error")
		if err != nil {
			log.Println(err)
		}
	} else {
		// Send Response
		err := json.NewEncoder(w).Encode(user)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// Parse User Authentification
	var jsonuser UserJson
	err := json.NewDecoder(r.Body).Decode(&jsonuser)
	if err != nil {
		log.Println("ERROR DECODING JSON")
		log.Println(err.Error())
	}

	var user User
	user.EMail = jsonuser.EMail
	user.Password = jsonuser.Password
	user.Username = jsonuser.Username

	if db.Where(&User{EMail: user.EMail}).First(&user).RecordNotFound() {
		// Create the USer
		err := db.Create(&user).Error
		if err != nil {
			log.Println(err.Error())
			log.Println("Failed to create user in database")
			err := json.NewEncoder(w).Encode("Failed to create user in database")
			if err != nil {
				log.Println(err.Error())
			}
		}
	} else {
		log.Println("Failed to create user, e-mail already exists")
		err := json.NewEncoder(w).Encode("Failed to create user, e-mail already exists")
		if err != nil {
			log.Println(err.Error())
		}
	}
}
