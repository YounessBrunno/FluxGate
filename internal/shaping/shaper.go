package shaping

import "net/http"

type Shaper interface {
	Allow(w http.ResponseWriter, r *http.Request)
}