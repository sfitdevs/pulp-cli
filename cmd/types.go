package cmd

type PostData struct {
	Content  string `json:"content"`
	Language string `json:"language"`
}

type PostResponse struct {
	AccessKey string `json:"accessKey"`
	Content   string `json:"content"`
	Key       string `json:"key"`
	Language  string `json:"language"`
	Size      int    `json:"size"`
	TimeStamp int64  `json:"timestamp"`
	Views     int    `json:"views"`
}

type GetResponse struct {
	Content   string `json:"content"`
	Key       string `json:"key"`
	Language  string `json:"language"`
	Size      int    `json:"size"`
	TimeStamp int64  `json:"timestamp"`
	Views     int    `json:"views"`
}

type DelData struct {
	AccessKey string `json:"accessKey"`
}

type Pulp struct {
	Key       string `json:"key"`
	AccessKey string `json:"accessKey"`
}
