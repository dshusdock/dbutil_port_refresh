package dbsidenavvw

import (
	"dshusdock/go_project/config"
	con "dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/services/session"
	"encoding/gob"
	"log/slog"
	"net/http"
)

type DBSideNavVw struct {
	App *config.AppConfig
}

var AppDBSideNavVw *DBSideNavVw

func init() {
	AppDBSideNavVw = &DBSideNavVw{
		App: nil,
	}
	gob.Register(AppDBSideNavVwData{})
	gob.Register(HeaderItem{})
}

func (m *DBSideNavVw) RegisterHandler() con.ViewHandler {
	slog.Info("[DBSideNavVw] - RegisterView")
	return &DBSideNavVw{}
}

func (m *DBSideNavVw) HandleRequest(w http.ResponseWriter, event con.AppEvent) any {
	slog.Info("[DBSideNavVw] - HandleRequest")
	var obj AppDBSideNavVwData

	if session.SessionSvc.SessionMgr.Exists(event.Request.Context(), "DBSideNavVw") {
		obj = session.SessionSvc.SessionMgr.Pop(event.Request.Context(), "DBSideNavVw").(AppDBSideNavVwData)
	} else {
		obj = *CreateDBSideNavVwData()	
	}

	obj.ProcessHttpRequest(w, event)	
	session.SessionSvc.SessionMgr.Put(event.Request.Context(), "DBSideNavVw", obj)

	return obj
}
 
///////////////////// Layout View Data //////////////////////

type AppDBSideNavVwData struct {
	Base 		*con.BaseTemplateparams
	Data 		[]ButtonData
	AuditData 	[]string
	HomeMenu 	bool
	AuditMenu 	bool
	DBList		[]string
	DBHost 		string
	SelectedDB 	string
	MenuType 	uint8
}

type ElementId struct {
	Name 		string
	Label 		string
}

type ButtonData struct {
	Category 	string
	ElId		[]ElementId
	Open 		bool
	Icon 		string
}

type HeaderItem struct {
	ID 				string
	Lbl     		string
	Class			string
}

type DropDownData struct {
	Class string
	Style string
}

var AUDIT_MENU1 = []string{
	"Create Audit",
	"Run Audit",
}

var AUDIT_MENU2 = []string{
	"Audit Menu",
}

var AUDIT_MENU3 = []string{
	"Audit Menu",
	"Audit Results",
}

const ( 
	MENU_1 = 1
	MENU_2 = 2
	MENU_3 = 3
)

var ICON_MAP = make (map[string]string)


func CreateDBSideNavVwData() *AppDBSideNavVwData {
	ICON_MAP["System"] = "fa-solid fa-server"
	ICON_MAP["User"] = "fa-solid fa-user"
	ICON_MAP["Recording"] = "fa-solid fa-microphone"
	ICON_MAP["Button"] = "fa-solid fa-stop"
	ICON_MAP["Resource AOR"] = "fa-solid fa-phone"
	ICON_MAP["Open Connection"] = "fa-solid fa-volume-high"
	ICON_MAP["Line"] = "fa-solid fa-line"
	ICON_MAP["Zone"] = "fa-solid fa-server"

	return &AppDBSideNavVwData{
		Base: 		&con.BaseTemplateparams{},
		Data: 		createButtonData(),
		AuditData: 	AUDIT_MENU1,
		HomeMenu: 	true,
		AuditMenu: 	false,
		DBList: 	[]string{},
		DBHost: 	"",
		SelectedDB: "",
		MenuType: 	MENU_1,
	}
}

func (m *AppDBSideNavVwData) ProcessHttpRequest(w http.ResponseWriter, event con.AppEvent) *AppDBSideNavVwData{
	slog.Info("[DBSideNavVwData] - ProcessHttpRequest")
	
	if event.EventId == con.EVENT_STARTUP {
		return m
	}

	switch event.EventStr {
	case "dbsidenavvw_btn_System":
		for i, _ := range m.Data {
			if m.Data[i].Category == "System" {
				m.Data[i].Open = !m.Data[i].Open
			}
		}
		m.HomeMenu = true
	}


	return m
}

func createButtonData() []ButtonData {
	var bd = []ButtonData{}
	found := false
	
	// Range over the JsonData and map the Labels to the Categories
	for _, v := range con.ButtonMap { // took out ii for print statements
		found = false
		if len(bd)==0 {
			bd = append(bd, ButtonData{Category: v.Category, ElId: []ElementId{{Name: v.Name, Label: v.Label }, }, Open: false, Icon: ""})
			continue
		} 

		for i, _ := range bd {
			if bd[i].Category == v.Category {	
				bd[i].ElId = append(bd[i].ElId, ElementId{ Name: v.Name, Label: v.Label })
				found = true
				break
			} 
		}

		if !found {
			bd = append(bd, ButtonData{Category: v.Category, ElId: []ElementId{{Name: v.Name, Label: v.Label }, }, Open: false, Icon: ""})
		}
	}

	// Populate the Icons
	for i, _ := range bd {
		bd[i].Icon = ICON_MAP[bd[i].Category]
	}

	return bd
}


