package requests

type SystemSettingsRequest struct {
	SetKey   string `json:"set_key"`
	SetValue string `json:"set_value"`
}
