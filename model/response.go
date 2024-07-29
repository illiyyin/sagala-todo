package model

type GETAllTaskResponse struct {
	Data    []TaskResponse `json:"data"`
	Message string         `json:"message"`
}

type GETTaskResponse struct {
	Data    TaskResponse `json:"data"`
	Message string       `json:"message"`
}

type DELETETaskResponse struct {
	Message string       `json:"message"`
}

type PATCHTaskResponse struct {
	Data    TaskResponse `json:"data"`
	Message string       `json:"message"`
}

type POSTTaskResponse struct {
	Data    TaskResponse `json:"data"`
	Message string       `json:"message"`
}