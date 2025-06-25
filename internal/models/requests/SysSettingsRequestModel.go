package requests

type SystemSettingsRequest struct {
	SetKey   string `json:"set_key"`
	SetValue int    `json:"set_value"`
}
