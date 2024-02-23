package main

import "errors"

type Account struct {
	fullName string
	userName string
	password string
}

// Account reciever functions
func (account Account) login() error {
	// WIP
	return errors.New("WIP")
}

func (account Account) logout() error {
	// WIP
	return errors.New("WIP")
}

type Administrator struct {
	projects map[string]Project
	// deletedData ???
	user User // comp
}

// Administrator reciever functions
func (administrator Administrator) addUser(user User) error {
	// WIP
	return errors.New("WIP")
}

func (administrator Administrator) removerUser(user User) error {
	// WIP
	return errors.New("WIP")
}

func (administrator Administrator) deleteProject(project *Project) {
	// WIP
}

func (administrator Administrator) changeUserRole() {
	// WIP
}
func (administrator Administrator) recoverDeletedData() error { // Hur fan ska detta funka dock haha...
	// WIP
	return errors.New("WIP")
}

type ProjectManager struct {
	managedProjects map[string]Project // projekt som förvaltas av projektledaren
	totalTime       uint64             // total totalt tid arbetat av projektledaren

	user User // comp, för att kunna skapa och
}

// ProjectManager reciever functions
func (projectManager ProjectManager) signReport(timeReport *TimeReport, user User) {
	// WIP
}

func (projectManager ProjectManager) unsignReport(timeReport *TimeReport, user User) {
	// WIP
}

func (projectManager ProjectManager) viewAllReports(project *Project) ([]TimeReport, error) {
	// WIP
	return project.timeReports, errors.New("WIP")
}

func (projectManager ProjectManager) giveUserRole(user *User, newRole string) error {
	// WIP
	return errors.New("WIP")
}

func (projectManager ProjectManager) removeUserFromProject(user *User, project *Project) error {
	// WIP
	return errors.New("WIP")
}

func (projectManager ProjectManager) viewTotalTime(project *Project) uint64 {
	// WIP
	return 0
}

type User struct {
	account     Account // comp
	timereports []TimeReport
	role        string // ?????
}

// User reciever functions

// function used to create a time report
func (user User) createTimeReport() *TimeReport {
	// WIP
	return &TimeReport{}
}

func (user User) viewTimeReport(timereports *[]TimeReport) *TimeReport {
	// WIP
	return &TimeReport{}
}

func (user User) editTimeReport(timereport *TimeReport) {
	// timereport.editReport()
	// WIP
}

func (user User) deleteTimeReport(*TimeReport) error { // Ska bara project manager kunna göra detta? fråga ledarna!
	// WIP
	return errors.New("WIP")
}
