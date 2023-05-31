package database

import (
	"context"

	"github.com/jserva90/Erply-test-task/models"
)

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
