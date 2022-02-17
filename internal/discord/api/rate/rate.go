package rate

import (
	"net/http"
	"sync"
)

type Limiter struct {
	mu sync.Mutex

	global  *bucket
	buckets map[string]*bucket
}

func New() *Limiter {
	return &Limiter{
		global:  new(bucket),
		buckets: make(map[string]*bucket),
	}
}

func (l *Limiter) update(r *http.Response) {
	global := r.Header.Get("X-RateLimit-Global")
	if global != "" {
		l.global.Update(r)
	} else {
		bucket := r.Header.Get("X-RateLimit-Bucket")
		if bucket != "" {
			b, ok := l.buckets[bucket]
			if !ok {
				b = newBucketFromResponse(r)
				l.buckets[bucket] = b
			} else {
				b.Update(r)
			}
		}
	}
}

func (l *Limiter) wait(r *http.Response) {
	global := r.Header.Get("X-RateLimit-Global")
	if global != "" {
		l.global.Wait()
	} else {
		bkt := r.Header.Get("X-RateLimit-Bucket")
		b := l.buckets[bkt]
		b.Wait()
	}
}

// Wait updates the Limiter's state and sleeps if necessary.
func (l *Limiter) Wait(r *http.Response) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.update(r)
	l.wait(r)
}
