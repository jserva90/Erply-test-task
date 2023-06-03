package models

type Status struct {
	Request            string  `json:"request"`
	RequestUnixTime    int64   `json:"requestUnixTime"`
	ResponseStatus     string  `json:"responseStatus"`
	ErrorCode          int     `json:"errorCode"`
	GenerationTime     float64 `json:"generationTime"`
	RecordsTotal       int     `json:"recordsTotal"`
	RecordsInResponse  int     `json:"recordsInResponse"`
	RequestFromLocalDB bool    `json:"requestFromLocalDB"`
}
