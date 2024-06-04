package resiliency1

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

const (
	PkgPath   = "github/advanced-go/documents/resiliency1"
	routeName = "documents"
	timeout   = 500
)

func errorInvalidURL(path string) *core.Status {
	return core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("invalid argument: URL path is invalid %v", path)))
}

func buildURL(values url.Values) string {
	if values == nil {
		return fmt.Sprintf("docs://docs-host.com/documents/resiliency")
	}
	return fmt.Sprintf("docs://docs-host.com/documents/resiliency?%v", values.Encode())
}

// Get - resource GET
func Get(ctx context.Context, h http.Header, values url.Values) ([]Entry, *core.Status) {
	return getDocuments(ctx, access.NewRequest(http.MethodGet, buildURL(values), core.AddRequestId(h), routeName, timeout), values)
}

// Put - resource PUT, with optional content override
func Put(r *http.Request, body []Entry) *core.Status {
	if r == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid argument: request is nil"))
	}
	if body == nil {
		content, status := json2.New[[]Entry](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			return status
		}
		body = content
	}
	return addDocuments(r.Context(), access.NewRequest(r.Method, buildURL(nil), core.AddRequestId(r.Header), routeName, timeout), body)
}
