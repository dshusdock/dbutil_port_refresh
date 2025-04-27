package constants

import (
	"net/http"
)

const (
	FILESA = iota
	FILESB
	FILESC
)

const (
	EVENT_STARTUP = "Event_Startup"
	EVENT_CLICK  = "Event_Click"
	EVENT_CHANGE  = "Event_Change"
	EVENT_STATECHANGE = "Event_StateChange"
)

const (
	VW_INDEX = iota
)

// AppMode constants
const (
	NONE = iota
	HOME
	QUERY
	AUDIT
	LOADDB
	DBQUERY
	REMTASKS
	TABLEMAPPER
	COLUMNFINDER
)

type EventData struct {
	Id        string
	EventType string
	Event     string
}

type HtmxInfo struct {
	Url string
}

type SubElement struct {
	Type string
	Lbl  string
}

type ViewInterface interface {
	// HandleHttpRequest(w http.ResponseWriter, d url.Values /*ViewInfo*/)
	HandleHttpRequest(w http.ResponseWriter, r *http.Request)
}

type ViewHandler interface {
	HandleRequest(w http.ResponseWriter, event AppEvent) any
	//HandleMBusRequest(w http.ResponseWriter, r *http.Request) any
}

type ViewInfo struct {
	Event   int
	Type    string
	Label   string
	ViewId  string
	ViewStr string
}

type AppEvent struct {
	View 		string
	Src   		string
	Type   		string
	EventId 	string
	EventStr 	string
	Target   	string
	Request 	*http.Request
	Data 		string
}

type RenderInfo struct {
	TemplateName string
	TemplateFiles []string
}

type RowData struct {
	Data []string
}

// Table supports the display of data in a table format
type UnigyTable struct {
	Name 				string
	Columns 			[]string
	Rows 				[]RowData
	RowsSlice 			[]RowData
	Start 				int
	End 				int
	TableSize 			int
	ShowSize 			int
	PageSize 			int
	CurrentPage 		int
	PageAry				[]string
	OptionValues 		[]OptionsState
	SqlQuery 			string
	PageBtns 			[]PageButton
	Htmx	 			bool
	SearchInput 		string
	SortSqlQuery 		string
	SortColumn 			string
	SortDirection 		string
	Detail 				bool
}

func CopyUnigyTable(m *UnigyTable) *UnigyTable {
	return &UnigyTable{
		Name: m.Name,
		Columns: m.Columns,
		Rows: m.Rows,
		RowsSlice: m.RowsSlice,
		Start: m.Start,
		End: m.End,
		TableSize: m.TableSize,
		ShowSize: m.ShowSize,
		PageSize: m.PageSize,
		CurrentPage: m.CurrentPage,
		PageAry: m.PageAry,
		OptionValues: m.OptionValues,
		SqlQuery: m.SqlQuery,
		PageBtns: m.PageBtns,
		Htmx: m.Htmx,
		SearchInput: m.SearchInput,
		SortSqlQuery: m.SortSqlQuery,
		SortColumn: m.SortColumn,
		SortDirection: m.SortDirection,
		Detail: m.Detail,
	}
}

type TableControls struct {
	OptionValues 		[]OptionsState
	PageBtns 			[]PageButton
}

type PageButton struct {
	Label 		string
	Disabled 	bool
	Visible 	bool
	Class 		string
	Name 		string
	Htmx 		bool
	SrcId 		string
}

type OptionsState struct {
	OptionValue 	string
	OptionSelected 	bool
}

// UnigyTable constructor
func NewUnigyTable() *UnigyTable {
	return &UnigyTable{
		Name: "",
		Columns: []string{},
		Rows: []RowData{},
		RowsSlice: []RowData{},
		Start: 0,
		End: 0,
		TableSize: 0,
		ShowSize: 50,
		PageSize: 0,
		CurrentPage: 1,
		PageAry: []string{},
		OptionValues: []OptionsState{
			{OptionValue: "10", OptionSelected: false},
			{OptionValue: "25", OptionSelected: false},
			{OptionValue: "50", OptionSelected: true},
			{OptionValue: "100", OptionSelected: false},
		},
		SqlQuery: "",
		PageBtns: []PageButton{
			{Label: "", Name: "xrev", Disabled: true, Class: "fa-solid fa-angle-double-left fa-lg", SrcId: "xrev", Htmx: false},
			{Label: "", Name: "srev", Disabled: true, Class: "fa-solid fa-angle-left fa-lg", SrcId: "srev", Htmx: false},
			{Label: "Rows 10 to 20 of 100", Name: "page", Disabled: true, Class: "page_select", SrcId: "pgselect", Htmx: true},
			{Label: "", Name: "sfwd", Disabled: true, Class: "fa-solid fa-angle-right fa-lg",  SrcId: "sfwd", Htmx: false},
			{Label: "", Name: "xfwd", Disabled: true, Class: "fa-solid fa-angle-double-right fa-lg",  SrcId: "xfwd", Htmx: false},
		},
		Htmx: false,
		SearchInput: "",
		SortSqlQuery: "",
		SortColumn: "",
		SortDirection: "",
		Detail: true,
	}
} 

type DBLoaderInfo struct {
	SrcIP 			string
	SrcUser 		string
	SrcPassword 	string
	DBName 			string
	UnzipRequired 	bool
	SQLFileName 	string
	FileDir 		string
	FileInNas 		bool
	UseLocalServer 	bool
	TargetIP 		string
}

type QueryInfo struct {
	Sql 	string
	DbType 	string
	DBHost 	string
	DBName 	string
}
 




