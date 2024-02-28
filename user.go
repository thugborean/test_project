package main

import "errors"

type Account struct {
	fullName string
	userName string
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
	projects    map[string]Project
	ProjectUser // comp
}

// Administrator reciever functions
func (administrator Administrator) addUser(user *User) error {
	// WIP
	return errors.New("WIP")
}

func (administrator Administrator) removerUser(user *User) error {
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

	ProjectMember // comp
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

func (projectManager ProjectManager) giveUserRole(user *User, project *Project, newRole string) error {
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

type ProjectMember struct {
	timereports []TimeReport
	role        string // ?????
	Account            // comp
}

// User reciever functions

// function used to create a time report, returning a *TimeReport is questionable
func (ProjectMember ProjectMember) createTimeReport(Project *Project) *TimeReport {
	// WIP
	return &TimeReport{}
}

func (ProjectMember ProjectMember) viewTimeReport(timereports *[]TimeReport) *TimeReport {
	// WIP
	return &TimeReport{}
}

func (ProjectMember ProjectMember) editTimeReport(timereport *TimeReport) {
	// timereport.editReport()
	// WIP
}

func (projectUser ProjectMember) deleteTimeReport(*TimeReport) error { // Ska bara project manager kunna göra detta? fråga ledarna!
	// WIP
	return errors.New("WIP")
}

func (ProjectMember ProjectMember) printRole() string {
	return ProjectMember.role
}

type User interface {
	ProjectMember
	ProjectManager
	Administrator
}
