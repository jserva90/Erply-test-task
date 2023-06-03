package utils

import (
	"strconv"
	"testing"
	"time"

	"github.com/jserva90/Erply-test-task/models"
)

func TestIsSessionExpired(t *testing.T) {
	currentTime := time.Now().Unix()

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

func TestIsDatabaseCustomerExpired(t *testing.T) {
	currentTime := time.Now().Unix()

	if !IsDatabaseCustomerExpired(currentTime - 600) {
		t.Error("Expected expired timestamp, got not expired")
	}

	if IsDatabaseCustomerExpired(currentTime + 600) {
		t.Error("Expected not expired timestamp, got expired")
	}
}
