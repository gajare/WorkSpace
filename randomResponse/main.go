package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Define the User struct
type User struct {
	Name    string                 `json:"name"`
	Salary  float64                `json:"salary"`
	Address map[string]interface{} `json:"address"`
}

// Handler function to respond with a User JSON
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Example user with full address details
	user := map[string]interface{}{
		"name":   "John Doe",
		"salary": 50000.0,
		"address": map[string]interface{}{
			"city":    "New York",
			"pincode": "10001",
			"nearby":  "Central Park",
		},
		"extra1": "unexpected value 1",
		"extra2": 42,
	}

	// Example user with partial address details
	user2 := map[string]interface{}{
		"name":   "Jane Smith",
		"salary": 60000.0,
		"address": map[string]interface{}{
			"city":    "Los Angeles",
			"pincode": "90001",
		},
	}

	// Example user with only city in address
	user3 := map[string]interface{}{
		"name":   "Alice Johnson",
		"salary": 70000.0,
		"address": map[string]interface{}{
			"city": "San Francisco",
		},
	}

	response := []map[string]interface{}{user, user2, user3}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Function to handle dynamic fields in JSON response
func handleDynamicFields(data []byte) ([]User, error) {
	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		return users, err
	}
	return users, nil
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUserHandler).Methods("GET")

	http.Handle("/", r)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
