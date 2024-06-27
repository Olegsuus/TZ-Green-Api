package models

type Params struct {
	IDInstance       string `json:"idInstance"`
	ApiTokenInstance string `json:"apiTokenInstance"`
	PhoneNumber      string `json:"phoneNumber"`
	Message          string `json:"message"`
	FileUrl          string `json:"fileUrl"`
}
