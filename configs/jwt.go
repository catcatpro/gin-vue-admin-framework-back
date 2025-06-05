package configs

type Jwt struct {
	Secret           string
	RefreshSecret    string
	ExpiresAt        string
	RefreshExpiresAt string
	Issuer           string
}
