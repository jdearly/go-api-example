package app

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) Setup() {
	app.Router.
		Methods("GET").
		Path("/users").
		HandlerFunc(app.getAllFunc)
	app.Router.
		Methods("POST").
		Path("/users").
		HandlerFunc(app.postFunc)
	app.Router.Methods("GET").
		Path("/users/{id}").
		HandlerFunc(app.getByIDFunc)
	app.Router.Methods("PUT").
		Path("/users/{id}").
		HandlerFunc(app.putFunc)
	app.Router.Methods("DELETE").
		Path("/users/{id}").
		HandlerFunc(app.deleteFunc)
}

func (app *App) getByIDFunc(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
	}

	dbdata := &User{}
	err := app.Database.QueryRow("SELECT id, name FROM `example_table` WHERE id = ?", id).Scan(&dbdata.ID, &dbdata.Name)

	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dbdata); err != nil {
		panic(err)
	}
}

func (app *App) getAllFunc(w http.ResponseWriter, r *http.Request) {

	var users []User
	result, err := app.Database.Query("SELECT id, name FROM `example_table`")

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		dbdata := &User{}
		err := result.Scan(&dbdata.ID, &dbdata.Name)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, *dbdata)
	}

	json.NewEncoder(w).Encode(users)
}

func (app *App) putFunc(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
	}

	var dbdata User
	decodeErr := json.NewDecoder(r.Body).Decode(&dbdata)
	if decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	_, err := app.Database.Exec("UPDATE `example_table` SET name = ? WHERE id = ?", dbdata.Name, id)

	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dbdata); err != nil {
		panic(err)
	}
}

func (app *App) deleteFunc(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
	}

	_, err := app.Database.Exec("DELETE FROM `example_table` WHERE id = ?", id)

	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusOK)
}

func (app *App) postFunc(w http.ResponseWriter, r *http.Request) {

	var dbdata User
	decodeErr := json.NewDecoder(r.Body).Decode(&dbdata)
	if decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	_, err := app.Database.Exec("INSERT IGNORE INTO `example_table` (id, name) VALUES(?, ?)", dbdata.ID, dbdata.Name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}
