package main

import (
	"fmt"
	"log"
	"net/http"
)

// Request
// Reponse
// fun(req,res){}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Qt internship bank")
	fmt.Println("Endpoint Hit: homePage")
}

func Deposit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Account deposit page!")
	fmt.Println("Endpoint Hit: depositPage")
}
func Credit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Account Credit page!")
	fmt.Println("Endpoint Hit: creditPage")
}
func Balance(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Account balance page!")
	fmt.Println("Endpoint Hit: balancePage")
}
func New_account(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to new Account registration page!")
	fmt.Println("Endpoint Hit: newAccountPage")
}
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Page!")
	fmt.Println("Endpoint Hit: loginPage")
}
func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout Page!")
	fmt.Println("Endpoint Hit: logoutPage")
}

func handleRequests() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/deposit", Deposit)
	http.HandleFunc("/credit", Credit)
	http.HandleFunc("/balance", Balance)
	http.HandleFunc("/new_account", New_account)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/logout", Logout)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
