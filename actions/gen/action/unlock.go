package action

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-masonry/mortar/providers/types"
	"go.uber.org/fx"
	"io/ioutil"
	"log"
	"net/http"
)

var _ = bytes.MinRead
var _ = log.Ldate

type UnlockPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            UnlockArgs             `json:"input"`
}

func makeUnlockHandler(l UnlockLogic) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// set the response header as JSON
		w.Header().Set("Content-Type", "application/json")

		// read request body
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		// parse the body as action payload
		var actionPayload UnlockPayload
		err = json.Unmarshal(reqBody, &actionPayload)
		if err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		// Send the request params to the Action's generated handler function
		result, err := l.Unlock(r.Context(), actionPayload.Input)

		// throw if an error happens
		if err != nil {
			errorObject := GraphQLError{
				Message: err.Error(),
			}
			errorBody, _ := json.Marshal(errorObject)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorBody)
			return
		}

		// Write the response as JSON
		data, _ := json.Marshal(result)
		w.Write(data)
	}
}

type UnlockLogic interface {
	Unlock(ctx context.Context, args UnlockArgs) (response UnlockMutationOutput, err error)
}

type NewUnlockResult struct {
	fx.Out
	types.HTTPHandlerFuncPatternPair `group:"externalHttpHandlerFunctions"`
}

func NewUnlock(l UnlockLogic) (NewUnlockResult, error) {
	if l == nil {
		return NewUnlockResult{}, errors.New("need UnlockLogic")
	}
	return NewUnlockResult{
		HTTPHandlerFuncPatternPair: types.HTTPHandlerFuncPatternPair{
			Pattern:     "/unlock",
			HandlerFunc: makeUnlockHandler(l),
		},
	}, nil
}
