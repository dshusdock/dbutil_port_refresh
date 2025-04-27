package createauditvw

import (
	"dshusdock/go_project/config"
	con "dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/services/session"
	"encoding/gob"
	"log/slog"
	"net/http"
)

type CreateAuditVw struct {
	App *config.AppConfig
}

var AppCreateAuditVw *CreateAuditVw

func init() {
	AppCreateAuditVw = &CreateAuditVw{
		App: nil,
	}
	gob.Register(AppCreateAuditVwData{})
	gob.Register(QueryData{})
}

func (m *CreateAuditVw) RegisterHandler() con.ViewHandler {
	slog.Info("[CreateAuditVw] - RegisterView")
	return &CreateAuditVw{}
}

func (m *CreateAuditVw) HandleRequest(w http.ResponseWriter, event con.AppEvent) any {
	slog.Info("[CreateAuditVw] - HandleRequest")
	var obj AppCreateAuditVwData

	if session.SessionSvc.SessionMgr.Exists(event.Request.Context(), "CreateAuditVw") {
		obj = session.SessionSvc.SessionMgr.Pop(event.Request.Context(), "CreateAuditVw").(AppCreateAuditVwData)
	} else {
		obj = *CreateCreateAuditVwData()	
	}

	obj.ProcessHttpRequest(w, event)	
	session.SessionSvc.SessionMgr.Put(event.Request.Context(), "CreateAuditVw", obj)

	return obj
}
 
///////////////////// Layout View Data //////////////////////

type AppCreateAuditVwData struct {
	Base 			*con.BaseTemplateparams
	Data 			*QueryData
	TextAreaStr 	string
}

type QueryData struct {
	Placeholder 			string
}

func CreateCreateAuditVwData() *AppCreateAuditVwData {
	return &AppCreateAuditVwData{
		Base: nil,
		Data: nil,
		TextAreaStr: "This is a test",
	}
}

func (m *AppCreateAuditVwData) ProcessHttpRequest(w http.ResponseWriter, event con.AppEvent) *AppCreateAuditVwData{
	slog.Info("[CreateAuditVwData] - ProcessHttpRequest")
	
	return m
}

