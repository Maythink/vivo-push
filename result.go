package vivopush

type ResultItem struct {
	Result int    `json:"result"`
	Desc   string `json:"desc"`
}

type SendResult struct {
	ResultItem
	TaskId string `json:"taskId"`
}

type BatchStatusResult struct {
	ResultItem
	statistics []TaskData `json:"statistics"`
}

type TaskData struct {
	TaskId  string `json:"taskId"`
	Send    int    `json:"send"`
	Receive int    `json:"receive"`
	Display int    `json:"display"`
	Click   int    `json:"click"`
}
