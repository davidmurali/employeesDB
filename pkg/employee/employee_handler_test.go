package employee_test

import (
	database "employeesDB/pkg/database"
	mock_database "employeesDB/pkg/mocks"
	"fmt"
	"os"
	"testing"

	"github.com/golang/mock/gomock"

	_ "github.com/go-sql-driver/mysql"
)

func TestCreateEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// Create a mock DB instance
	mockDB := mock_database.NewMockDBInterface(ctrl)

	// Prepare test data
	emp := &database.Employee{
		ID:       1,
		Name:     "John",
		Position: "Developer",
		Salary:   50000,
	}

	// Set expectation on the mock
	mockDB.EXPECT().CreateEmployee(emp).Return(nil)

	// Call the function to be tested
	err := mockDB.CreateEmployee(emp)

	// Check if the error is nil (indicating success)
	if err != nil {
		t.Errorf("CreateEmployee returned unexpected error: %v", err)
	}
}
func TestGetEmployeeByID(t *testing.T) {
	// Create a mock DB instance
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock_database.NewMockDBInterface(ctrl)

	// Prepare test data
	empID := 1
	expectedEmp := &database.Employee{
		ID:       empID,
		Name:     "John",
		Position: "Developer",
		Salary:   50000,
	}

	// Set expectation on the mock
	mockDB.EXPECT().GetEmployeeByID(empID).Return(expectedEmp, nil)

	// Call the function to be tested
	emp, err := mockDB.GetEmployeeByID(empID)

	// Check if the error is nil (indicating success)
	if err != nil {
		t.Errorf("GetEmployeeByID returned unexpected error: %v", err)
	}

	// Check if the returned employee matches the expected employee
	if emp.ID != expectedEmp.ID || emp.Name != expectedEmp.Name || emp.Position != expectedEmp.Position || emp.Salary != expectedEmp.Salary {
		t.Errorf("GetEmployeeByID returned unexpected employee: got %v, want %v", emp, expectedEmp)
	}
	fmt.Println(emp)
}

func TestUpdateEmployee(t *testing.T) {
	// Create a mock DB instance
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock_database.NewMockDBInterface(ctrl)

	// Prepare test data
	emp := &database.Employee{
		ID:       1,
		Name:     "John",
		Position: "Manager",
		Salary:   50000,
	}

	// Set expectation on the mock
	mockDB.EXPECT().UpdateEmployee(emp).Return(emp, nil)

	// Call the function to be tested
	empUpdated, err := mockDB.UpdateEmployee(emp)

	// Check if the error is nil (indicating success)
	if err != nil {
		t.Errorf("UpdateEmployee returned unexpected error: %v", err)
	}

	// Check if the employee returned matches the expected employee
	if empUpdated != emp {
		t.Errorf("UpdateEmployee returned unexpected employee: got %+v, want %+v", empUpdated, emp)
	}
}

func TestDeleteEmployee(t *testing.T) {
	// Create a mock DB instance
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock_database.NewMockDBInterface(ctrl)

	// Prepare test data
	empID := 1

	// Set expectation on the mock
	mockDB.EXPECT().DeleteEmployee(empID).Return(nil, nil)

	// Call the function to be tested
	_, err := mockDB.DeleteEmployee(empID)

	// Check if the error is nil (indicating success)
	if err != nil {
		t.Errorf("DeleteEmployeeByID returned unexpected error: %v", err)
	}
}

// Write similar test functions for other CRUD operations

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}
