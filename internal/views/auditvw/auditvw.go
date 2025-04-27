package auditvw

import (
	"dshusdock/go_project/config"
	con "dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/services/dbservice/appdb/tables"
	"dshusdock/go_project/internal/services/session"
	"encoding/gob"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

type AuditVw struct {
	App *config.AppConfig
}

var AppAuditVw *AuditVw

func init() {
	AppAuditVw = &AuditVw{
		App: nil,
	}
	gob.Register(AppAuditVwData{})
	gob.Register(AuditData{})
}

func (m *AuditVw) RegisterHandler() con.ViewHandler {
	slog.Info("[AuditVw] - RegisterView")
	return &AuditVw{}
}

func (m *AuditVw) HandleRequest(w http.ResponseWriter, event con.AppEvent) any {
	slog.Info("[AuditVw] - HandleRequest")
	
	var obj AppAuditVwData

	if session.SessionSvc.SessionMgr.Exists(event.Request.Context(), "AuditVw") {
		obj = session.SessionSvc.SessionMgr.Pop(event.Request.Context(), "AuditVw").(AppAuditVwData)
	} else {
		obj = *CreateAuditVwData()	
	}

	obj.ProcessHttpRequest(w, event)	
	session.SessionSvc.SessionMgr.Put(event.Request.Context(), "AuditVw", obj)

	return obj
}
 
///////////////////// Layout View Data //////////////////////

type AppAuditVwData struct {
	Base 			*con.BaseTemplateparams
	Data 			[]con.AuditItem
	TextAreaStr 	string
	OldView 		bool
	Modal 			bool
	MasterCheck 	bool
}

type AuditData struct {
	Placeholder 	string
}

type DropDownData struct {
	Class string
	Style string
}

func CreateAuditVwData() *AppAuditVwData {
	return &AppAuditVwData{
		Base: 			nil,
		Data: 			PopulateAuditData(),
		TextAreaStr: 	"",		
		OldView: 		false,
		Modal: 			false,
		MasterCheck: 	true,
		
	}
}

var EvtStr = ""

func (m *AppAuditVwData) ProcessHttpRequest(w http.ResponseWriter, event con.AppEvent) *AppAuditVwData{
	slog.Info("[AuditVwData] - ProcessHttpRequest")
	slog.Info("[AuditVwData] - ProcessHttpRequest", "Data", event.Data)
	
	if event.EventStr == "control-hdrvw_button_CreateAudit" {
		slog.Info("[AuditVwData] - ProcessHttpRequest - control-hdrvw_button_CreateAudit")
		m.Modal = true
	}

	// Normalize the event string
	if strings.Contains(event.EventStr, "auditvw_Issue_") {
		EvtStr = "auditvw_checkbox_issue"
	} else {
		EvtStr = event.EventStr
	}

	switch EvtStr {			
		case "auditvw_checkbox_issue":
			slog.Info("[AuditVwData] - ProcessHttpRequest - auditvw_checkbox_audit")
			for idx, item := range m.Data {
				if item.Issue == event.Src {
					if event.Data == "true"{
						m.Data[idx].Checked = true
					} else {
						m.Data[idx].Checked = false
					}				
				}
			}
		case "auditvw_selectall_toggle":
			slog.Info("[AuditVwData] - ProcessHttpRequest - auditvw_btn_SetAll")
			var flag = false

			if event.Data == "true" {
				flag = true
				m.MasterCheck = true
			} else {
				flag = false
				m.MasterCheck = false
			}
			// Set the master checkbox
			for idx, _ := range m.Data {				
				m.Data[idx].Checked = flag				
			}
		case "auditvw_btn-clearall":
			slog.Info("[AuditVwData] - ProcessHttpRequest - auditvw_btn_ClearAll")
			for idx, _ := range m.Data {
				m.Data[idx].Checked = false
			}
		case "auditvw_btn_Submit":
			data := event.Request.PostForm
			obj := tables.Query{}
			
			obj.Category = data.Get("category")
			obj.SubCategory = data.Get("subcategory")
			obj.Name = data.Get("issue_name")
			obj.Query = data.Get("issue_query")
			obj.Type = data.Get("type")
			obj.User = session.SessionSvc.SessionMgr.GetString(event.Request.Context(), "user")
			obj.Created = time.Now().Format("2006-01-02 15:04:05")
			obj.LastModified = time.Now().Format("2006-01-02 15:04:05")
	
			_, err := tables.InsertQuery(obj)
			if err != nil {
				slog.Error("[AuditVwData] - ProcessHttpRequest - InsertQuery", "Error", err)
			}
		default:
			slog.Info("[AuditVwData] - ProcessHttpRequest - default")	
	}	
	return m
}

func PopulateAuditData() []con.AuditItem{
	slog.Info("[AuditVwData] - PopulateAuditData")
	var items []con.AuditItem
	//var x con.AuditItem

	cnt := tables.GetQueryCountBySubCategory("Issue")

	fmt.Println("Count: ", cnt)

	for i := 0; i < cnt; i++ {
		x := con.AuditItem{}
		str := fmt.Sprintf("IssueDetail_%02d", i+1)
		qry := tables.GetQueryBySubCategory(str)
		
		x.Issue = fmt.Sprintf("Issue_%02d", i+1)
		x.Query = qry.Query
		x.ElId = qry.Id
		x.Checked = true
		x.Category = qry.Category	
		x.IType = qry.Type
		str = fmt.Sprintf("Issue_%02d", i+1)
		qry = tables.GetQueryBySubCategory(str)
		x.IssueDef = qry.Query
		x.IssueName = qry.Name

		items = append(items, x)
	}
	return items
}


