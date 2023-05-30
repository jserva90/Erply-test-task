package models

type GetSessionKeyInfoResponse struct {
	Status  Status             `json:"status"`
	Records []SessionKeyRecord `json:"records"`
}

type SessionKeyRecord struct {
	CreationUnixTime string `json:"creationUnixTime"`
	ExpireUnixTime   string `json:"expireUnixTime"`
}
