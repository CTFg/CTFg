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

type SubmitPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            SubmitArgs             `json:"input"`
}

func makeSubmitHandler(l SubmitLogic) http.HandlerFunc {
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
		var actionPayload SubmitPayload
		err = json.Unmarshal(reqBody, &actionPayload)
		if err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		// Send the request params to the Action's generated handler function
		result, err := l.Submit(r.Context(), actionPayload.Input)

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

type SubmitLogic interface {
	Submit(ctx context.Context, args SubmitArgs) (response SubmitAttemptMutationOutput, err error)
}

type NewSubmitResult struct {
	fx.Out
	types.HTTPHandlerFuncPatternPair `group:"externalHttpHandlerFunctions"`
}

func NewSubmit(l SubmitLogic) (NewSubmitResult, error) {
	if l == nil {
		return NewSubmitResult{}, errors.New("need SubmitLogic")
	}
	return NewSubmitResult{
		HTTPHandlerFuncPatternPair: types.HTTPHandlerFuncPatternPair{
			Pattern:     "/submit",
			HandlerFunc: makeSubmitHandler(l),
		},
	}, nil
}
