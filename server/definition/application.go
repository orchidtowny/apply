package definition

import "gorm.io/gorm"

type Application struct {
	gorm.Model

	Id                       string `json:"id"`
	Username                 string `json:"username"` // Ingame username
	Age                      int    `json:"age"`
	WhereDidYouFindTheServer string `json:"where_did_you_find_the_server"`
	Bio                      string `json:"bio"`
	Status                   int    `json:"status"` // 0: Pending, 1: Rejected, 2: Approved
}
