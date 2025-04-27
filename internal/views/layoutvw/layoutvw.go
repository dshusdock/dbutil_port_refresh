package layoutvw

import (
	"dshusdock/go_project/config"
	con "dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/services/session"
	"encoding/gob"
	"log/slog"
	"net/http"
)

type LayoutVw struct {
	App *config.AppConfig
}

var COMP_NAME = "[layoutvw] - "

var AppLayoutVw *LayoutVw

func init() {
	AppLayoutVw = &LayoutVw{
		App: nil,
	}
	gob.Register(LayoutVwData{})
}

func (m *LayoutVw) RegisterHandler() con.ViewHandler {
	slog.Info( COMP_NAME + "RegisterView")
	return &LayoutVw{}
}

func (m *LayoutVw) HandleRequest(w http.ResponseWriter, event con.AppEvent) any {
	slog.Info(COMP_NAME + "HandleRequest")
	var obj LayoutVwData

	if session.SessionSvc.SessionMgr.Exists(event.Request.Context(), "layoutvw") {
		obj = session.SessionSvc.SessionMgr.Pop(event.Request.Context(), "layoutvw").(LayoutVwData)
	} else {
		obj = *CreateLayoutVwData()	
	}

	obj.ProcessHttpRequest(w, event)	
	session.SessionSvc.SessionMgr.Put(event.Request.Context(), "layoutvw", obj)

	return obj
}
 
///////////////////// Layout View Data //////////////////////

type LayoutVwData struct {
	Base *con.BaseTemplateparams
	Data any
	View int
}

type AppLytVwData struct {
	Lbl string
}

func CreateLayoutVwData() *LayoutVwData {
	return &LayoutVwData{
		Base: nil,
		Data: nil,
		View: 0,
	}
}

func (m *LayoutVwData) ProcessHttpRequest(w http.ResponseWriter, event con.AppEvent) *LayoutVwData{
	return m
}

func (m *LayoutVwData) ProcessMBusRequest(w http.ResponseWriter, r *http.Request) {}
