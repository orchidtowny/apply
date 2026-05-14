package definition

type InfoResponse struct {
	Open  bool              `json:"open"`
	Rules map[string]string `json:"rules"`
}
