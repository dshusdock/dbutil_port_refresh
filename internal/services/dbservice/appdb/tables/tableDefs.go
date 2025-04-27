package tables

import (
	"dshusdock/go_project/internal/services/dbservice/appdb"
	"dshusdock/go_project/internal/services/utilitysvc"
	"fmt"
)

type PBIs struct {
	Id          	int    		`json:"id"`
	Assignee		string 		`json:"assignee"`
	AssigneeGroup 	string 		`json:"assigneeGroup"`
	Client 			string 		`json:"client"`
	ProblemId 		string 		`json:"problemId"`
	SubmiteDate 	string		`json:"submiteDate"`
	//SubmiteDate 	time.Time		`json:"submiteDate"`
	Status 			string 		`json:"status"`
	Summary 		string 		`json:"summary"`
	LastModified 	string 		`json:"lastModified"`
	//LastModified 	time.Time 		`json:"lastModified"`
}

// InsertPBI inserts a new PBI row into the PBIs table
func InsertPBI(pbi PBIs) (int64, error) {
	db := appdb.Connect2DB("127.0.0.1")
	defer db.Close()

	// Insert a row
	result, err := db.Exec("INSERT INTO PBIs (assignee, assigneeGroup, client, problemId, submiteDate, status, summary, lastModified) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", 
		pbi.Assignee, pbi.AssigneeGroup, pbi.Client, pbi.ProblemId, pbi.SubmiteDate, pbi.Status, pbi.Summary, pbi.LastModified)

	if err != nil {
		return 0, fmt.Errorf("InsertPBI: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("InsertPBI: %v", err)
	}
	return id, nil
}

type Query struct {
	Id 				uint16
	Category 		string
	Created 		string
	Name 			string
	Query 			string
	SubCategory 	string
	Type 			string
	User 			string
	LastModified 	string
}

// InsertQuery inserts a new Query row into the Queries table
func InsertQuery(query Query) (int64, error) {
	db := appdb.Connect2DB("127.0.0.1")
	defer db.Close()

	// Insert a row
	result, err := db.Exec("INSERT INTO Query (category, created, name, query, subCategory, type, user, lastModified) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		query.Category, query.Created, query.Name, query.Query, query.SubCategory, query.Type, query.User, query.LastModified)
	if err != nil {
		return 0, fmt.Errorf("InsertPBI: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("InsertPBI: %v", err)
	}

	return id, nil
}

// GetQuery retrieves a query from the Queries table
func GetQueryBySubCategory(_type string) Query {
	db := appdb.Connect2DB("")
	defer db.Close()

	// Query the database
	row := db.QueryRow("SELECT id, category, created, name, query, subCategory, type, user, lastModified FROM Query WHERE subCategory = ?", _type)

	var qry Query
	err := row.Scan(&qry.Id, &qry.Category, &qry.Created, &qry.Name, &qry.Query, &qry.SubCategory, &qry.Type, &qry.User, &qry.LastModified)
	if err != nil {
		return Query{}
	}
	return qry
}

func GetQueryCountBySubCategory(_type string) int {
	db := appdb.Connect2DB("")
	defer db.Close()

	// Query the database
	var count int
	_= db.QueryRow("SELECT COUNT(*) FROM Query WHERE category like ?", _type).Scan(&count)
	return count
}


type UserData struct {
	FirstName string
	LastName string
	Email string
	Username string
	Password []byte
}

// InsertUser inserts a new User row into the Users table
func InsertUserInfo(user UserData) error {
	db := appdb.Connect2DB("")
	defer db.Close()

	// Insert a row
	_, err := db.Exec("INSERT INTO UserInfo (firstname, lastname, email, username, password) VALUES (?, ?, ?, ?, ?)",
		user.FirstName, user.LastName, user.Email, user.Username, user.Password)
	if err != nil {
		return fmt.Errorf("InsertUser: %v", err)
	}
	return nil
}

// GetUserInfo retrieves a user's information from the Users table
func GetUserInfo(username string) (UserData, error) {
	db := appdb.Connect2DB("")
	defer db.Close()

	// Query the database
	row := db.QueryRow("SELECT firstName, lastName, email, username, password FROM UserInfo WHERE username = ?", username)

	var user UserData
	err := row.Scan(&user.FirstName, &user.LastName, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return UserData{}, fmt.Errorf("GetUserInfo: %v", err)
	}
	return user, nil
}

// UpdateUserInfo updates a user's information in the Users table
func UpdateUserInfo(user UserData) error {
	db := appdb.Connect2DB("")
	defer db.Close()

	// Update the row
	_, err := db.Exec("UPDATE UserInfo SET firstName = ?, lastName = ?, email = ?, password = ? WHERE username = ?",
		user.FirstName, user.LastName, user.Email, user.Password, user.Username)
	if err != nil {
		return fmt.Errorf("UpdateUserInfo: %v", err)
	}
	return nil
}	

// DeleteUserInfo deletes a user's information from the Users table
func DeleteUserInfo(username string) error {
	db := appdb.Connect2DB("")
	defer db.Close()

	// Delete the row
	_, err := db.Exec("DELETE FROM UserInfo WHERE username = ?", username)
	if err != nil {
		return fmt.Errorf("DeleteUserInfo: %v", err)
	}
	return nil
}

// validateUserInfo checks if the user's information is valid
func Check4Username(username string) bool {
	db := appdb.Connect2DB("")
	defer db.Close()
	var user string

	// Query the database
	row := db.QueryRow("SELECT username FROM UserInfo WHERE username = ?", username)

	err := row.Scan(&user)
	if err != nil {
		return false
	} 
	if user != username {
		return false
	}
	return true
}

func ValidatePassword(username string, password string) bool {
	db := appdb.Connect2DB("")
	defer db.Close()
	var pass string

	// Query the database
	row := db.QueryRow("SELECT password FROM UserInfo WHERE username = ?", username)

	err := row.Scan(&pass)
	if err != nil {
		return false
	} 
	pw, _ := utilitysvc.DecryptValue([]byte(pass))
	return pw == password
}