package resiliency1

type Entry struct {
	Region    string `json:"region"`
	Zone      string `json:"zone"`
	SubZone   string `json:"sub-zone"`
	Host      string `json:"host"`
	Status    string `json:"status"`
	CreatedTS string `json:"created-ts"`
	UpdatedTS string `json:"updated-ts"`

	// Timeout
	Timeout int `json:"timeout"`

	// Rate Limiting
	RateLimit float64 `json:"rate-limit"`
	RateBurst int     `json:"rate-burst"`
}
