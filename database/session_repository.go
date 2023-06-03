package database

import (
	"context"
	"database/sql"

	"github.com/jserva90/Erply-test-task/models"
	"github.com/jserva90/Erply-test-task/utils"
)

func (m *SqliteDB) AddSession(clientCode, username, password, sessionKey, sessionToken string) error {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	var session models.Session
	session.ClientCode, _ = utils.Encrypt(clientCode, utils.GetSecretKey())
	session.Username = username
	session.Password, _ = utils.Encrypt(password, utils.GetSecretKey())
	session.SessionKey, _ = utils.Encrypt(sessionKey, utils.GetSecretKey())

	checkStmt := `SELECT EXISTS(SELECT 1 FROM sessions WHERE username = ? LIMIT 1)`
	var exists bool
	err := m.DB.QueryRowContext(ctx, checkStmt, session.Username).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	var stmt string
	if exists {
		stmt = `UPDATE sessions
			SET client_code = ?, username = ?, password = ?, session_key = ?, session_token = ?
			WHERE username = ?`
	} else {
		stmt = `INSERT INTO sessions (client_code, username, password, session_key, session_token)
			VALUES (?,?,?,?,?)`
	}

	var args []interface{}
	if exists {
		args = []interface{}{session.ClientCode, session.Username, session.Password, session.SessionKey, sessionToken, session.Username}
	} else {
		args = []interface{}{session.ClientCode, session.Username, session.Password, session.SessionKey, sessionToken}
	}

	_, err = m.DB.ExecContext(ctx, stmt, args...)
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

func (m *SqliteDB) GetSession(sessionToken string) (*models.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	query := `SELECT client_code, username, password, session_key FROM sessions WHERE session_token = ?`

	row := m.DB.QueryRowContext(ctx, query, sessionToken)
	var session models.Session

	err := row.Scan(
		&session.ClientCode,
		&session.Username,
		&session.Password,
		&session.SessionKey,
	)
	if err != nil {
		return nil, err
	}

	return &session, nil
}
