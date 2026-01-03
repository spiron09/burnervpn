package store

import (
	"sync"

	"github.com/spiron09/burnervpn/server/models"
)

type SessionStore struct {
	sessions map[string]*models.Session
	mu       sync.RWMutex
}

func SessionStoreInit() *SessionStore {
	return &SessionStore{
		sessions: make(map[string]*models.Session),
	}
}

func (s *SessionStore) GetSession(id string) *models.Session {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.sessions[id]
}

func (s *SessionStore) SetSession(session *models.Session) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sessions[session.ID] = session
}

func (s *SessionStore) DeleteSession(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sessions, id)
}
