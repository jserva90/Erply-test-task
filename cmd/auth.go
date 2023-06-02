package main

import (
	"fmt"

	"github.com/jserva90/Erply-test-task/utils"
)

func (app *application) getClientCodeAndSessionKey(sessionCookie string) (string, string, error) {
	session, err := app.DB.GetSession(sessionCookie)
	if err != nil {
		return "", "", fmt.Errorf("failed to get session: %w", err)
	}

	decryptedClientCode, err := utils.Decrypt(session.ClientCode, utils.GetSecretKey())
	if err != nil {
		return "", "", fmt.Errorf("failed to decrypt client code: %w", err)
	}

	decryptedSessionKey, err := utils.Decrypt(session.SessionKey, utils.GetSecretKey())
	if err != nil {
		return "", "", fmt.Errorf("failed to decrypt session key: %w", err)
	}

	sessionInfo, err := app.getSessionKeyInfo(decryptedClientCode, decryptedSessionKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to get session info: %w", err)
	}

	if utils.IsSessionExpired(*sessionInfo) {
		return "", "", fmt.Errorf("session expired")
	}

	return decryptedClientCode, decryptedSessionKey, nil
}
