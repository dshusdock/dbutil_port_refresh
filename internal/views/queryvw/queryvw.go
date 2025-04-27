package queryvw

import (
	"dshusdock/go_project/config"
	con "dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/services/session"
	"encoding/gob"
	"fmt"
	"net/http"
)

type QueryVw struct {
	App *config.AppConfig
}

var AppQueryVw *QueryVw

func init() {
	AppQueryVw = &QueryVw{
		App: nil,
	}
	gob.Register(AppQueryVwData{})
	gob.Register(QueryData{})
}

func (m *QueryVw) RegisterHandler() con.ViewHandler {
	return &QueryVw{}
}

func (m *QueryVw) HandleRequest(w http.ResponseWriter, event con.AppEvent) any {
	fmt.Println("[QueryVw] - HandleRequest")
	var obj AppQueryVwData

	if session.SessionSvc.SessionMgr.Exists(event.Request.Context(), "QueryVw") {
		obj = session.SessionSvc.SessionMgr.Pop(event.Request.Context(), "QueryVw").(AppQueryVwData)
	} else {
		obj = *CreateQueryVwData()	
	}

	obj.ProcessHttpRequest(w, event)	
	session.SessionSvc.SessionMgr.Put(event.Request.Context(), "QueryVw", obj)

	return obj
}
 
///////////////////// Layout View Data //////////////////////

type AppQueryVwData struct {
	Base 			*con.BaseTemplateparams
	Data 			*QueryData
	TextAreaStr 	string
}

type QueryData struct {
	Placeholder 			string
}

func CreateQueryVwData() *AppQueryVwData {
	return &AppQueryVwData{
		Base: nil,
		Data: nil,
		TextAreaStr: "select d.name as EnterpriseName, d.domainName as EnterpriseDomainName, c.name as InstanceName, c.domainName as InstanceDomainName, a.id as ZoneId, a.certificateAuthorityIP, a.name, a.softwareVersion, b.replicationEnabled, b.replicationRole from Zone a, InterzoneCommConfigZone b, Instance c, Enterprise d where d.id=c.parentEnterpriseId and c.id=a.parentInstanceId and b.id=a.interzoneCommConfigZoneId",
	}
}

func (m *AppQueryVwData) ProcessHttpRequest(w http.ResponseWriter, event con.AppEvent) *AppQueryVwData{
	fmt.Println("[QueryVwData] - ProcessHttpRequest")
	if event.EventStr == "queryvw_btn_Execute" {
		m.TextAreaStr = event.Request.FormValue("querytext")
	}
	return m
}

