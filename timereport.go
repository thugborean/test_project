package main

import "errors"

type TimeReport struct {
	reportId     int
	projectName  string
	userName     string
	userRole     string
	reportDate   string
	hoursWorked  int
	reportStatus string // Example "draft", "signed", "unsigned"

}

// TimeReport reciever functions

func (timeReport TimeReport) creatReport() {
	// WIP
}

func (timeReport TimeReport) editReport() {
	// WIP
}

func (timeReport TimeReport) signReport() {
	// WIP
}

func (timeReport TimeReport) unsignReport() {
	// WIP
}

func (timeReport TimeReport) viewTotalTime() uint64 {
	// WIP
	return 0
}

func (timeReport TimeReport) printReport() string {
	// WIP
	return ""
}

type Project struct {
	timeReports    []TimeReport
	projectName    string
	projectMembers map[string]string
	projectRoles   map[string]string
	timeReport     TimeReport // comp
}

// Project reciever functions // WIP WIP WIP

func (project Project) addMember(member string, role string) error {
	//WIP
	return errors.New("WIP")
}

func (project Project) removeMember(member string) error {
	//WIP
	return errors.New("WIP")
}
