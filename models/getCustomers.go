package models

type CustomerRecord struct {
	ID                      int           `json:"id"`
	CustomerID              int           `json:"customerID"`
	FullName                string        `json:"fullName"`
	CompanyName             string        `json:"companyName"`
	CompanyTypeID           int           `json:"companyTypeID"`
	FirstName               string        `json:"firstName"`
	LastName                string        `json:"lastName"`
	PersonTitleID           int           `json:"personTitleID"`
	EInvoiceEmail           string        `json:"eInvoiceEmail"`
	EInvoiceReference       string        `json:"eInvoiceReference"`
	EmailEnabled            int           `json:"emailEnabled"`
	EInvoiceEnabled         int           `json:"eInvoiceEnabled"`
	DocuraEDIEnabled        int           `json:"docuraEDIEnabled"`
	MailEnabled             int           `json:"mailEnabled"`
	OperatorIdentifier      string        `json:"operatorIdentifier"`
	EDI                     string        `json:"EDI"`
	DoNotSell               int           `json:"doNotSell"`
	PartialTaxExemption     int           `json:"partialTaxExemption"`
	GroupID                 int           `json:"groupID"`
	CountryID               string        `json:"countryID"`
	PayerID                 int           `json:"payerID"`
	Phone                   string        `json:"phone"`
	Mobile                  string        `json:"mobile"`
	Email                   string        `json:"email"`
	Fax                     string        `json:"fax"`
	Code                    string        `json:"code"`
	Birthday                string        `json:"birthday"`
	IntegrationCode         string        `json:"integrationCode"`
	FlagStatus              int           `json:"flagStatus"`
	ColorStatus             string        `json:"colorStatus"`
	Credit                  int           `json:"credit"`
	SalesBlocked            int           `json:"salesBlocked"`
	ReferenceNumber         string        `json:"referenceNumber"`
	CustomerCardNumber      string        `json:"customerCardNumber"`
	FactoringContractNumber string        `json:"factoringContractNumber"`
	GroupName               string        `json:"groupName"`
	CustomerType            string        `json:"customerType"`
	Address                 string        `json:"address"`
	Street                  string        `json:"street"`
	Address2                string        `json:"address2"`
	City                    string        `json:"city"`
	PostalCode              string        `json:"postalCode"`
	Country                 string        `json:"country"`
	State                   string        `json:"state"`
	AddressTypeID           int           `json:"addressTypeID"`
	AddressTypeName         string        `json:"addressTypeName"`
	IsPOSDefaultCustomer    int           `json:"isPOSDefaultCustomer"`
	EUCustomerType          string        `json:"euCustomerType"`
	EDIType                 string        `json:"ediType"`
	LastModifierUsername    string        `json:"lastModifierUsername"`
	LastModifierEmployeeID  int           `json:"lastModifierEmployeeID"`
	TaxExempt               int           `json:"taxExempt"`
	PaysViaFactoring        int           `json:"paysViaFactoring"`
	RewardPoints            int           `json:"rewardPoints"`
	TwitterID               string        `json:"twitterID"`
	FacebookName            string        `json:"facebookName"`
	CreditCardLastNumbers   string        `json:"creditCardLastNumbers"`
	GLN                     string        `json:"GLN"`
	DeliveryTypeID          int           `json:"deliveryTypeID"`
	Image                   string        `json:"image"`
	CustomerBalanceDisabled int           `json:"customerBalanceDisabled"`
	RewardPointsDisabled    int           `json:"rewardPointsDisabled"`
	POSCouponsDisabled      int           `json:"posCouponsDisabled"`
	EmailOptOut             int           `json:"emailOptOut"`
	SignUpStoreID           int           `json:"signUpStoreID"`
	HomeStoreID             int           `json:"homeStoreID"`
	Gender                  string        `json:"gender"`
	PeppolID                string        `json:"PeppolID"`
	ExternalIDs             []interface{} `json:"externalIDs"`
}

type CustomerResponse struct {
	Status  Status           `json:"status"`
	Records []CustomerRecord `json:"records"`
}
