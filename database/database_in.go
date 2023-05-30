package database

import (
	"context"

	"github.com/jserva90/Erply-test-task/models"
	"github.com/jserva90/Erply-test-task/utils"
)

func (m *SqliteDB) AddSession(clientCode, username, password, sessionKey, sessionToken string) error {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	var session models.Session
	session.ClientCode, _ = utils.HashInput(clientCode)
	session.Username, _ = utils.HashInput(username)
	session.Password, _ = utils.HashInput(password)
	session.SessionKey, _ = utils.HashInput(sessionKey)

	stmt := `INSERT INTO
				 sessions (client_code, username, password, session_key, session_token)
			 VALUES (?,?,?,?,?)`

	_, err := m.DB.ExecContext(ctx, stmt, session.ClientCode, session.Username, session.Password, session.SessionKey, sessionToken)
	if err != nil {
		return err
	}
	return nil
}

func (m *SqliteDB) RemoveSession(sessionToken string) error {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	stmt := `DELETE FROM sessions WHERE session_token = ?`

	_, err := m.DB.ExecContext(ctx, stmt, sessionToken)
	if err != nil {
		return err
	}
	return nil
}
