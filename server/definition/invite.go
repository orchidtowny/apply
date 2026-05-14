package definition

type Invite struct {
	Code    string `json:"code"`
	Creator string `json:"creator"` // ID of user
}
