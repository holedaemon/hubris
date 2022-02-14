package types

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Error struct {
	StatusCode int                    `json:"-"`
	Code       int                    `json:"code"`
	Message    string                 `json:"message"`
	Errors     map[string]interface{} `json:"errors"`
}

func (e *Error) Error() string {
	var sb strings.Builder

	if e.Code != 0 {
		sb.WriteString(fmt.Sprintf("discord/types: %d: ", e.Code))
	}

	if e.Message != "" {
		sb.WriteString(e.Message)
	}

	return sb.String()
}

// NewError creates a new Error from an http.Response
func NewError(res *http.Response) error {
	var e *Error
	if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
		return err
	}

	return e
}
