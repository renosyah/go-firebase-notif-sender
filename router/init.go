package router

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/renosyah/go-firebase-notif-sender/api"
)

var (
	dbPool        *sql.DB
	payloadModule *api.PayloadModule
)

func Init() {
	payloadModule = api.NewPayloadModule()
}

func ParseBodyData(ctx context.Context, r *http.Request, data interface{}) error {
	bBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(err, "read")
	}

	err = json.Unmarshal(bBody, data)
	if err != nil {
		return errors.Wrap(err, "json")
	}

	valid, err := govalidator.ValidateStruct(data)
	if err != nil {
		return errors.Wrap(err, "validate")
	}

	if !valid {
		return errors.New("invalid data")
	}

	return nil
}
