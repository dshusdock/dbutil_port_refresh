package constants

import "log/slog"

/* Be sure that the template you are trying to render contains the file
   you are trying to render
*/

const (
	BASE = "./ui/html/views/"
	APP = "app/"
	DBUTILS = "dbutils/"
)

const (
	// BASE FILES
	FILE_BASE_TMPL = BASE + APP + "base.tmpl"
	FILE_HEADER_TMPL = BASE + APP + "header.tmpl"
	FILE_MAINVW_TMPL = BASE + APP + "main_vw.tmpl"
	FILE_LAYOUT_TMPL = BASE + APP + "layout.tmpl"
	FILE_LOGINVW_TMPL = BASE + APP + "loginvw.tmpl"
	FILE_CREATEACCNTVW_TMPL = BASE + APP + "createaccountvw.tmpl"
	FILE_SNIPPETS_TMPL = BASE + APP + "snippets.tmpl"
	FILE_APPSIDENAV_TMPL = BASE + APP + "appsidenavvw.tmpl"
	FILE_DBUTILVW_TMPL = BASE + DBUTILS + "dbutilvw.tmpl"
	FILE_DBSIDENAVVW_TMPL = BASE + DBUTILS + "dbsidenavvw.tmpl"
	FILE_DBSOURCEVW_TMPL = BASE + DBUTILS + "dbsourcevw.tmpl"
	FILE_TABLE_TMPL = BASE + DBUTILS + "tablevw.tmpl"	
	FILE_QUERYVW_TMPL = BASE + DBUTILS + "queryvw.tmpl"
	FILE_AUDITVW_TMPL = BASE + DBUTILS + "auditvw.tmpl"
	FILE_CREATEAUDITVW_TMPL = BASE + DBUTILS + "createauditvw.tmpl"
)

// /////////////Rendered File Map///////////////
// The order of the the following two structures have to match
// The first structure is the key and the second is the value
const (
	RM_HOME = iota
	RM_HEADER
	RM_MAINVW
	RM_DBMAINVW
	RM_SNIPPETS
	RM_LOGINVW
	RM_CREATEACCNTVW
	RM_APPSIDENAVVW
	RM_DBUTILVW
	RM_DBSIDENAVVW	// 10
	RM_DBSOURCEVW
	RM_TABLEVW
	RM_QUERYVW
	RM_AUDITVW
	RM_CREATEAUDITVW
	RM_NONE
)

type RenderedFileMap struct {
	HOME           	[]string
	HEADER 	   		[]string
	MAINVW			[]string
	DBMAINVW		[]string
	SNIPPETS		[]string
	LOGINVW			[]string
	CREATEACCNTVW	[]string
	APPSIDENAVVW	[]string
	DBUTILVW		[]string
	DBSIDENAVVW		[]string //10
	DBSOURCEVW		[]string
	TABLEVW			[]string
	QUERYVW			[]string
	AUDITVW			[]string
	CREATEAUDITVW	[]string
	NONE           	[]string
}

func RENDERED_FILE_MAP() *RenderedFileMap {
	return &RenderedFileMap{
		HOME: []string{
			FILE_BASE_TMPL,
			FILE_HEADER_TMPL,
			FILE_MAINVW_TMPL,
			FILE_LAYOUT_TMPL,
			FILE_LOGINVW_TMPL,
			FILE_CREATEACCNTVW_TMPL,
			FILE_APPSIDENAV_TMPL,
			FILE_DBUTILVW_TMPL,
			FILE_DBSIDENAVVW_TMPL,	
			FILE_TABLE_TMPL,	
			FILE_QUERYVW_TMPL,	
			FILE_AUDITVW_TMPL,
			FILE_CREATEAUDITVW_TMPL,
		},	
		HEADER: []string{
			FILE_HEADER_TMPL,
		},	
		MAINVW: []string{
			FILE_MAINVW_TMPL,			
			FILE_DBUTILVW_TMPL,
			FILE_DBSIDENAVVW_TMPL,	
			FILE_TABLE_TMPL,		
			FILE_QUERYVW_TMPL,
			FILE_AUDITVW_TMPL,
			FILE_CREATEAUDITVW_TMPL,
		},
		DBMAINVW: []string{
			FILE_DBSIDENAVVW_TMPL,
			FILE_DBUTILVW_TMPL,
			FILE_AUDITVW_TMPL,
			FILE_CREATEAUDITVW_TMPL,
			FILE_QUERYVW_TMPL,
			FILE_TABLE_TMPL,
		},
		SNIPPETS: []string{
			FILE_SNIPPETS_TMPL,			
			//"./ui/html/views/snippets.tmpl",
		},
		LOGINVW: []string{
			FILE_BASE_TMPL,
			FILE_LAYOUT_TMPL,
			FILE_LOGINVW_TMPL,
		},
		CREATEACCNTVW: []string{
			FILE_BASE_TMPL,
			FILE_CREATEACCNTVW_TMPL,
		},
		APPSIDENAVVW: []string{
			FILE_APPSIDENAV_TMPL,
		},
		DBUTILVW: []string{
			FILE_DBUTILVW_TMPL,
			FILE_DBSIDENAVVW_TMPL,
			FILE_TABLE_TMPL,
			FILE_QUERYVW_TMPL,			
			FILE_AUDITVW_TMPL,
			FILE_CREATEAUDITVW_TMPL,
			
		},
		DBSIDENAVVW: []string{
			FILE_DBSIDENAVVW_TMPL,
		},
		DBSOURCEVW: []string{
			FILE_DBSOURCEVW_TMPL,
		},	
		TABLEVW: []string{
			FILE_TABLE_TMPL,
		},	
		QUERYVW: []string{
			FILE_QUERYVW_TMPL,
			FILE_DBSIDENAVVW_TMPL,
		},
		AUDITVW: []string{
			FILE_AUDITVW_TMPL,
			FILE_CREATEAUDITVW_TMPL,
		},
		CREATEAUDITVW: []string{
			FILE_CREATEAUDITVW_TMPL,
		},
		NONE: []string{
			"",	
		},
	}
}

func GetRenderInfo(viewType int) RenderInfo{
	files := RENDERED_FILE_MAP()

	slog.Debug("GetRenderInfo: ", "viewType", viewType)

	ri := RenderInfo{
		TemplateName: "",
		TemplateFiles: []string{},
	}

	switch viewType {
	case RM_HOME:
		ri.TemplateFiles = files.HOME
		ri.TemplateName = "base"
	case RM_MAINVW:
		ri.TemplateFiles = files.MAINVW
		ri.TemplateName = "main_vw"	
	case RM_SNIPPETS:
		ri.TemplateFiles = files.SNIPPETS
		ri.TemplateName = "snippets"		
	case RM_LOGINVW:
		ri.TemplateFiles = files.LOGINVW
		ri.TemplateName = "base"
	case RM_NONE:
		ri.TemplateFiles = files.NONE
		ri.TemplateName = ""	
	case RM_DBSOURCEVW:
		ri.TemplateFiles = files.DBSOURCEVW
		ri.TemplateName = "dbsourcevw"
	case RM_DBMAINVW:
		ri.TemplateFiles = files.DBMAINVW
		ri.TemplateName = "dbutilvw"
	case RM_TABLEVW:
		ri.TemplateFiles = files.TABLEVW
		ri.TemplateName = "tablevw"
	case RM_QUERYVW:
		ri.TemplateFiles = files.QUERYVW
		ri.TemplateName = "queryvw"
	case RM_AUDITVW:
		ri.TemplateFiles = files.AUDITVW
		ri.TemplateName = "auditvw"
	case RM_CREATEAUDITVW:
		ri.TemplateFiles = files.CREATEAUDITVW
		ri.TemplateName = "createauditvw"
	default:
		slog.Debug("No view to render")
		ri.TemplateFiles = files.HOME
		ri.TemplateName = "base"
	}

	return ri
}