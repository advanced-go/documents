package resiliency1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

var storage []Entry

func getDocuments(_ context.Context, h http.Header, values url.Values) (docs []Entry, status *core.Status) {
	if len(storage) == 0 {
		return nil, core.StatusNotFound()
	}
	filter := core.NewOrigin(values)
	for _, item := range storage {
		target := core.Origin{Region: item.Region, Zone: item.Zone, SubZone: item.SubZone, Host: item.Host}
		if core.OriginMatch(target, filter) {
			docs = append(docs, item)
		}
	}
	if len(docs) == 0 {
		return nil, core.StatusNotFound()
	}
	return docs, core.StatusOK()
}

func addDocuments(_ context.Context, _ http.Header, docs []Entry) *core.Status {
	if len(docs) > 0 {
		storage = append(storage, docs...)
	}
	return core.StatusOK()
}
