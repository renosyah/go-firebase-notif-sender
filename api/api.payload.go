package api

import (
	"context"
	"fmt"
	"net/http"

	fcm "github.com/appleboy/go-fcm"
	"github.com/renosyah/go-firebase-notif-sender/model"
)

type (
	PayloadModule struct {
		Name string
	}
)

func NewPayloadModule() *PayloadModule {
	return &PayloadModule{
		Name: "module/Payload",
	}
}

func (m PayloadModule) Add(ctx context.Context, param model.Payload) (interface{}, *Error) {

	msg := &fcm.Message{
		Condition:    fmt.Sprintf(`'%s' in topics`, param.Topic),
		Data:         param.Data,
		Notification: param.Notification,
	}

	client, err := fcm.NewClient(param.ApiKey)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on sending notif"

		return map[string]string{"status": "not ok"}, NewErrorWrap(err, m.Name, "add/payload",
			message, status)
	}

	_, err = client.Send(msg)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on sending notif"

		return map[string]string{"status": "not ok"}, NewErrorWrap(err, m.Name, "add/payload",
			message, status)
	}

	return map[string]string{"status": "ok"}, nil
}
