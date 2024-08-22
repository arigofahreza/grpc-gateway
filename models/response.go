package models

type ResponseModel struct {
	Metadata Metadata    `json:"metadata"`
	Data     interface{} `json:"data"`
}

type Metadata struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
