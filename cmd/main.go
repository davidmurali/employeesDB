package main

import (
	config "employeesDB/pkg/config"
	constants "employeesDB/pkg/constants"
	employee "employeesDB/pkg/employee"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	err := config.Init()
	if err != nil {
		logrus.Error(constants.InitError + ": " + err.Error())
	}
	http.ListenAndServe(":9000", employee.CreateEmployeeEndpoints())
}
