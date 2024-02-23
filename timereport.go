package main

import "errors"

type TimeReport struct {
	reportId     string
	projectName  string
	userName     string
	userRole     string
	reportDate   string
	hoursWorked  uint32
	isSigned     bool
	reportStatus string // Example "draft", "signed", "unsigned"

}

// TimeReport reciever functions

func (timeReport *TimeReport) editReport() { // func is called in editTimeReport
	// WIP
}

func (timeReport *TimeReport) signReport() { // is called from ProjectManager
	// WIP
}

func (timeReport *TimeReport) unsignReport() { // is called from ProjectManager
	// WIP
}

func (timeReport TimeReport) printReport() string {
	// WIP
	return ""
}

type Project struct {
	timeReports    []TimeReport
	projectName    string
	projectMembers map[string]User
	projectRoles   map[string]string
	timeReport     TimeReport // comp
}

// Project reciever functions // WIP WIP WIP

func (project *Project) addMember(projectmembers *map[string]User, role string) error {
	//WIP
	return errors.New("WIP")
}

func (project *Project) removeMember(member string) error {
	//WIP
	return errors.New("WIP")
}

func (project Project) viewTotalTime() uint64 {
	// WIP
	return 0
}
