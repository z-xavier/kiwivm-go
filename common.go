package kiwivm

const (
	Host = "https://api.64clouds.com"
)

const (
	plainTextType = "text/plain;charset=UTF-8"
)

const (
	EnvVeID   = "KIWIVM_VEID"
	EnvApiKey = "KIWIVM_API_KEY"
)

type Auth struct {
	VeID   string `url:"veid"`
	APIKey string `url:"api_key"`
}

type ErrorRsp struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}
