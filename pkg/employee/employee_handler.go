package employee

import (
	"database/sql"
	config "employeesDB/pkg/config"
	constants "employeesDB/pkg/constants"
	database "employeesDB/pkg/database"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	dbPort, _ := strconv.Atoi(config.GetMySQLDBPort())
	db, err := database.GetDb(config.GetMySQLDBUsername(), config.GetMySQLDBPassword(), config.GetMySQLDBProtocol(), config.GetMySQLDBHost(), dbPort, config.GetMySQLDBName())
	if err != nil {
		err = fmt.Errorf(constants.DBOpenError + ": " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var emp Employee
	err = json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		err = fmt.Errorf(constants.DecodeError + ": " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// config.SelectDatabase(db, "employees")
	_, err = db.Exec("INSERT INTO employees (name, position, salary) VALUES (?, ?, ?)", emp.Name, emp.Position, emp.Salary)
	if err != nil {
		err = fmt.Errorf(constants.DBInsertError + ": " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {

	dbPort, _ := strconv.Atoi(config.GetMySQLDBPort())
	db, err := database.GetDb(config.GetMySQLDBUsername(), config.GetMySQLDBPassword(), config.GetMySQLDBProtocol(), config.GetMySQLDBHost(), dbPort, config.GetMySQLDBName())
	if err != nil {
		err = fmt.Errorf(constants.DBOpenError + ": " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := r.URL.Query().Get("id")
	var emp Employee
	err = db.QueryRow("SELECT id, name, position, salary FROM employees WHERE id = ?", id).Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary)
	if err == sql.ErrNoRows {
		err = fmt.Errorf(constants.SqlNoRowsError + ": " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if err != nil {
		err = fmt.Errorf(constants.DBGetError + ": " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(emp)

}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {

	dbPort, _ := strconv.Atoi(config.GetMySQLDBPort())
	db, err := database.GetDb(config.GetMySQLDBUsername(), config.GetMySQLDBPassword(), config.GetMySQLDBProtocol(), config.GetMySQLDBHost(), dbPort, config.GetMySQLDBName())
	if err != nil {
		err = fmt.Errorf(constants.DBOpenError + ": " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := r.URL.Query().Get("id")
	var emp Employee
	err = json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		err = fmt.Errorf(constants.DecodeError + ": " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = db.Exec("UPDATE employees SET name = ?, position = ?, salary = ? WHERE id = ?", emp.Name, emp.Position, emp.Salary, id)
	if err == sql.ErrNoRows {
		err = fmt.Errorf(constants.SqlNoRowsError + ": " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if err != nil {
		err = fmt.Errorf(constants.DBUpdateError + ": " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	dbPort, _ := strconv.Atoi(config.GetMySQLDBPort())
	db, err := database.GetDb(config.GetMySQLDBUsername(), config.GetMySQLDBPassword(), config.GetMySQLDBProtocol(), config.GetMySQLDBHost(), dbPort, config.GetMySQLDBName())
	if err != nil {
		err = fmt.Errorf(constants.DBOpenError + ": " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := r.URL.Query().Get("id")
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM employees WHERE id = ?", id).Scan(&count)
	if err != nil {
		err = fmt.Errorf(constants.DBGetError + ": " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if count == 0 {
		err = fmt.Errorf("Employee with ID %s not found", id)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = db.Exec("DELETE FROM employees WHERE id = ?", id)
	if err != nil {
		err = fmt.Errorf(constants.DBDeleteError + ": " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
