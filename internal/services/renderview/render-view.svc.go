package renderview

import (
	"dshusdock/go_project/config"
	con "dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/render"
	"dshusdock/go_project/internal/services/appmgmntsvc"
	"dshusdock/go_project/internal/services/messagebus"
	"dshusdock/go_project/internal/services/session"
	"dshusdock/go_project/internal/views/appsidenavvw"
	"dshusdock/go_project/internal/views/auditvw"
	"dshusdock/go_project/internal/views/base"
	"dshusdock/go_project/internal/views/createauditvw"
	"dshusdock/go_project/internal/views/dbsidenavvw"
	"dshusdock/go_project/internal/views/dbsourcevw"
	"dshusdock/go_project/internal/views/dbutilvw"
	"dshusdock/go_project/internal/views/headervw"
	"dshusdock/go_project/internal/views/layoutvw"
	"dshusdock/go_project/internal/views/queryvw"
	"dshusdock/go_project/internal/views/tablevw"
	"log/slog"
	"net/http"
	"strings"
)

type RenderView struct {
	App 		*config.AppConfig 
	ViewHandlers 	map[string]con.ViewHandler
}

type DisplayData struct {
	Base 		*con.BaseTemplateparams
	Tmplt   	map[string]*any
}

var RenderViewSvc *RenderView

func MapRenderViewSvc(r *RenderView) {
	RenderViewSvc = r
}

func InitRouteHandlers() {
	// Register the views
	RenderViewSvc.ViewHandlers["basevw"] = base.AppBaseVw.RegisterHandler()
	RenderViewSvc.ViewHandlers["header"] = headervw.AppHeaderVw.RegisterHandler()
	RenderViewSvc.ViewHandlers["appsidenavvw"] = appsidenavvw.AppAppSideNavVw.RegisterHandler()
	RenderViewSvc.ViewHandlers["layoutvw"] = layoutvw.AppLayoutVw.RegisterHandler()
	RenderViewSvc.ViewHandlers["dbutilvw"] = dbutilvw.AppDBUtilVw.RegisterHandler()
	RenderViewSvc.ViewHandlers["dbsidenavvw"] = dbsidenavvw.AppDBSideNavVw.RegisterHandler()
	RenderViewSvc.ViewHandlers["dbsourcevw"] = dbsourcevw.AppDBSourceVw.RegisterHandler()
	RenderViewSvc.ViewHandlers["tablevw"] = tablevw.AppTableVw.RegisterHandler()
	RenderViewSvc.ViewHandlers["queryvw"] = queryvw.AppQueryVw.RegisterHandler()
	RenderViewSvc.ViewHandlers["auditvw"] = auditvw.AppAuditVw.RegisterHandler()
	RenderViewSvc.ViewHandlers["createauditvw"] = createauditvw.AppCreateAuditVw.RegisterHandler()
	RenderViewSvc.ViewHandlers["auditvw"] = auditvw.AppAuditVw.RegisterHandler()
}

func NewRenderViewSvc(app *config.AppConfig) *RenderView {
	
	obj := &RenderView{
		App: app,
		ViewHandlers: make(map[string]con.ViewHandler),
	}
	RenderViewSvc = obj

	messagebus.GetBus().Subscribe("Event:Click", RenderViewSvc.HandleMBusRequest)
	return RenderViewSvc
	
}

func (rv *RenderView) ProcessInit(w http.ResponseWriter, r *http.Request) {
	slog.Info("Processing Init Event")
	obj := DisplayData{
		Base: nil,
		Tmplt: make(map[string]*any),
	}

	event := con.AppEvent{
		Request: r,
		EventId: con.EVENT_STARTUP,
		EventStr: "startup",
	}	

	evt := event.EventStr
	for _, v := range con.APP_VIEWS[evt].Views {
		result := rv.ViewHandlers[v].HandleRequest(w, event)
		obj.Tmplt[v] = &result
	}
	obj.Base = con.GetBaseTemplateObj(con.APP_VIEWS[evt].BaseVals)

	render.RenderAppTemplate(w, nil, obj, con.APP_VIEWS[evt].Tmplt)
	slog.Info("Processing Init Event - Done")
}

func (rv *RenderView) ProcessStateChange(w http.ResponseWriter, r *http.Request) {
	slog.Info("----------[Processing State Change]----------")
	data := r.PostForm

	event := con.AppEvent{
		Src: data.Get("src"),
		Target: data.Get("target"),
		Request: r,
		EventId: con.EVENT_STATECHANGE,		
	}
	
	for _, v := range con.APP_VIEWS["state_change"].Views {
		_ = rv.ViewHandlers[v].HandleRequest(w, event)
	}

	slog.Info("Processing State Change - Done")
}

func (rv *RenderView)  ProcessEvent(w http.ResponseWriter, r *http.Request ) {

	slog.Debug("----------[Entering ProcessEvent]----------")
	data := r.PostForm

	event := con.AppEvent{
		Src: data.Get("src"),
		Request: r,
		EventId: data.Get("event"),
		EventStr: data.Get("view") + "_" + data.Get("src"),
		Data: data.Get("data"),
	}

	slog.Info("EventStr: ", "value", event.EventStr)

	obj := DisplayData{
		Base: nil,
		Tmplt: make(map[string]*any),
	}

	evt := con.ExtractEventStr(event.EventStr)

	// Verify the user is logged in if not, redirect to login page
	if event.EventStr != "loginvw_login" && event.EventStr != "loginvw_create-account" {
		if !session.SessionSvc.SessionMgr.Exists(r.Context(), "LoggedIn") {
			slog.Info("User not logged in")
			w.Header().Set("HX-Redirect", "/")
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	
	for _, v := range con.APP_VIEWS[evt].Views {
		var result interface{} = nil
		result = rv.ViewHandlers[v].HandleRequest(w, event)
		if result == nil {
			slog.Info("No view to render")
			return
		}
		obj.Tmplt[v] = &result
		
	}
	loggedInBits := appmgmntsvc.GetCurrentLogonState(r.Context())
	obj.Base = con.GetBaseTemplateObj(con.APP_VIEWS[evt].BaseVals | loggedInBits)
	
	if con.APP_VIEWS[evt].Tmplt != con.RM_NONE {
		render.RenderAppTemplate(w, nil, obj, con.APP_VIEWS[evt].Tmplt)
	}
	slog.Debug("Processing Event - Done")
}

func (rv *RenderView) ProcessChangeEvent(w http.ResponseWriter, r *http.Request) {
	slog.Info("----------[Processing Change Event]----------")
	data := r.PostForm

	event := con.AppEvent{
		Src: data.Get("src"),
		Request: r,
		EventId: con.EVENT_CHANGE,
	}

	evt := con.ExtractEventStr(event.EventStr)

	for _, v := range con.APP_VIEWS[evt].Views {
		_ = rv.ViewHandlers[v].HandleRequest(w, event)
	}

	slog.Info("Processing Change Event - Done")
}


// Deprecated
func (rv *RenderView) ProcessRequest(w http.ResponseWriter, r *http.Request, view string) {
	var rslt any
	var _view int

	ev := con.AppEvent{
		Request: r,
		EventId: con.EVENT_STARTUP,
	}
	
	slog.Info("[renderview] - ProcessRequest")
	obj := DisplayData{
		Base: nil,
		Tmplt: make(map[string]*any),
	}

	if (false) { // some special condition) 
		// do something special that returns rslt
	} else {
		rslt = rv.ViewHandlers[view].HandleRequest(w, ev)
	}
	obj.Tmplt[view] = &rslt

	switch view {
	case "basevw":		
		//_view = rslt.(base.BaseVwData).View
		obj.Base = nil
	default:
	}	

	if _view == con.RM_NONE { 
		slog.Info("No view to render")
		return
	}

	// Now go build the view and render it
	render.RenderAppTemplate(w, nil, obj, _view)
}

func (rv *RenderView) HandleMBusRequest(w http.ResponseWriter, r *http.Request) {
	slog.Info("[renderview] - HandleMBusRequest")
    d := r.PostForm
	id := d.Get("view")	

	switch id {
	case "basevw":
	}
}

func stripChars(str, chr string) string {
    return strings.Map(func(r rune) rune {
        if strings.IndexRune(chr, r) < 0 {
            return r
        }
        return -1
    }, str)
}