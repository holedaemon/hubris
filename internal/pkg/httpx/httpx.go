package httpx

import (
	"net/http"
	"time"
)

func init() {
	http.DefaultClient = &http.Client{
		Timeout: time.Second * 15,
	}
}
