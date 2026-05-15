package definition

type ApplicationRequest struct {
	Username                 string `json:"username"`
	Age                      int    `json:"age"`
	WhereDidYouFindTheServer string `json:"where_did_you_find_the_server"`
	Bio                      string `json:"bio"`
}
