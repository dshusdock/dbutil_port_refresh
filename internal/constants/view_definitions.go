package constants

import (
	"log/slog"
	"reflect"
	"strings"
)

var APP_VIEWS = make (map[string]ViewDef)

type ViewDef struct {
	Views 		[]string
	BaseVals 	Bits
	Tmplt 		int
}

func init() {
	APP_VIEWS["startup"] = ViewDef{Views: []string{"basevw"},  BaseVals: BIT_NONE, Tmplt: RM_HOME,}
	APP_VIEWS["loginvw_create-account"] = ViewDef{Views: []string{"basevw"},  BaseVals: BIT_CreateAccount, Tmplt: RM_HOME,}
	APP_VIEWS["loginvw_login"] = ViewDef{Views: []string{"basevw",},  BaseVals: BIT_NONE, Tmplt: RM_HOME,}

	APP_VIEWS["createacctvw_submit"] = ViewDef{Views: []string{"basevw"},  BaseVals: BIT_CreateAccount, Tmplt: RM_HOME,}
	APP_VIEWS["createacctvw_ok"] = ViewDef{Views: []string{"basevw"},  BaseVals: BIT_CreateAccount, Tmplt: RM_HOME,}
	
	APP_VIEWS["appsidenavvw_db-analyzer"] = ViewDef{Views: []string{"appsidenavvw","dbsidenavvw",},BaseVals: BIT_DBUtilApp,Tmplt: RM_MAINVW,}
	APP_VIEWS["appsidenavvw_"] = ViewDef{Views: []string{"appsidenavvw"},BaseVals: BIT_NONE,Tmplt: RM_NONE,}
	
	APP_VIEWS["dbsidenavvw_"] = ViewDef{Views: []string{"dbsidenavvw", "tablevw"},BaseVals: BIT_DBUtilApp,Tmplt: RM_TABLEVW,}
	APP_VIEWS["dbsidenavvw_dbsource"] = ViewDef{Views: []string{"dbsidenavvw", "dbsourcevw"},BaseVals: BIT_NONE, Tmplt: RM_DBSOURCEVW,}
	APP_VIEWS["dbsidenavvw_ffquery"] = ViewDef{Views: []string{"dbsidenavvw", "queryvw"},BaseVals: BIT_DBUtilApp, Tmplt: RM_QUERYVW,}
	APP_VIEWS["dbsidenavvw_audit_config"] = ViewDef{Views: []string{"dbutilvw", "auditvw", "createauditvw", "dbsidenavvw"},BaseVals: BIT_DBUtilApp, Tmplt: RM_AUDITVW,}
	APP_VIEWS["dbsidenavvw_audit_create"] = ViewDef{Views: []string{"dbutilvw", "auditvw", "createauditvw", "dbsidenavvw"},BaseVals: BIT_DBUtilApp, Tmplt: RM_CREATEAUDITVW,}
	APP_VIEWS["dbsidenavvw_audit_execute"] = ViewDef{Views: []string{"auditvw", "tablevw", },BaseVals: BIT_DBUtilApp, Tmplt: RM_DBMAINVW,}

	APP_VIEWS["dbsourcevw_target-select"] = ViewDef{Views: []string{"dbsourcevw",},BaseVals: BIT_DBUtilApp,Tmplt: RM_DBSOURCEVW,}
	APP_VIEWS["dbsourcevw_db-select"] = ViewDef{Views: []string{"dbsourcevw",},BaseVals: BIT_DBUtilApp,Tmplt: RM_DBSOURCEVW,}

	APP_VIEWS["tablevw_data-cell"] = ViewDef{Views: []string{"tablevw",},BaseVals: BIT_DBUtilApp,Tmplt: RM_TABLEVW,}
	APP_VIEWS["tablevw_table-search"] = ViewDef{Views: []string{"tablevw",},BaseVals: BIT_DBUtilApp,Tmplt: RM_TABLEVW,}
	APP_VIEWS["tablevw_"] = ViewDef{Views: []string{"tablevw",},BaseVals: BIT_DBUtilApp,Tmplt: RM_TABLEVW,}

	APP_VIEWS["queryvw_btn-execute"] = ViewDef{Views: []string{"queryvw", "dbutilvw", "dbsidenavvw", "tablevw"},BaseVals: BIT_DBUtilApp,Tmplt: RM_MAINVW,}

	APP_VIEWS["auditvw_"] = ViewDef{Views: []string{"auditvw",},BaseVals: BIT_DBUtilApp,Tmplt: RM_NONE,}
	APP_VIEWS["auditvw_selectall_toggle"] = ViewDef{Views: []string{"auditvw",},BaseVals: BIT_DBUtilApp, Tmplt: RM_AUDITVW,}
	
	//dbsidenavvw_audit


	
	APP_VIEWS["state_change"] = ViewDef{Views: []string{"headervw", }, BaseVals: BIT_NONE, Tmplt: RM_NONE,}
}

type BaseTemplateparams struct {
	Title 						string
	LoggedIn 					bool	
	DisplayLogin  				bool
	CreateAccount 				bool
	CreatAcctResponse 			bool
	MainTable	  				bool
	DBUtilApp					bool
	
}
 
type Bits uint16

const (
	BIT_LoggedIn Bits = 1 << iota
	BIT_DisplayLogin
	BIT_CreateAccount
	BIT_CreatAcctResponse
	BIT_MainTable
	BIT_DBUtilApp
	
	BIT_NONE 
)

func GetBaseTemplateObj(b Bits) *BaseTemplateparams{
	return SetBits(b)
}

func SetBits(bits Bits) *BaseTemplateparams {
	slog.Info("number of elements in struct:", "value", reflect.TypeOf(BaseTemplateparams{}).NumField())
	base := BaseTemplateparams{
		Title: "",
		LoggedIn: false,
		DisplayLogin: true,
		CreateAccount: false,
		CreatAcctResponse: false,		
		MainTable: false,
		DBUtilApp: false,		
	}
 
	fType := reflect.TypeOf(base) 
    fVal := reflect.New(fType)

	slen := reflect.TypeOf(BaseTemplateparams{}).NumField()
	
	var x=0
	for i := 1; i < slen; i++ {				
		if bits & (1 << x) > 0 {
			fVal.Elem().Field(i).SetBool(true)			
		} else {
			fVal.Elem().Field(i).SetBool(false)
		}
		x++
	}
	val := fVal.Elem().Interface().(BaseTemplateparams)

	return &val
}

func ExtractEventStr(str string) string {
	var subKey = ""

	// Find all keys matching the pattern "user_*"
	// pattern := regexp.MustCompile(`user_\d+`)
	for key := range APP_VIEWS {
		if strings.Compare(str, key) == 0 {
			slog.Info("Found Key (exact): ",  "value", key)
			return key
		}

		if strings.Contains(str, key) {
			slog.Info("Found Key: ",  "value", key)
			subKey = key
		}
	}
	return subKey
}
