package utils

import (
	"strconv"
	"testing"
	"time"

	"github.com/jserva90/Erply-test-task/models"
)

func TestIsSessionExpired(t *testing.T) {
	currentTime := time.Now().Unix()

	// Create a sample sessionInfo with expiry time 10 seconds in the past
	expiryUnixTime := strconv.FormatInt(currentTime-10, 10)
	sessionInfo := models.GetSessionKeyInfoResponse{
		Records: []models.SessionKeyRecord{
			{
				ExpireUnixTime: expiryUnixTime,
			},
		},
	}

	isExpired := IsSessionExpired(sessionInfo)

	if !isExpired {
		t.Errorf("Expected session to be expired, but got not expired")
	}

	// Create a sample sessionInfo with expiry time 10 seconds in the future
	expiryUnixTime = strconv.FormatInt(currentTime+10, 10)
	sessionInfo = models.GetSessionKeyInfoResponse{
		Records: []models.SessionKeyRecord{
			{
				ExpireUnixTime: expiryUnixTime,
			},
		},
	}

	isExpired = IsSessionExpired(sessionInfo)

	if isExpired {
		t.Errorf("Expected session not to be expired, but got expired")
	}
}
