package models

// Response is http response
type Response struct {
	Response responseTime
	Status   status
	Data     map[string]interface{}
}

// ResponseNodata is http response
type ResponseNodata struct {
	Response responseTime
	Status   status
}

type responseTime struct {
	Datetime string
}

type status struct {
	Code        int
	Description string
	Message     string
}
