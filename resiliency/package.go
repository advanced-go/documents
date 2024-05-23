package resiliency

import (
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

const (
	PkgPath = "github/advanced-go/documents/resiliency"
)

type Document struct {
	Origin  core.Origin
	Version string
	Content string
}

type PutBodyConstraints interface {
	[]Document | []byte | *http.Request
}
