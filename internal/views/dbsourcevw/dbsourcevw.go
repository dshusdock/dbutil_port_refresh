package dbsourcevw

import (
	"dshusdock/go_project/config"
	con "dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/services/appmgmntsvc"
	"dshusdock/go_project/internal/services/dbservice/unigydb"
	"dshusdock/go_project/internal/services/session"
	"encoding/gob"
	"log/slog"
	"net/http"
)

type DBSourceVw struct {
	App *config.AppConfig
}

var AppDBSourceVw *DBSourceVw

func init() {
	AppDBSourceVw = &DBSourceVw{
		App: nil,
	}
	gob.Register(AppDBSourceVwData{})
}

func (m *DBSourceVw) RegisterHandler() con.ViewHandler {
	slog.Info("[DBSourceVw] - RegisterView")
	return &DBSourceVw{}
}

func (m *DBSourceVw) HandleRequest(w http.ResponseWriter, event con.AppEvent) any {
	slog.Info("[DBSourceVw] - HandleRequest")
	var obj AppDBSourceVwData

	if session.SessionSvc.SessionMgr.Exists(event.Request.Context(), "DBSourceVw") {
		obj = session.SessionSvc.SessionMgr.Pop(event.Request.Context(), "DBSourceVw").(AppDBSourceVwData)
	} else {
		obj = *CreateDBSourceVwData()	
	}

	obj.ProcessHttpRequest(w, event)	
	session.SessionSvc.SessionMgr.Put(event.Request.Context(), "DBSourceVw", obj)

	return obj
}
 
///////////////////// Layout View Data //////////////////////

type AppDBSourceVwData struct {
	DBList			[]string
	DBHost 			string
	SelectedDB 		string
	DBConfirmStr 	string
}

func CreateDBSourceVwData() *AppDBSourceVwData {
	return &AppDBSourceVwData{
		DBList: []string{},
		DBHost: "",
		SelectedDB: "",	
		DBConfirmStr: "Select Server/Database",			
	}
}

type SomeStruct struct{
	Result string	`json:"result"`
}

func (m *AppDBSourceVwData) ProcessHttpRequest(w http.ResponseWriter, event con.AppEvent) *AppDBSourceVwData{
	slog.Info("[DBSourceVwData] - ProcessHttpRequest")
	slog.Info("event.Src: ", "value", event.Src)

	if event.Src == "target-select" {
		slog.Info("[DBSourceVwData] - Target-Select")
		ip := event.Request.FormValue("targetserver")
		slog.Info("IP: ",  "value", ip)
		unigydb.Host = ip
		m.DBHost = ip
		m.DBList = unigydb.GetDBList(con.QueryInfo{Sql: "show databases", DBHost: ip, DBName: "", DbType: "unigy",}, )
		slog.Info("DBList: ",  "value", m.DBList)
		m.SelectedDB = "Select Database..."
	} 

	if event.Src == "db-select" {
		slog.Info("[DBSourceVwData] - DB-Select")
		db := event.Request.FormValue("selecteddb")
		m.SelectedDB = db

		count := unigydb.PerformCountQuery(con.QueryInfo{
			Sql: "SELECT COUNT(*) FROM Zone;",
			DBHost: m.DBHost,
			DBName: db,
			DbType: "unigy",}) 
		
		slog.Info("Count: ",  "value", count)
		if count > 0 {
			m.DBConfirmStr = db + " - Available"
			appmgmntsvc.SetActiveDBHostAndName(event.Request.Context(), m.DBHost, db)
		} else {
			m.DBConfirmStr = db + " - Invalid"
		}
	}
	
	return m
}


