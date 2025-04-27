package session

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"github.com/alexedwards/scs/v2"
)

type SessionService struct {
	SessionMgr *scs.SessionManager
}

var SessionSvc *SessionService

func NewSessionService() *SessionService {
	return &SessionService{}
}

func init() {
	SessionSvc = NewSessionService()
}

func (ss *SessionService) RegisterSessionManager(sessionMgr *scs.SessionManager) {
	ss.SessionMgr = sessionMgr
}

func (ss *SessionService) DeleteSession(sessionID string) error {
	if ss.SessionMgr == nil {
		return fmt.Errorf("session manager is not registered")
	}
	return ss.SessionMgr.Destroy(nil)
}

func (ss *SessionService) DeleteUserSessions(w http.ResponseWriter, r *http.Request, userID string) error {
	err := ss.SessionMgr.Iterate(r.Context(), func(ctx context.Context) error {
		uid := ss.SessionMgr.GetString(ctx, "userID")
	
		if userID == uid {
			return ss.SessionMgr.Destroy(ctx)
		}
	
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

