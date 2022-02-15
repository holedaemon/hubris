package exp

import (
	"math/rand"
	"time"
)

type Backoff struct {
	base time.Duration
	max  time.Duration
	f    float64
	j    float64
}

func New(base, max time.Duration, f, j float64) *Backoff {
	return &Backoff{
		base: base,
		max:  max,
		f:    f,
		j:    j,
	}
}

func (b *Backoff) Attempt(n int) time.Duration {
	if n == 0 {
		return b.base
	}

	bkf := float64(b.base)
	max := float64(b.max)

	for bkf < max && n > 0 {
		bkf *= b.f
		n--
	}

	if bkf > max {
		bkf = max
	}

	bkf *= 1 + b.j*(rand.Float64()*2-1)

	if bkf < 0 {
		return 0
	}

	return time.Duration(bkf)
}
