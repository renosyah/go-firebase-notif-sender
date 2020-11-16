package model

type (
	Payload struct {
		ApiKey string                 `json:"api_key"`
		Topic  string                 `json:"topic"`
		Data   map[string]interface{} `json:"data"`
	}
)
