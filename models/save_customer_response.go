package models

type SaveCustomerResponse struct {
	Status  Status               `json:"status"`
	Records []SaveCustomerRecord `json:"records"`
}

type SaveCustomerRecord struct {
	ClientID      int  `json:"clientID"`
	CustomerID    int  `json:"customerID"`
	AlreadyExists bool `json:"alreadyExists"`
}
