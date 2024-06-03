package resiliency1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/url"
	"time"
)

var storage []Entry

func getDocuments(_ context.Context, req *request, values url.Values) (docs []Entry, status *core.Status) {
	if len(storage) == 0 {
		return nil, core.StatusNotFound()
	}
	var start = time.Now().UTC()

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
	log(start, time.Since(start), req, status, "")
	return docs, core.StatusOK()
}

func addDocuments(_ context.Context, req *request, docs []Entry) *core.Status {
	var start = time.Now().UTC()

	if len(docs) > 0 {
		storage = append(storage, docs...)
	}
	log(start, time.Since(start), req, core.StatusOK(), "")
	return core.StatusOK()
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
