package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type RegistrationRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Product   string `json:"product"`
	Date      string `json:"date"`
	Time      string `json:"time"`
}

type JsonResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", handleRegister).Methods("POST")
	r.HandleFunc("/registration", serveRegistrationPage).Methods("GET")

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", r)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the homepage!")
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	var requestData RegistrationRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Check for required fields
	if requestData.Email == "" || requestData.FirstName == "" || requestData.LastName == "" ||
		requestData.Phone == "" || requestData.Product == "" || requestData.Date == "" || requestData.Time == "" {
		response := JsonResponse{
			Status:  "400",
			Message: "Missing required fields in the registration data",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Process registration logic here

	response := JsonResponse{
		Status:  "success",
		Message: "Appointment booked successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func serveRegistrationPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Registration.html")
}
