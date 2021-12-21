package main

import (
	"encoding/json"
	"log"
	"os"
	"github.com/joho/godotenv"
	"math/rand"
	
	"net/http"
	"strconv"
    "github.com/gorilla/mux"
	
)



 type Person struct{
   Id string `json:"personid"`
	Mail string `json:"email"` 
   Password string `json:"password"` 
   Role string `json:"role"`
}


var persons []Person

func getPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}
func loginCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	type Result struct{
		Status string `json:"status"`
		Role string `json:"role"`
	}
	res:=Result{Status:"False",Role:""}
	for _, item := range persons {
		if item.Mail== params["email"]  && item.Password==params["password"]{
			res.Role=item.Role
			res.Status="True"
			json.NewEncoder(w).Encode(res)
			return
		}
	}
	json.NewEncoder(w).Encode(res)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) 
	for _, item := range persons {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// Add new book
func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Person
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	persons = append(persons, book)
	json.NewEncoder(w).Encode(book)
}

// Update book
func updatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range persons {
		if item.Id == params["id"] {
			persons = append(persons[:index], persons[index+1:]...)
			var book Person
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.Id = params["id"]
			persons = append(persons, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}
func home(w http.ResponseWriter, r *http.Request){
type Result struct{
		Content string `json:"content"`

	}
	res:=Result{Content:"use postman for this api /people get req gives people int the db.."}

	json.NewEncoder(w).Encode(res)
}
// Delete book
func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range persons {
		if item.Id == params["id"] {
			persons = append(persons[:index], persons[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(persons)
}

func login(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	type Result struct{
		Status string `json:"status"`
		Role string `json:"role"`
	}
	res:=Result{Status:"False",Role:""}
	for _, item := range persons {

		if item.Mail== params["email"]{

			res.Role=item.Role
			res.Status="True"
			json.NewEncoder(w).Encode(res)
			return
		}
	}
	json.NewEncoder(w).Encode(res)
}
// Main function
func main() {
	// Init router
	r := mux.NewRouter()

		godotenv.Load()
	    port := os.Getenv("PORT")
 persons=append(persons,Person{Id:"1",Mail:"18pa1a0561@gmail.com",Password:"Ilovemydad7@",Role: "student"})
 persons=append(persons,Person{Id:"2",Mail:"vamsijavvadi@gmail.com",Password:"Ilovemydad7@",Role: "teacher"})
 	persons=append(persons,Person{Id:"3",Mail:"pramilacheruku001@gmail.com",Password:"Ilovemydad7@",Role: "student"})
 	
 
	// Route handles & endpoints
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/people", getPeople).Methods("GET")
	r.HandleFunc("/login/email/{email}",login).Methods("GET")
		r.HandleFunc("/login/{email}/{password}",loginCheck).Methods("GET")
			
	r.HandleFunc("/people/{id}", getPerson).Methods("GET")
	r.HandleFunc("/people", createPerson).Methods("POST")
	r.HandleFunc("/people/{id}", updatePerson).Methods("PUT")
	r.HandleFunc("/people/{id}", deletePerson).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":"+port, r))
}
