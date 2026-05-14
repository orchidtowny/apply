package definition

type Invite struct {
	Id      string `json:"id"`
	Code    string `json:"code"`
	Creator string `json:"creator"` // ID of user
}
