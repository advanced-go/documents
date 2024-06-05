package resiliency1

import (
	"context"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
	"time"
)

var storage []Entry

func getDocuments(_ context.Context, req access.Request, values url.Values) (docs []Entry, h2 http.Header, status *core.Status) {
	if len(storage) == 0 {
		return nil, nil, core.StatusNotFound()
	}
	var start = time.Now().UTC()
	h2 = make(http.Header)

	filter := core.NewOrigin(values)
	for _, item := range storage {
		target := core.Origin{Region: item.Region, Zone: item.Zone, SubZone: item.SubZone, Host: item.Host}
		if core.OriginMatch(target, filter) {
			docs = append(docs, item)
		}
	}
	if len(docs) == 0 {
		status = core.StatusNotFound()
	} else {
		status = core.StatusOK()
	}
	access.LogEgress(start, time.Since(start), req, status, routeName, "", timeout, "")
	return docs, h2, status
}

func addDocuments(_ context.Context, req access.Request, docs []Entry) (http.Header, *core.Status) {
	var start = time.Now().UTC()
	h2 := make(http.Header)

	if len(docs) > 0 {
		storage = append(storage, docs...)
	}
	access.LogEgress(start, time.Since(start), req, core.StatusOK(), routeName, "", timeout, "")
	return h2, core.StatusOK()
}

func setTimeout(ctx context.Context, duration time.Duration) (context.Context, context.CancelFunc) {
	if ctx == nil {
		ctx = context.Background()
	}
	if _, ok := ctx.Deadline(); ok {
		return ctx, nil
	}
	return context.WithTimeout(ctx, duration)
}
