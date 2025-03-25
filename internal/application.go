package internal

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	charmLog "github.com/charmbracelet/log"
	"github.com/gorilla/mux"
)

type App struct {
	logger *charmLog.Logger
	DB *sql.DB
}

func NewApp(logger *charmLog.Logger, db *sql.DB) *App {
	return &App{
		logger: logger,
		DB: db,
	}
}

func (a *App) getAllBreeds(w http.ResponseWriter, r *http.Request) {
	rows, err := a.DB.Query("SELECT * FROM breeds")
	if err != nil {
		http.Error(w, "Failed to fetch breeds", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var breeds []Breed
	for rows.Next() {
		var breed Breed
		if err := rows.Scan(&breed.Id, &breed.Species, &breed.PetSize, &breed.Name, &breed.AverageMaleAdultWeight, &breed.AverageFemaleAdultWeight); err != nil {
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			return
		}
		breeds = append(breeds, breed)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating over rows", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(breeds)
}

func (a *App) getBreed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	query := "SELECT * FROM breeds WHERE id = ?"

	var breed Breed

	err := a.DB.QueryRow(query, id).Scan(
		&breed.Id, &breed.Species, &breed.PetSize, &breed.Name,
		&breed.AverageMaleAdultWeight, &breed.AverageFemaleAdultWeight,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Breed not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to fetch breed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(breed)
}

func validBreed(breed *Breed) error {
	if strings.TrimSpace(breed.Species) == "" {
		return errors.New("species is required")
	} else if breed.Species != "dog" && breed.Species != "cat" {
		return errors.New("species must be either a dog or a cat")
	}
	if strings.TrimSpace(breed.PetSize) == "" {
		return errors.New("petSize is required")
	} else if breed.PetSize != "small" && breed.PetSize != "medium" && breed.PetSize != "tall" {
		return errors.New("petSize must be either small, medium or tall")
	}
	if strings.TrimSpace(breed.Name) == "" {
		return errors.New("name is required")
	} else if len(breed.Name) > 80 {
		return errors.New("name is too long. Cannot exceed 80 characters")
	}

	if breed.AverageMaleAdultWeight <= 0 {
		return errors.New("average_male_adult_weight must be a positive number")
	}
	if breed.AverageFemaleAdultWeight <= 0 {
		return errors.New("average_female_adult_weight must be a positive number")
	}

	return nil
}

func (a *App) createBreed(w http.ResponseWriter, r *http.Request) {
	var newBreed Breed
	err := json.NewDecoder(r.Body).Decode(&newBreed)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = validBreed(&newBreed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO breeds (species, pet_size, name, average_male_adult_weight, average_female_adult_weight)
	          VALUES (?, ?, ?, ?, ?)`
	result, err := a.DB.Exec(query, newBreed.Species, newBreed.PetSize, newBreed.Name, newBreed.AverageMaleAdultWeight, newBreed.AverageFemaleAdultWeight)
	if err != nil {
		fmt.Println("Database Error: ", err)
		http.Error(w, "failed to create breed", http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Error retrieving breed ID", http.StatusInternalServerError)
		return
	}
	newBreed.Id = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBreed)
}

func (a *App) updateBreed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var newBreed Breed
	err := json.NewDecoder(r.Body).Decode(&newBreed)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = validBreed(&newBreed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var exists bool
	err = a.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM breeds WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "Breed not found", http.StatusNotFound)
		return
	}

	query := `UPDATE breeds
		SET species = ?, pet_size = ?, name = ?, average_male_adult_weight = ?, average_female_adult_weight = ?
		WHERE id = ?`

	_, err = a.DB.Exec(query, newBreed.Species, newBreed.PetSize, newBreed.Name, newBreed.AverageMaleAdultWeight, newBreed.AverageFemaleAdultWeight, id)
	if err != nil {
		http.Error(w, "Failed to update breed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newBreed)
}

func (a *App) deleteBreed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	query := `DELETE FROM breeds WHERE id = ?`
	result, err := a.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to delete breed", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to check if breed was deleted", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Breed not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Breed with id: " + id + " deleted successfully")
}

func (a *App) searchBreed(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	species := vars.Get("species")
	weight := vars.Get("weight")

	var args []any
	query := "SELECT * FROM breeds WHERE 1=1"

	if species != "" {
		query += " AND species = ?"
		args = append(args, species)
	}
	if weight != "" {
		_, err := strconv.Atoi(weight)
		if err != nil {
			http.Error(w, "Invalid weight parameter", http.StatusBadRequest)
		}
		query += " AND average_male_adult_weight = ? OR average_female_adult_weight = ?"
		args = append(args, weight, weight)
	}

	rows, err := a.DB.Query(query, args...)
	if err != nil {
		fmt.Println("Error executing query: ", err)
		http.Error(w, "Failed to fetch breeds", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var breeds []Breed
	for rows.Next() {
		var breed Breed
		if err := rows.Scan(&breed.Id, &breed.Species, &breed.PetSize, &breed.Name, &breed.AverageMaleAdultWeight, &breed.AverageFemaleAdultWeight); err != nil {
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			return
		}
		breeds = append(breeds, breed)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating over rows", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(breeds)
}

func (a *App) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/search", a.searchBreed).Methods("GET")
	r.HandleFunc("/", a.getAllBreeds).Methods("GET")
	r.HandleFunc("/{id}", a.getBreed).Methods("GET")
	r.HandleFunc("/", a.createBreed).Methods("POST")
	r.HandleFunc("/{id}", a.updateBreed).Methods("PUT")
	r.HandleFunc("/{id}", a.deleteBreed).Methods("DELETE")
}
