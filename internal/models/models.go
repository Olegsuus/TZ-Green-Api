package models

type Params struct {
	IDInstance       string `json:"idInstance"`
	ApiTokenInstance string `json:"apiTokenInstance"`
	PhoneNumber      string `json:"phoneNumber"`
	Message          string `json:"message"`
	FileUrl          string `json:"fileUrl"`
}

type ParamsFileUrl struct {
	IDInstance       string `json:"idInstance"`
	ApiTokenInstance string `json:"apiTokenInstance"`
	ChatId           string `json:"chatId"`
	UrlFile          string `json:"urlFile"`
	FileName         string `json:"fileName"`
	Caption          string `json:"caption"`
	QuotedMessageId  string `json:"quotedMessageId"`
}
