package router

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/renosyah/go-firebase-notif-sender/api"
	"github.com/renosyah/go-firebase-notif-sender/model"
)

func HandlerAddPayload(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.Payload

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Payload/create/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return payloadModule.Add(ctx, param)
}
