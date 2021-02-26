package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"pall.com/controllers"
	"pall.com/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "h0prsvmv."
	dbname   = "pall_com"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	checkErr(err)
	defer us.Close()
	us.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/testcookie", usersC.TestCookie).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	log.Println("Server is started!")
	http.ListenAndServe(":3000", r)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
