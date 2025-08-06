package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "modernc.org/sqlite" // Go SQLite driver
)

var db *sql.DB

func main() {
	var err error

	// Open or create the SQLite database file
	db, err = sql.Open("sqlite", "./items.db")
	if err != nil {
		log.Fatal("❌ Database open error:", err)
	}

	// Create the table if it doesn't exist
	createTable := `
	CREATE TABLE IF NOT EXISTS Items (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT NOT NULL
	);`
	if _, err = db.Exec(createTable); err != nil {
		log.Fatal("❌ Table creation error:", err)
	}

	//fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	// http.Handle("/", http.RedirectHandler("/index.html", http.StatusFound)) // error here because was calling index("/") while looking static
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// API routes
	http.HandleFunc("/items", handleItems)
	http.HandleFunc("/items/", handleItemByID)

	fmt.Println("🚀 Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleItems(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rows, err := db.Query("SELECT ID, Name FROM Items")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		result := make(map[int]string)
		for rows.Next() {
			var id int
			var name string
			rows.Scan(&id, &name)
			result[id] = name
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)

	case "POST":
		var input struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		_, err := db.Exec("INSERT INTO Items (Name) VALUES (?)", input.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Item created")

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleItemByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/items/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		var name string
		err := db.QueryRow("SELECT Name FROM Items WHERE ID = ?", id).Scan(&name)
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"name": name})

	case "PUT":
		var input struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		res, err := db.Exec("UPDATE Items SET Name = ? WHERE ID = ?", input.Name, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		count, _ := res.RowsAffected()
		if count == 0 {
			http.NotFound(w, r)
			return
		}

		fmt.Fprint(w, "Item updated")

	case "DELETE":
		res, err := db.Exec("DELETE FROM Items WHERE ID = ?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		count, _ := res.RowsAffected()
		if count == 0 {
			http.NotFound(w, r)
			return
		}

		fmt.Fprint(w, "Item deleted")

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
