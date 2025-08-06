package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// usin map to store here, no database yet
var items = map[int]string{}
var idCounter = 1

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/items", handleItems)     // GET, POST
	http.HandleFunc("/items/", handleItemByID) // GET, PUT, DELETE

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handle /items (Create and Read All)
func handleItems(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(items)

	case "POST":
		var input struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		items[idCounter] = input.Name
		idCounter++
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Item created")

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handle /items/{id} (Read One, Update, Delete)
func handleItemByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/items/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		name, ok := items[id]
		if !ok {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"name": name})

	case "PUT":
		var input struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		items[id] = input.Name
		fmt.Fprintf(w, "Item updated")

	case "DELETE":
		if _, ok := items[id]; !ok {
			http.NotFound(w, r)
			return
		}
		delete(items, id)
		fmt.Fprintf(w, "Item deleted")

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
