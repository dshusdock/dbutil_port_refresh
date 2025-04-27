package unigydb

import (
	"bytes"
	"database/sql"
	con "dshusdock/go_project/internal/constants"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"log/slog"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Host = "10.201.206.8"
var ChosenDB = "dunkin"

func init() {
	slog.Info("UnigyDB init")
}

func SetHost(host string) {
	Host = host
}

func SetChosenDB(db string) {
	ChosenDB = db
}

func Connect2UnigyDB(ip string) *sql.DB {
	// temporary default value
	if ip == "" {
		ip = "10.201.206.8"
	}
	dbn := "dunkin"

	dbStr := fmt.Sprintf("dunkin:dunkin123@tcp(%s:3306)/%s?multiStatements=true", ip, dbn)
	// Open database connection (SQLite in this example)
	db, err := sql.Open("mysql", dbStr)
	if err != nil {
		slog.Info(err.Error())
	}
	return db
}

func Connect2LocalDB(ip string) *sql.DB {
	// temporary default value
	if ip == "" {
		ip = "db"
	}
	dbn := "udu"

	dbStr := fmt.Sprintf("root:my-secret-pw@tcp(%s:3306)/%s?multiStatements=true", ip, dbn)
	db, err := sql.Open("mysql", dbStr)
	if err != nil {
		slog.Info(err.Error())
	}
	
	return db
}

func GetDBList(qi con.QueryInfo) []string {
	db := Connect2UnigyDB(qi.DBHost)
	defer db.Close()

	rows, err := db.Query("show databases")
	if err != nil {
		slog.Info(err.Error())
	}
	defer rows.Close()
	cols, _ := rows.Columns()

	w := tabwriter.NewWriter(os.Stdout, 0, 2, 1, ' ', 0)
	defer w.Flush()
	sep := []byte("\t")
	//	newLine := []byte("\n")

	w.Write([]byte(strings.Join(cols, "\t") + "\n"))

	row := make([][]byte, len(cols))
	rowPtr := make([]any, len(cols))

	for i := range row {
		rowPtr[i] = &row[i]
	}
	bfr := []string{}
	for rows.Next() {
		_ = rows.Scan(rowPtr...)

		x := bytes.Join(row, sep)
		bfr = append(bfr, string(x))
	}
	db.Close()
	return bfr
}

func PerformQuery(qi con.QueryInfo) *con.UnigyTable {
	slog.Debug("PerformQuery: ", "query", qi.Sql)
	slog.Debug("Target: ", "host", Host)

	var db *sql.DB

	if qi.DbType == "unigy" || qi.DbType == "" {
		db = Connect2UnigyDB(qi.DBHost)

		_, err := db.Exec("USE " + qi.DBName)
		if err != nil {
			slog.Info(err.Error())
			return nil		
		}
	}
	
	if qi.DbType == "local" {
		db = Connect2LocalDB("")
	}
	defer db.Close()

	// Query the database
	rows, err := db.Query(qi.Sql)
	if err != nil {
		slog.Info(err.Error())
		return nil		
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		slog.Info(err.Error())
		return nil		
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
			slog.Info(err.Error())
			return nil		
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
		slog.Info(err.Error())
		return nil		
	}
	return convertRawData2Strings(columns, dataRows)
}

func PerformCountQuery(qi con.QueryInfo) int {
	slog.Debug("PerformCountQuery: ", "query", qi.Sql)
	
	var db *sql.DB

	if qi.DbType == "unigy" || qi.DbType == "" {
		db = Connect2UnigyDB(qi.DBHost)

		_, err := db.Exec("USE " + qi.DBName)
		if err != nil {
			slog.Info(err.Error())
			return 0		
		}
	}
	
	if qi.DbType == "local" {
		db = Connect2LocalDB("")
	}
	defer db.Close()

	var count int
	err := db.QueryRow(qi.Sql).Scan(&count)
	if err != nil {
		slog.Info(err.Error())
		return 0		
	}
	return count
}

func convertRawData2Strings(cols []string, rows []map[string]interface{}) *con.UnigyTable {
	ii:=0
	
	uData := con.NewUnigyTable()
	uData.Columns = cols

	uData.Rows = []con.RowData{}
	uData.TableSize = len(rows)

	for i:=0; i<len(rows); i++ {
		val := rows[i]
		el := con.RowData{}
		// iterate through data columns and convert to string
		for _, v := range cols {					
			if intValue, ok := val[v].(int64); ok {				
				el.Data = append(el.Data, strconv.Itoa(int(intValue))) 
			} else if strValue, ok := val[v].(string); ok {
				el.Data = append(el.Data, strValue)
			} else {
				el.Data = append(el.Data, "NULL")
			}			
			ii++
		}
		uData.Rows = append(uData.Rows, el)
		uData.RowsSlice = append(uData.RowsSlice, el)		
	}
	return uData
}	

func CreateDB(ip string, dbName string) error{
	db := Connect2LocalDB(ip)
	defer db.Close()

	// Create a new database
	_, err := db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		slog.Info("Error creating database", "error", err)
		return err
	}

	slog.Info("Database created successfully")
	return nil
}

