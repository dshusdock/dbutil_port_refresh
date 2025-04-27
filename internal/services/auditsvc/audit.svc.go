package auditsvc

import (
	"context"
	"dshusdock/go_project/internal/services/session"
	"dshusdock/go_project/internal/views/auditvw"
	"log/slog"
	"strconv"
)

type AuditService struct {
}

var AuditSvc *AuditService

func NewAuditService() *AuditService {
	return &AuditService{}
}

func init() {
	AuditSvc = NewAuditService()
}

func (m *AuditService) GetSQLString(ctx context.Context, val string) (string, string) {
	var obj auditvw.AppAuditVwData

	if session.SessionSvc.SessionMgr.Exists(ctx, "AuditVw") {
		obj = session.SessionSvc.SessionMgr.Get(ctx, "AuditVw").(auditvw.AppAuditVwData)

	}

	id, err := strconv.Atoi(val)
	if err != nil {
		slog.Error("Error converting row.Name to int: ", "Error", err.Error())
		return "", ""
	}
	return obj.Data[id-1].Query, obj.Data[id-1].IssueName
}

func (m *AuditService) GetAuditSQLString(ctx context.Context) string{
	var obj auditvw.AppAuditVwData
	var strAry []string
	finalStr := "" 

	if session.SessionSvc.SessionMgr.Exists(ctx, "AuditVw") {
		obj = session.SessionSvc.SessionMgr.Get(ctx, "AuditVw").(auditvw.AppAuditVwData)

	}

	for x:=0; x<len(obj.Data); x++ {
		item := obj.Data[x]
		if item.Checked{
			strAry = append(strAry, item.IssueDef)
		}		
	}

	for i, str := range strAry {
		finalStr += str
		if i == len(strAry)-1 {
          finalStr += ";"
        } else {
           finalStr += " union "
        }
	}

	return finalStr
}