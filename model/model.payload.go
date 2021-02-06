package model

import (
	fcm "github.com/appleboy/go-fcm"
)

type (
	Payload struct {
		ApiKey       string                 `json:"api_key"`
		Topic        string                 `json:"topic"`
		Notification *fcm.Notification      `json:"notification"`
		Data         map[string]interface{} `json:"data"`
	}
)
