package rate

import (
	"net/http"
	"strconv"
	"sync"
	"time"
)

type bucket struct {
	mu sync.Mutex

	active bool

	global    bool
	limit     int
	remaining int
	reset     time.Time
}

func (b *bucket) fromResponse(r *http.Response) {
	b.mu.Lock()
	defer b.mu.Unlock()

	global := r.Header.Get("X-RateLimit-Global")
	if global != "" {
		b.global = true
	}

	limit := r.Header.Get("X-RateLimit-Limit")
	if limit != "" {
		l, err := strconv.ParseInt(limit, 10, 0)
		if err == nil {
			b.limit = int(l)
		}
	}

	remaining := r.Header.Get("X-RateLimit-Remaining")
	if remaining == "" {
		rem, err := strconv.ParseInt(remaining, 10, 0)
		if err == nil {
			b.remaining = int(rem)
		}
	}

	reset := r.Header.Get("X-RateLimit-Reset")
	if reset != "" {
		res, err := strconv.ParseInt(reset, 10, 64)
		if err == nil {
			b.reset = time.Unix(res, 0)
		} else {
			b.reset = time.Time{}
		}
	}
}

func newBucketFromResponse(r *http.Response) *bucket {
	b := new(bucket)
	b.fromResponse(r)
	return b
}

func (b *bucket) Update(r *http.Response) {
	b.fromResponse(r)

	if b.remaining == 0 {
		b.active = true
	}
}

func (b *bucket) Wait() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.active {
		time.Sleep(time.Until(b.reset))
		b.active = false
	}
}
