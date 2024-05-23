package resiliency

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

// Get - resource GET
func Get(ctx context.Context, h http.Header, values url.Values) ([]Document, *core.Status) {
	return get[core.Log](ctx, h, values)
}

func get[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (docs []Document, status *core.Status) {
	docs, status = getDocuments(ctx, h, values)
	if status.OK() || status.NotFound() || status.Timeout() {
		return
	}
	var e E
	e.Handle(status, core.RequestId(h))
	return
}
