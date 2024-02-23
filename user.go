package main

type Account struct {
	fullName string
	userName string
	password string
}

// Account reciever functions
func (account Account) login() {
	// WIP
}

func (account Account) logout() {
	// WIP
}

type Administrator struct {
	projects map[string]Project
	// deletedData ???
	user User // comp
}

// Administrator reciever functions
func (administrator Administrator) addUser(user User) {
	// WIP
}

func (administrator Administrator) removerUser(user User) {
	// WIP
}

func (administrator Administrator) deleteProject(project *Project) {
	// WIP
}

func (administrator Administrator) changeUserRole() {
	// WIP
}
func (administrator Administrator) recoverDeletedData() { // Hur fan ska detta funka dock haha...
	// WIP
}

type ProjectManager struct {
	managedProjects map[string]Project // projekt som förvaltas av projektledaren
	totalTime       uint64             // total totalt tid arbetat av projektledaren

	user User // comp
}

type User struct {
	account     Account // comp
	timereports []TimeReport
}

// User reciever functions

// function used to create a time report
func (administrator Administrator) createProject() *TimeReport {
	// WIP
	return &TimeReport{}
}
