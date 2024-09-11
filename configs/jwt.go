package configs

type Jwt struct {
	Secret    string
	ExpiresAt string
	Issuer    string
}
