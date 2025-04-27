package tablevw

import (
	"dshusdock/go_project/config"
	con "dshusdock/go_project/internal/constants"
 	"dshusdock/go_project/internal/services/auditsvc"

	//"dshusdock/go_project/internal/services/appmgmntsvc"
	//"dshusdock/go_project/internal/services/dbservice/unigydb"
	"dshusdock/go_project/internal/services/appmgmntsvc"
	"dshusdock/go_project/internal/services/dbservice/unigydb"
	"dshusdock/go_project/internal/services/session"

	//ct "dshusdock/go_project/internal/views/control_hdrvw"
	"encoding/gob"

	"encoding/json"
	//"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

func GetData() string {
	return "TableVw"
}

type TableVw struct {
	App *config.AppConfig
}

var AppTableVw *TableVw

func init() {
	AppTableVw = &TableVw{
		App: nil,
	}
	gob.Register(AppTableVwData{})
}

func (m *TableVw) RegisterHandler() con.ViewHandler {
	return &TableVw{}
}

func (m *TableVw) HandleRequest(w http.ResponseWriter, event con.AppEvent) any {
	slog.Info("[TableVw] - HandleRequest")
	var obj AppTableVwData

	if session.SessionSvc.SessionMgr.Exists(event.Request.Context(), "TableVw") {
		obj = session.SessionSvc.SessionMgr.Pop(event.Request.Context(), "TableVw").(AppTableVwData)
	} else {
		obj = *CreateTableVwData()	
	}

	obj.ProcessHttpRequest(w, event)	
	session.SessionSvc.SessionMgr.Put(event.Request.Context(), "TableVw", obj)

	return obj
}
 
///////////////////// Layout View Data //////////////////////

type AppTableVwData struct {
	Base 			*con.BaseTemplateparams
	Data 			*con.UnigyTable
	DataCopy 		*con.UnigyTable
	Detail 			bool
	DBHost 			string
	DBName 			string
	QryInfo 		con.QueryInfo
	DBQueryMode 	bool
	isPageEvent 	bool
}

// Page button states
const (
	OFF = 0x0
	ALL = REV | FWD
	XREV = 0x1 | PAGE
	SREV = 0x2 | PAGE
	SFWD = 0x8 | PAGE
	XFWD = 0x10 | PAGE
	REV = SREV | XREV | PAGE
	FWD = SFWD | XFWD | PAGE
	PAGE = 0x4
)

func CreateTableVwData() *AppTableVwData {
	return &AppTableVwData{
		Base: 	nil,
		Data: 	nil,
		Detail: true,
		DBHost: "",
		DBName: "",
		QryInfo: con.QueryInfo{},
		DBQueryMode: false,
		isPageEvent: false,
	}
}

func (m *AppTableVwData) ProcessHttpRequest(w http.ResponseWriter, event con.AppEvent) *AppTableVwData{
	slog.Info("[TableVwData]: - ProcessHttpRequest")
	
	host, name := appmgmntsvc.GetActiveDBHostAndName(event.Request.Context())
	m.DBHost = host
	m.DBName = name
	m.QryInfo = con.QueryInfo{
		DBHost: host,
		DBName: name,
	}

	if event.EventStr == "dbsidenavvw_ffquery" {
		m.Data = nil
	}

	if event.EventStr == "dbsourcevw_db-select" {
		m.Data = nil
		return m
	}

	if event.EventStr == "queryvw_btn-execute" || 
		event.EventStr == "auditvw_btn_ExecuteAudit" || 
		event.EventStr == "dbsidenavvw_audit" ||
		event.EventStr == "dbsidenavvw_query-cat" ||
		event.EventStr == "headervw_btn_RemTasks" ||
		event.EventStr == "headervw_btn_TableMapper" ||
		event.EventStr == "headervw_btn_ColumnFinder" {
		m = m.processTableRequest(event)
	}

	// Process the events originating from the table view....
	switch event.Src {
	case "data-cell":
		fallthrough
	case "button": // This appears to do nothing...leaving for now
		m = m.processTableRequest(event)
	case "table-search":
		m = m.processSearchEvent(event)
	case "pgsize":
		m = m.processSelectEvent(w, event)
	case "xrev":
		fallthrough
	case "srev":
		fallthrough
	case "sfwd":
		fallthrough
	case "xfwd":
		fallthrough
	case "pgselect":
		
		m = m.processViewBtnEvent(event)
	case "page-select":
		m = m.processPageSelectEvent(event)
	case "sort":
		m = m.processSortEvent(event)
	default:	
		m = m.processTableRequest(event)	
	}

	return m
}

type RowSQL struct {
	SQL string `json:"sql"`
	Name string `json:"name"`
}

func (m *AppTableVwData) processTableRequest(event con.AppEvent) *AppTableVwData {
	slog.Info("[TableVwData]: - processTableRequest")
	var sql string
	var label string
	ctx := event.Request.Context()

	if event.EventStr == "dbsidenavvw_audit" {
		m.Data = nil
		return m
	}

	if event.EventStr == "queryvw_btn-execute" || event.EventStr == "auditvw_btn_ExecuteAudit" {
		sql = event.Request.PostForm.Get("querytext")
		label = "User Query"
	} else {
		// Determine which button was clicked and map it to the appropriate SQL query
		// There should be only one button item to map to
		for _, item := range con.ButtonMap {
			if item.Name == event.Src {	
				sql = item.SQLString
				label = item.Label	
			}
		}
	}

	switch event.EventStr {
	case "headervw_btn_RemTasks":
		sql = "SELECT taskId, createDate, client, name, summary FROM Tasks"
		m.QryInfo.Sql = sql
		m.QryInfo.DbType = "local"
		m.Data = unigydb.PerformQuery(m.QryInfo)
	case "dbsidenavvw_query-cat":
		sql = "SELECT * FROM Query"
		m.QryInfo.Sql = sql
		m.QryInfo.DbType = "local"
		m.Data = unigydb.PerformQuery(m.QryInfo)
		//m.DBQueryMode = true
		appmgmntsvc.SetAppMode(event.Request.Context(), con.DBQUERY)		
	case "headervw_btn_TableMapper":
		sql = "select uc.id, ufk2.tableName as 'rvsFK_Tbl', uc.tableName as 'Table Name', uc.columnName as 'Column Name', " +
			"ufk.referenceTable as 'ForeignKey Table', ufk.referenceColumn as 'ForeignKey Column' from UnigyColumns uc " +
			"left join UnigyForeignKey ufk on (uc.tableName=ufk.tableName and uc.columnName = ufk.columnName) " +
			"left join UnigyForeignKey ufk2 on (uc.tableName=ufk2.referenceTable and uc.columnName='id') " +
			"order by uc.tableName, uc.id;"
		m.QryInfo.Sql = sql
		m.QryInfo.DbType = "local"
		m.Data = unigydb.PerformQuery(m.QryInfo)
	case "headervw_btn_ColumnFinder":
		sql = "select tableName as 'Table Name', columnName as 'Column Name' from UnigyColumns;"
		m.QryInfo.Sql = sql
		m.QryInfo.DbType = "local"
		m.Data = unigydb.PerformQuery(m.QryInfo)
	case "dbsidenavvw_audit_execute":
		m.QryInfo.Sql = auditsvc.AuditSvc.GetAuditSQLString(ctx)
		if m.QryInfo.Sql == "" {
			return m
		}
		m.Data = unigydb.PerformQuery(m.QryInfo)
		m.Data.Name = "Audit Results"
		appmgmntsvc.SetAppMode(event.Request.Context(), con.AUDIT)
		//m.DataCopy = m.Data
		return m
	case "control-hdrvw_button_AuditResults":
		m.Data = con.CopyUnigyTable(m.DataCopy)
		appmgmntsvc.SetAppMode(ctx, con.AUDIT)
	case "tablevw_data-cell":
		slog.Info("[TableVwData]: Row Selected - view_str: ", "ViewStr",event.Data)
		mode := appmgmntsvc.GetAppMode(ctx)
		
		
		var row RowSQL
		err :=json.Unmarshal([]byte(event.Data), &row)
		if err != nil {
			slog.Info("Error decoding JSON: ",  "Error", err.Error())
		}

		if mode == con.AUDIT {			
			slog.Debug("Handling audit case")
			m.DataCopy = con.CopyUnigyTable(m.Data)
			sql, tblName := auditsvc.AuditSvc.GetSQLString(ctx, row.Name)
			m.QryInfo.Sql = sql
			m.QryInfo.DbType = "unigy"
			m.Data = unigydb.PerformQuery(m.QryInfo)
			appmgmntsvc.SetAppMode(ctx, con.NONE)
			m.Data.Name = tblName
			return m
		}

		if mode != con.DBQUERY {
			m.processClickSearch(row.SQL)
		}
		
		sql = row.SQL
		// Verify that the SQL string is a select statement
		if strings.Contains(sql, "select") {
			m.QryInfo.Sql = sql
			m.QryInfo.DbType = "unigy"
			m.Data = unigydb.PerformQuery(m.QryInfo)
			label = row.Name
		} else {
			m.processClickSearch(sql)
			//m.DBQueryMode = false
			appmgmntsvc.SetAppMode(ctx, con.NONE)
			
			return m
		}
		
	default:		
		slog.Info("Handling default case")
		m.QryInfo.Sql = sql
		m.QryInfo.DbType = "unigy"
		m.Data = unigydb.PerformQuery(m.QryInfo)
	}

	if m.Data == nil {
		return m
	}

	m.setTableEndpoints()
	m.setPageSize()
	m.updatePagerLabels(true)
	m.Data.SqlQuery = sql
	m.Data.Name = label

	if event.EventStr == "auditvw_btn_ExecuteAudit" {
		m.Data.Detail = false
	}
	
	return m
}


func (m *AppTableVwData) processSortEvent(event con.AppEvent) *AppTableVwData {
	slog.Info("[TableVwData] - processSortEvent")
	slog.Debug("[TableVwData]:  ", "value", event.Src)
	var sql string
	var sortDir string
	
	if m.Data.SortColumn == "" || m.Data.SortColumn != event.Src {
		m.Data.SortColumn = event.Src
	} 

	origSql := m.Data.SqlQuery
	baseSql := strings.Split(origSql, " order by ")
	
	if m.Data.SortDirection == "" || m.Data.SortDirection == "asc" {
		sortDir = "desc"
		sql = baseSql[0] + " order by " + event.Src + " desc"
	} else {
		sortDir = "asc"
		sql = baseSql[0] + " order by " + event.Src + " asc"
	}
	slog.Debug("SQL: ", "sql", sql)

	tableName := m.Data.Name
	m.QryInfo.Sql = sql
	m.QryInfo.DbType = "unigy"
	m.Data = unigydb.PerformQuery(m.QryInfo)
	m.Data.SortDirection = sortDir
	m.setTableEndpoints()
	m.setPageSize()
	m.updatePagerLabels(true)
	m.Data.Name = tableName
	m.Data.SqlQuery = origSql

	return m
}

func (m *AppTableVwData) processSearchEvent(event con.AppEvent) *AppTableVwData {
	slog.Info("[TableVwData] - processSearchEvent")

	data := event.Request.PostForm
	key := data.Get("search")
	m.Data.SearchInput = key

	m.Data.RowsSlice = []con.RowData{}
	// Search the table for the key
	if key == "" {
		slog.Info("Key is null")
		m.setTableEndpoints()
		m.setPageSize()
		m.updatePagerLabels(true)
	} else {
		m.setPageButtonState(OFF)
		for x := 0; x < m.Data.TableSize; x++ {
			var row = m.Data.Rows[x]
			// if strings.Contains(strings.Join(row.Data, " "), key) {
			// 	slog.Debug("[TableVwData] - got Row", "row", row)
			// 	m.Data.RowsSlice = append(m.Data.RowsSlice, row)
			// }

			index := strings.Index(strings.ToLower(strings.Join(row.Data, " ")), 
				strings.ToLower(key))
			if index != -1 {
				m.Data.RowsSlice = append(m.Data.RowsSlice, row)
			}
		}		
	}
	
	slog.Debug("[TableVwData]: Search: ", "key", key)
	return m
}

func (m *AppTableVwData) processClickSearch(key string) *AppTableVwData {
	slog.Info("[TableVwData] - processClickSearch")

	m.Data.SearchInput = key

	m.Data.RowsSlice = []con.RowData{}
	// Search the table for the key
	if key == "" {
		slog.Debug("[TableVwDataKey]: key is null")
		m.setTableEndpoints()
		m.setPageSize()
		m.updatePagerLabels(true)
	} else {
		m.setPageButtonState(OFF)
		for x := 0; x < m.Data.TableSize; x++ {
			var row = m.Data.Rows[x]
			index := strings.Index(strings.ToLower(strings.Join(row.Data, " ")), 
				strings.ToLower(key))
			if index != -1 {
				m.Data.RowsSlice = append(m.Data.RowsSlice, row)
			}
		}
	}
	
	slog.Info("Search Key: ", "key", key)
	return m
}

// Process the select event for choosing the number of rows to display in the table
func (m *AppTableVwData) processSelectEvent(w http.ResponseWriter, event con.AppEvent) *AppTableVwData {
	slog.Info("[TableVwData] - processSelectEvent")
	
	data := event.Request.PostForm
	size := data.Get("table_items")

	m.updateSelectedSize(size)
	m.setTableEndpoints()
	m.updatePagerLabels(false)
	return m
}

func (m *AppTableVwData) processViewBtnEvent(event con.AppEvent) *AppTableVwData {
	slog.Info("[TableVwData]: - processViewBtnEvent")
	slog.Debug("[TableVwData]:  ", "value", event.Src)

	if m.isPageEvent {
		m.Data.PageAry = []string{}
		length := len(m.Data.Rows)/m.Data.ShowSize
		for i := 0; i < length; i++ {
			m.Data.PageAry = append(m.Data.PageAry, strconv.Itoa(i + 1))
		}
		if len(m.Data.Rows)%m.Data.ShowSize > 0 {
			m.Data.PageAry = append(m.Data.PageAry, strconv.Itoa(length + 1))
		}
	} else {
		m.setDisplayState(event.Src)
		m.Data.RowsSlice = m.Data.Rows[m.Data.Start:m.Data.End]
	}
	m.isPageEvent = false
	
	return m
}

func (m *AppTableVwData) processPageSelectEvent(event con.AppEvent) *AppTableVwData {
	slog.Info("[TableVwData] - processPageSelectEvent")

	page, _ := strconv.Atoi(event.Src)
	m.Data.Start = (m.Data.ShowSize * (page - 1))
	m.Data.CurrentPage = page

	if (m.Data.Start + m.Data.ShowSize) < m.Data.TableSize {
		m.Data.End = m.Data.Start + m.Data.ShowSize
	} else {
		m.Data.End = m.Data.TableSize
	}

	m.Data.RowsSlice = m.Data.Rows[m.Data.Start:m.Data.End]
	m.updatePagerLabels(false)

	if page == 1 {
		m.setPageButtonState(FWD)
	} else if page == len(m.Data.PageAry) {
		m.setPageButtonState(REV)
	} else {
		m.setPageButtonState(FWD | REV)
	}
	
	return m
}

// Setup the start and end points for the table based on the view size. Also
// update the state of the page controls
func (m *AppTableVwData) setDisplayState(label string) {

	switch label {
	case "sfwd":
		if (m.Data.End + 1) < (m.Data.TableSize)  {
			m.Data.Start = m.Data.Start + 1
			m.Data.End = m.Data.Start + m.Data.ShowSize
			m.setPageButtonState(FWD | REV)
		} else {
			m.Data.Start = m.Data.Start + 1
			m.Data.End = m.Data.TableSize
			m.setPageButtonState(REV)
		}
			
		if m.Data.Start % m.Data.ShowSize == 0 {
			m.Data.CurrentPage = (m.Data.Start/m.Data.ShowSize) + 1
		}			
	case "xfwd":
		if (m.Data.End + m.Data.ShowSize) < (m.Data.TableSize - 1)  {
			m.Data.Start = m.Data.Start + m.Data.ShowSize
			m.Data.End = m.Data.Start + m.Data.ShowSize
			m.setPageButtonState(FWD | REV)
		} else {
			m.Data.Start = m.Data.Start + m.Data.ShowSize
			m.Data.End = m.Data.TableSize
			m.setPageButtonState(REV)
		}			
		
		// Check if start point is on a showsize boundary	
		mod := m.Data.Start % m.Data.ShowSize
		if mod != 0 {
			m.Data.Start = m.Data.Start + (m.Data.ShowSize - mod)		
			m.Data.End = m.Data.Start + m.Data.ShowSize
		} else {
			m.Data.CurrentPage++
		}
	case "srev":
		if (m.Data.Start - 1) > 0 {
			m.Data.Start = m.Data.Start - 1
			m.Data.End = m.Data.Start + m.Data.ShowSize
			m.setPageButtonState(REV | FWD)
		} else {
			m.Data.Start = 0
			m.Data.End = m.Data.ShowSize
			m.setPageButtonState(FWD)
		}
		if m.Data.Start % m.Data.ShowSize == 0 {
			m.Data.CurrentPage = m.Data.Start/m.Data.ShowSize
			if m.Data.CurrentPage == 0 {
				m.Data.CurrentPage = 1
			}		
		} else {
			m.Data.CurrentPage = m.Data.Start/m.Data.ShowSize + 1
		}
	case "xrev":
		mod := m.Data.Start % m.Data.ShowSize
		if mod != 0 {
			
			m.Data.Start = m.Data.Start - mod		
			m.Data.End = m.Data.Start + m.Data.ShowSize

		} else {

			if (m.Data.Start - m.Data.ShowSize) > 0 {
				m.Data.Start = m.Data.Start - m.Data.ShowSize
				m.Data.End = m.Data.Start + m.Data.ShowSize
				m.setPageButtonState(REV | FWD)
			} else {
				m.Data.Start = 0
				m.Data.End = m.Data.ShowSize
				m.setPageButtonState(FWD)
			}

			m.Data.CurrentPage--
			if m.Data.CurrentPage == 0 {
				m.Data.CurrentPage = 1
			}
		}
	}
	m.updatePagerLabels(false)
}

// Set the start and end points for the table to display
func (m *AppTableVwData) setTableEndpoints() {
	m.Data.Start = 0
	if m.Data.ShowSize > m.Data.TableSize {
		m.setPageButtonState(OFF)
		m.Data.End = m.Data.TableSize
	} else {
		m.Data.End = m.Data.ShowSize
		m.setPageButtonState(FWD)
	}
	m.Data.RowsSlice = m.Data.Rows[m.Data.Start:m.Data.End]
	
}

// Set the number of rows to display in the table
func (m *AppTableVwData) updateSelectedSize(size string) {
	for i := 0; i < len(m.Data.OptionValues); i++ {
		if m.Data.OptionValues[i].OptionValue == size {
			m.Data.OptionValues[i].OptionSelected = true
		} else {
			m.Data.OptionValues[i].OptionSelected = false
		}
	}
	m.Data.ShowSize, _ = strconv.Atoi(size)
}

// Set the state of the page buttons
func (m *AppTableVwData) setPageButtonState(state int16) {
		
	for i := 0; i < 5; i++ {
		if 1 << i & state != 0 {
			m.Data.PageBtns[i].Disabled = false
		} else {
			m.Data.PageBtns[i].Disabled = true
		}
	}
	if state > 0 {
		m.Data.Htmx = true
	} else {
		m.Data.Htmx = false
	}
} 

func (m *AppTableVwData) setPageSize() {
	m.Data.PageSize = len(m.Data.Rows)/m.Data.ShowSize
	if len(m.Data.Rows)%m.Data.ShowSize > 0 {
		m.Data.PageSize++
	}
}

func (m *AppTableVwData) updatePagerLabels(reset bool) {
	
	begin := m.Data.Start + 1
	end := m.Data.Start + m.Data.ShowSize + 1
	size := m.Data.TableSize

	if end > size {
		end = size
	}


	m.Data.PageBtns[2].Label = `Row ` + strconv.Itoa(begin) + ` to ` + strconv.Itoa(end) + ` of ` + strconv.Itoa(size)
}





