package models

type Response struct {
	Status  Status   `json:"status"`
	Records []Record `json:"records"`
}

type Record struct {
	UserID                     string        `json:"userID"`
	UserName                   string        `json:"userName"`
	EmployeeID                 string        `json:"employeeID"`
	EmployeeName               string        `json:"employeeName"`
	GroupID                    string        `json:"groupID"`
	GroupName                  string        `json:"groupName"`
	IPAddress                  string        `json:"ipAddress"`
	SessionKey                 string        `json:"sessionKey"`
	SessionLength              int           `json:"sessionLength"`
	IsPasswordExpired          bool          `json:"isPasswordExpired"`
	LoginURL                   string        `json:"loginUrl"`
	BerlinPOSVersion           string        `json:"berlinPOSVersion"`
	BerlinPOSAssetsURL         string        `json:"berlinPOSAssetsURL"`
	EpsiURL                    string        `json:"epsiURL"`
	RemindUserToUpdateUsername int           `json:"remindUserToUpdateUsername"`
	CustomerRegistryURLs       []string      `json:"customerRegistryURLs"`
	CouponRegistryURLs         []string      `json:"couponRegistryURLs"`
	DisplayAdManagerURLs       []string      `json:"displayAdManagerURLs"`
	EpsiDownloadURLs           []DownloadURL `json:"epsiDownloadURLs"`
	IdentityToken              string        `json:"identityToken"`
	Token                      string        `json:"token"`
}

type DownloadURL struct {
	OperatingSystem string `json:"operatingSystem"`
	URL             string `json:"url"`
}
