package myutil

import (
	"errors"
	"fmt"
	"time"
	"unsafe"
)

type Employee struct {
	id          string
	firstName   string
	lastName    string
	age         int32
	joiningDate time.Time
}

type Manager struct {
	// e Employee
	Employee
	team    string
	workers []string
}

func (e Employee) getEmployeeFullName() {
	fmt.Println("Full name of the employee is:", e.firstName, e.lastName)
}

func newEmployee(fname, lname string, age int32) (*Employee, error) {
	if fname == "" || lname == "" || age <= 18 {
		return nil, errors.New("Please enter valid input")
	}
	return &Employee{
		"000000",
		fname,
		lname,
		age,
		time.Now(),
	}, nil
}

func newManager(fname, lname string, age int32, team string, emps ...string) (*Manager, error) {
	if fname == "" || lname == "" || age <= 18 {
		return nil, errors.New("Please enter valid input")
	}
	e, _ := newEmployee(fname, lname, age)
	return &Manager{
		*e,
		team,
		emps,
	}, nil
}

func Structs() {
	// STRUCTS
	e, _ := newEmployee("John", "Doe", 32)
	e.getEmployeeFullName()
	adm, _ := newManager("Jack", "Smith", 45, "BlackOps")
	adm.getEmployeeFullName()
	var emp Employee
	var m Manager
	fmt.Printf("%d %d\n", unsafe.Sizeof(emp), unsafe.Sizeof(m))
	fmt.Println()
}
