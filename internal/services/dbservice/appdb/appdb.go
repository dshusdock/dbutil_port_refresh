package appdb

import (
	"database/sql"
	con "dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/services/libraries/ssh"
	"dshusdock/go_project/internal/services/messagebus"
	u "dshusdock/go_project/internal/services/utilitysvc"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

// AppDB is the interface that wraps the basic methods for the database

func Connect2DB(ip string) *sql.DB {
	// temporary default value
	if ip == "" {
		ip = "db"
		//ip = "localhost"
	}
	dbn := "udu"

	dbStr := fmt.Sprintf("root:my-secret-pw@tcp(%s:3306)/%s?multiStatements=true", ip, dbn)
	db, err := sql.Open("mysql", dbStr)
	if err != nil {
		log.Fatal(err)
	}
	
	return db
}


func Connect2DBwithConfig(ip string) *sql.DB {	
	var db *sql.DB
	cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   ip+":3306",
        DBName: "recordings",
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
	return db
}

// PerformTableQuery is used to perform a query on a table with a known structure.
// It returns a slice of RowData where each RowData represents a row in the table.
func PerformTableQuery[T any](query string) []con.RowData {
	db := Connect2DB("")
	var tableDef []T

	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var e T

		s := reflect.ValueOf(&e).Elem()

		numCols := s.NumField()
		fmt.Println("numCols is", numCols)
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		if err := rows.Scan(columns...); err != nil {
			log.Fatal(err)
		}
		tableDef = append(tableDef, e)
	}

	var rd = []con.RowData{}

	for i := 0; i < len(tableDef); i++ {
		values := reflect.ValueOf(tableDef[i])

		r := con.RowData{
			Data: nil,
		}
		for ii := 0; ii < values.NumField(); ii++ {
			f := values.Field(ii)
			r.Data = append(r.Data, checkReflect(f))
		}
		rd = append(rd, r)
	}

	return rd
}

// This function is used to perform a generic query on the database without knowing the 
// structure of the table upfront. It returns a slice of maps where each map represents 
// a row in the table and the keys are the column names.
func PerformGenericQuery(query string) ([]map[string]interface{}, []interface{}) {

	db := Connect2DB("")
	defer db.Close()

	// Query the database
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	// Prepare a slice of interface{} to hold values for each column
	columnValues := make([]interface{}, len(columns))
	// Prepare a slice of pointers to hold references to each column value
	columnPointers := make([]interface{}, len(columns))
	for i := range columnValues {
		columnPointers[i] = &columnValues[i]
	}

	dataRows := []map[string]interface{}{}

	// Iterate through rows
	for rows.Next() {
		// Scan the row into the column pointers
		err := rows.Scan(columnPointers...)
		if err != nil {
			log.Fatal(err)
		}

		// Convert the column values into a map
		rowData := make(map[string]interface{})

		for i, colName := range columns {
			// Handle NULL values
			val := columnValues[i]
			if b, ok := val.([]byte); ok {
				rowData[colName] = string(b)			
			} else {
				rowData[colName] = val			
			}
		}		
		dataRows = append(dataRows, rowData)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return dataRows, columnValues
}

func dbExec(query string) sql.Result {	
	db := Connect2DB("")
	defer db.Close()

	// Insert a row
	result, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func dbQuery(query string) *sql.Rows {
	db := Connect2DB("")
	defer db.Close()

	// Insert a row
	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func checkReflect(f reflect.Value) string {
	if f.Kind().String() == "struct" {
		val := f.Interface().(sql.NullString)

		if val.Valid {
			fmt.Println("Valid Data:", val.String)
			return val.String
		} else {
			return "null"
		}
	}
	return f.String()
}

func DBLoaderFromLocal(info con.DBLoaderInfo) error{
	
	slog.Info("DBLoaderFromLocal start...")
	
	messagebus.GetBus().Publish("Event:Poll", "DBLoaderFromLocal start...")

	///////////////////////////////////////////////////////
	// Sanity check to see if the allicat dirs are mounted
	messagebus.GetBus().Publish("Event:Poll", "Checking for allicat...")
	cmd := exec.Command("ls", "-l", "/mnt/allicatnasn1/allicat_inactive")
	output, err := cmd.Output()
    if err != nil {
		handleError("Failed to run command")
		return err
    }

	if strings.Contains(string(output), "a_amer_general") {
		slog.Info("allicat directories mounted")
		messagebus.GetBus().Publish("Event:Poll", "Allicat confirmed...")
	} else {
		handleError("allicat directories not mounted")
		return errors.New("allicat directories not mounted")
	}
	///////////////////////////////////////////////////////
	// Verify that the requested sql file exists
	messagebus.GetBus().Publish("Event:Poll", "Checking for SQL file...")	
	cmd = exec.Command("find", u.GroomAllicatStr(info.FileDir), "-name", info.SQLFileName)	
	err = cmd.Run()
    if err != nil {
		handleError("Failed to run find command")
		return err
    }
	messagebus.GetBus().Publish("Event:Poll", "SQL file found...")
	///////////////////////////////////////////////////////
	// Copy the file to /tmp	
	messagebus.GetBus().Publish("Event:Poll", "Copying file to /tmp...")
	cmd = exec.Command("cp", u.GroomAllicatStr(info.FileDir) + "/" + info.SQLFileName, "/tmp")
	err = cmd.Run()
	//output, err = cmd.Output()
	//fmt.Println(string(output))
	if err != nil {
		handleError("Failed to run copy command")
		return err
	}

	///////////////////////////////////////////////////////
	//gzip -d /tmp/" + p.DBLdrInfo.SQLFileName
	messagebus.GetBus().Publish("Event:Poll", "Uncompressing File...")
	cmd = exec.Command("gzip", "-d", "/tmp/"+info.SQLFileName) 
	err = cmd.Run()
	//output, err = cmd.Output()
	//fmt.Println(string(output))
	if err != nil {
		handleError("Failed to run gzip command")
		return err
	}

	///////////////////////////////////////////////////////
	// Create the database
	messagebus.GetBus().Publish("Event:Poll", "Creating Database...")
	cmd = exec.Command("/usr/bin/mysql", "-h"+info.TargetIP, "-udunkin", "-pdunkin123", "-e", " create database " + info.DBName)
	err = cmd.Run()
	if err != nil {
		handleError("Failed to run create database command")
		return err
	}
	
	///////////////////////////////////////////////////////
	// Load the database
	messagebus.GetBus().Publish("Event:Poll", "Loading Database...")
	err = loadDb2Mysql("dunkin", "dunkin123", info.TargetIP, info.DBName, "/tmp/" + strings.Split(info.SQLFileName, ".")[0] + ".sql")
	if err != nil {
		handleError("Failed to load database")
		return err
	}
	messagebus.GetBus().Publish("Event:Poll", "Database load successful...")
	slog.Info("Database loaded successfully.")

	str := strings.Split(info.SQLFileName, ".")[0] + ".sql"
	cmd = exec.Command("ls", "-lh", "/tmp/"+ str)
	output, err = cmd.Output()
	if err != nil {
		handleError("Failed to run command")
	}
	messagebus.GetBus().Publish("Event:Poll", string(output))
	time.Sleep(5 * time.Second)
	messagebus.GetBus().Publish("Event:Poll", "done")

	return nil
}

func handleError(str string) {
	slog.Error(str)
	messagebus.GetBus().Publish("Event:Poll", str)
	time.Sleep(5 * time.Second)
	messagebus.GetBus().Publish("Event:Poll", "done")
}

func loadDb2Mysql(user string, password string, host string, database string, dumpFile string) error{
    // MySQL credentials and database details
	slog.Debug("loadDb2Target start...")

    // Construct the MySQL import command
    cmd := exec.Command("mysql", "-u"+user, "-p"+password, "-h "+host, database)

    // Open the dump file
    file, err := os.Open(dumpFile)
    if err != nil {
        slog.Error("Error opening file:", err)
        return err
    }
    defer file.Close()

    // Set the command's stdin to the dump file
    cmd.Stdin = file

    // Run the command
    err = cmd.Run()
    if err != nil {
        slog.Error("Error running command:", err)
        return err
    }
	return nil
}

func DBLoaderFromTarget(info con.DBLoaderInfo) error{
	util := u.NewUtilitySvc()
	
	// Log into the database host and create a client
	client := ssh.NewSSHClient(info.SrcIP, info.SrcUser, info.SrcPassword)
	if client == nil {
		slog.Info("Failed to create client")
		return nil
	}

	
	cmd := util.CmdBuilder(u.DBLoadParams{CmdType: "find", DBLdrInfo: info})
	slog.Info(">>>>>>>>>>>Running cmd " + cmd + "<<<<<<<<<<<<<<")
	str, err := ssh.RunCommand(client, cmd)
	if err != nil {
		slog.Info("Failed to run command: ", "Error", err.Error())
		return err
	}

	// Verify that the file exists
	slog.Info(">>>>>>>>>>>Verifying file exists<<<<<<<<<<<<<<")
	if strings.Contains(str, info.SQLFileName) {
		slog.Info("File found")
	} else {
		slog.Info("File not found")
		return nil	
	}
    
	// Copy the file to /tmp
	cmd = util.CmdBuilder(u.DBLoadParams{CmdType: "copy", DBLdrInfo: info})
	slog.Info(">>>>>>>>>>>Copy file to tmp<<<<<<<<<<<<<< " + "\n" + cmd)
	_, err = ssh.RunCommand(client, cmd)
	if err != nil {
		slog.Info("Failed to run command: ", err.Error())
		return err
	}

	// Unzip the file if required
	cmd = util.CmdBuilder(u.DBLoadParams{CmdType: "unzip", DBLdrInfo: info})
	slog.Info(">>>>>>>>>>>Unzip if required<<<<<<<<<<<<<<" + "\n" + cmd)
	if info.UnzipRequired {
		_, err = ssh.RunCommand(client, cmd)
		if err != nil {
			slog.Info("Failed to run command" + err.Error())
		}
	}

	// Create the database
	cmd = util.CmdBuilder(u.DBLoadParams{CmdType: "createdb", DBLdrInfo: info})
	slog.Info(">>>>>>>>>>>Create the database<<<<<<<<<<<<<<" + "\n" + cmd)
	_, err = ssh.RunCommand(client, cmd)
	if err != nil {
		slog.Info("Failed to run command" + err.Error())
	}

	// Load the database
	cmd = util.CmdBuilder(u.DBLoadParams{CmdType: "loaddb", DBLdrInfo: info})
	slog.Info(">>>>>>>>>>>Load the database " + cmd + "<<<<<<<<<<<<<<")
	str, err = ssh.RunCommand(client, cmd)
	if err != nil {
		slog.Info("Failed to run command")
	}

	slog.Info("-->" + str)
	return nil
}