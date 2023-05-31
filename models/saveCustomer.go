package models

type SaveCustomerResponse struct {
	Status  SaveCustomerResponseStatus `json:"status"`
	Records []SaveCustomerRecord       `json:"records"`
}

type SaveCustomerResponseStatus struct {
	Request           string  `json:"request"`
	RequestUnixTime   int64   `json:"requestUnixTime"`
	ResponseStatus    string  `json:"responseStatus"`
	ErrorCode         int     `json:"errorCode"`
	GenerationTime    float64 `json:"generationTime"`
	RecordsTotal      int     `json:"recordsTotal"`
	RecordsInResponse int     `json:"recordsInResponse"`
}

type SaveCustomerRecord struct {
	ClientID      int  `json:"clientID"`
	CustomerID    int  `json:"customerID"`
	AlreadyExists bool `json:"alreadyExists"`
}
