package apilog

type ApiLogData struct {
	Body string `json:"body"`
	Header string `json:"header"`
	Url string `json:"url"`
	Ip string `json:"ip"`
	ServerType string `json:"serverType"`
	Error string `json:"error"`
}