package utils

import (
	"strconv"
	"time"

	"github.com/jserva90/Erply-test-task/models"
)

func IsSessionExpired(sessionInfo models.GetSessionKeyInfoResponse) bool {
	currentTime := time.Now().Unix()
	convertedExpiryUnixTime, _ := strconv.ParseInt(sessionInfo.Records[0].ExpireUnixTime, 10, 64)
	return currentTime > convertedExpiryUnixTime
}

func IsDatabaseCustomerExpired(timestamp int64) bool {
	expirationTime := time.Unix(timestamp, 0).Add(10 * time.Minute)
	return time.Now().After(expirationTime)
}
