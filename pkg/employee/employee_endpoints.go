package employee

import (
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEmployeeEndpoints() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/createemployee", CreateEmployee).Methods(http.MethodPost)
	router.HandleFunc("/getemployeebyid", GetEmployeeByID).Methods(http.MethodGet)
	router.HandleFunc("/updateemployee", UpdateEmployee).Methods(http.MethodPut)
	router.HandleFunc("/deleteemployee", DeleteEmployee).Methods(http.MethodDelete)
	return router
}
