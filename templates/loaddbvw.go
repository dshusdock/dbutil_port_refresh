package loaddbvw

import (
	"dshusdock/go_project/config"
	con "dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/services/session"
	"encoding/gob"
	"fmt"
	"net/http"
)

type LoadDBVw struct {
	App *config.AppConfig
}

var AppLoadDBVw *LoadDBVw

func init() {
	AppLoadDBVw = &LoadDBVw{
		App: nil,
	}
	gob.Register(AppLoadDBVwData{})
	gob.Register(LoadDB{})
}

func (m *LoadDBVw) RegisterHandler() con.ViewHandler {
	return &LoadDBVw{}
}

func (m *LoadDBVw) HandleRequest(w http.ResponseWriter, event con.AppEvent) any {
	fmt.Println("[LoadDBVw] - HandleRequest")
	var obj AppLoadDBVwData

	if session.SessionSvc.SessionMgr.Exists(event.Request.Context(), "LoadDBVw") {
		obj = session.SessionSvc.SessionMgr.Pop(event.Request.Context(), "LoadDBVw").(AppLoadDBVwData)
	} else {
		obj = *CreateLoadDBVwData()	
	}

	obj.ProcessHttpRequest(w, event)	
	session.SessionSvc.SessionMgr.Put(event.Request.Context(), "LoadDBVw", obj)

	return obj
}
 
///////////////////// Layout View Data //////////////////////

type AppLoadDBVwData struct {
	Base *con.BaseTemplateparams
	Data []LoadDB
}

type LoadDB struct {
	ID 				string
	Lbl     		string
	Class			string
}

func CreateLoadDBVwData() *AppLoadDBVwData {
	return &AppLoadDBVwData{
		Base: nil,
		Data: nil,
	}
}

func (m *AppLoadDBVwData) ProcessHttpRequest(w http.ResponseWriter, event con.AppEvent) *AppLoadDBVwData{
	fmt.Println("[LoadDBVwData] - ProcessHttpRequest")
	fmt.Println("event.Src: ", "value", event.Src)

	return m
}


