package main

import "errors"

type Account struct {
	fullName string
	userName string
}

type Administrator struct {
	projects map[string]Project
	Account
	ProjectMember // comp
}

// Administrator reciever functions
func (administrator Administrator) addUser(project *Project, user *User) error {
	// WIP
	return errors.New("WIP")
}

func (administrator Administrator) removerUser(project *Project, user *User) error {
	// WIP
	return errors.New("WIP")
}

func (administrator Administrator) deleteProject(project *Project) error {
	// WIP
	return errors.New("WIP")
}

func (administrator Administrator) changeUserRole(project *Project, user *User) error {
	// WIP
	return errors.New("WIP")
}

func (administrator *Administrator) login() error {
	// WIP
	return errors.New("WIP")
}

func (administrator *Administrator) logout() error {
	// WIP
	return errors.New("WIP")
}

type ProjectManager struct {
	managedProjects map[string]Project // projekt som förvaltas av projektledaren
	totalTime       uint64             // total totalt tid arbetat av projektledaren
	Account
	ProjectMember // comp
}

// ProjectManager reciever functions
func (projectManager ProjectManager) signReport(timeReport *TimeReport, user User) error {
	// WIP
	return errors.New("WIP")
}

func (projectManager ProjectManager) unsignReport(timeReport *TimeReport, user User) error {
	// WIP
	return errors.New("WIP")
}

func (projectManager ProjectManager) getAllReports(project *Project) ([]TimeReport, error) {
	// WIP
	return project.timeReports, errors.New("WIP")
}

func (projectManager ProjectManager) assignRole(user *User, project *Project, newRole string) error {
	// WIP
	return errors.New("WIP")
}

func (projectManager ProjectManager) removeMember(project *Project, user *User) error {
	// WIP
	return errors.New("WIP")
}

func (projectManager ProjectManager) getTotalTime(project *Project) (uint64, error) {
	// WIP
	return 0, errors.New("WIP")
}

func (projectManager *ProjectManager) login() error {
	// WIP
	return errors.New("WIP")
}

func (projectManager *ProjectManager) logout() error {
	// WIP
	return errors.New("WIP")
}

type ProjectMember struct {
	timereports []TimeReport
	role        string // ?????
	Account            // comp
}

// User reciever functions

// function used to create a time report, returning a *TimeReport is questionable
func (ProjectMember *ProjectMember) createTimeReport(Project *Project) (*TimeReport, error) {
	// WIP
	return &TimeReport{}, errors.New("WIP")
}

func (ProjectMember ProjectMember) getTimeReport(timereports *[]TimeReport) (*TimeReport, error) {
	// WIP
	return &TimeReport{}, errors.New("WIP")
}

func (ProjectMember *ProjectMember) editTimeReport(timereport *TimeReport) {
	// timereport.editReport()
	// WIP
}

func (projectUser ProjectMember) deleteTimeReport(timeReport *TimeReport) error { // Ska bara project manager kunna göra detta? fråga ledarna!
	// WIP
	return errors.New("WIP")
}

func (projectUser *ProjectMember) login() error {
	// WIP
	return errors.New("WIP")
}

func (projectUser *ProjectMember) logout() error {
	// WIP
	return errors.New("WIP")
}

type User interface {
	login()
	logout()
}
