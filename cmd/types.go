package cmd

type PostData struct {
	Content     string `json:"content"`
	Language    string `json:"language"`
	Title string `json:"title"`
	Images []string `json:"images"`
}

type PostResponse struct {
	AccessKey string `json:"accessKey"`
	Content   string `json:"content"`
	Key       string `json:"key"`
	Language  string `json:"language"`
	TimeStamp int64  `json:"timestamp"`
	Views     int    `json:"views"`
	Title string `json:"title"`
	Images []string `json:"images"`
}

type GetResponse struct {
	Content   string `json:"content"`
	Key       string `json:"key"`
	Language  string `json:"language"`
	Size      int    `json:"size"`
	TimeStamp int64  `json:"timestamp"`
	Views     int    `json:"views"`
	Title string `json:"title"`
}

type DelData struct {
	AccessKey string `json:"accessKey"`
	Key       string `json:"key"`
}

type Pulp struct {
	Key       string `json:"key"`
	AccessKey string `json:"accessKey"`
	TimeStamp int64  `json:"timestamp"`
}
