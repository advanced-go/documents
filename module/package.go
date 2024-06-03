package module

const (
	Authority          = "github/advanced-go/documents"
	Name               = "documents"
	Version            = "1.1.1"
	Ver1               = "v1"
	Ver2               = "v2"
	ResiliencyResource = "resiliency"
)

// Configuration keys used on startup for map values
const (
	PackageNameUserKey     = "user"    // type:text, empty:false
	PackageNamePasswordKey = "pswd"    // type:text, empty:false
	PackageNameRetriesKey  = "retries" // type:int, default:-1, range:0-100
)
