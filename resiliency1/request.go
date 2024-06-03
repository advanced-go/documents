package resiliency1

import (
	"fmt"
	"net/http"
	"time"
)

type request struct {
	method    string
	resource  string
	uri       string
	routeName string
	duration  time.Duration
	h         http.Header
}

func NewRequest(method, resource, routeName string, h http.Header, duration time.Duration) *request {
	r := new(request)
	r.method = method
	r.resource = resource
	r.routeName = routeName
	r.h = h
	r.duration = duration
	r.uri = fmt.Sprintf("documents://host-name/%v", resource)
	return r
}
