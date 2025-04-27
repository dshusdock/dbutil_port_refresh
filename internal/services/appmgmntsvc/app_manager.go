package appmgmntsvc

import (
	"context"
	con "dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/services/session"
	"encoding/gob"
	"log/slog"
)

type UserInfo struct {
	Username 	string
	LoggedIn 	bool
	DBHost 		string
	DBName 		string
}

type AppManagerData struct {
	Users 			[]UserInfo
	AppMode 		int
	ActiveDBHost 	string
	ActiveDBName 	string
}

func NewAppManager() *AppManagerData {
	return &AppManagerData{
		Users: 	[]UserInfo{},
		AppMode: con.HOME,
	}
}

func init() {
	gob.Register(AppManagerData{})
	slog.Debug("[AppManager]: AppManagerData init")
}

func CreateAppManager(ctx context.Context) *AppManagerData {
	slog.Debug("[AppManager]: Creating AppManager")
	obj := NewAppManager()
	session.SessionSvc.SessionMgr.Put(ctx, "AppManager", obj)
	return obj
}

func GetAppManager(ctx context.Context) *AppManagerData {
	var obj AppManagerData
	if session.SessionSvc.SessionMgr.Exists(ctx, "AppManager") {
		obj = session.SessionSvc.SessionMgr.Get(ctx, "AppManager").(AppManagerData)
		return &obj
	}
	return nil
}

func GetCurrentLogonState(ctx context.Context) con.Bits {
	obj := session.SessionSvc.SessionMgr.Get(ctx, "AppManager")
	if obj == nil {
		return con.BIT_NONE
	}
	return con.BIT_LoggedIn	
}

func SetAppMode(ctx context.Context, mode int) {
	var obj AppManagerData

	slog.Debug("[AppManager]: Setting AppMode", "mode", mode)

	if session.SessionSvc.SessionMgr.Exists(ctx, "AppManager") {
		obj = session.SessionSvc.SessionMgr.Get(ctx, "AppManager").(AppManagerData)
		obj.AppMode = mode
		session.SessionSvc.SessionMgr.Put(ctx, "AppManager", obj)
	} else {
		slog.Error("[AppManager]: AppManager session object not found")
	}
}

func GetAppMode(ctx context.Context) int {	
	obj := session.SessionSvc.SessionMgr.Get(ctx, "AppManager")
	if obj == nil {
		return con.NONE
	}
	return obj.(AppManagerData).AppMode
}

func GetUserAuditList(ctx context.Context) []string {
	obj := session.SessionSvc.SessionMgr.Get(ctx, "AppManager")
	if obj == nil {
		return nil
	}
	return []string{"User1", "User2", "User3"}
}

func SetActiveDBHostAndName(ctx context.Context, host string, name string) {
	var obj AppManagerData

	slog.Debug("[AppManager]: Setting ActiveDBHostAndName", "host", host, "name", name)

	if session.SessionSvc.SessionMgr.Exists(ctx, "AppManager") {
		obj = session.SessionSvc.SessionMgr.Get(ctx, "AppManager").(AppManagerData)
		obj.ActiveDBHost = host
		obj.ActiveDBName = name
		session.SessionSvc.SessionMgr.Put(ctx, "AppManager", obj)
	} else {
		slog.Error("[AppManager]: AppManager session object not found")
	}
}


func GetActiveDBHostAndName(ctx context.Context) (string, string) {	
	obj := session.SessionSvc.SessionMgr.Get(ctx, "AppManager")
	if obj == nil {
		return "", ""
	}
	return obj.(AppManagerData).ActiveDBHost, obj.(AppManagerData).ActiveDBName
}






