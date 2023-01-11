package igapi

import (
	"encoding/json"
	"log"
)

type Following struct {
	Users []struct {
		PK                            int    `json:"pk"`
		Username                      string `json:"username"`
		Is_verified                   bool   `json:"is_verified"`
		Profile_pic_id                string `json:"profile_pic_id"`
		Profile_pic_url               string `json:"profile_pic_url"`
		Full_name                     string `json:"full_name"`
		Pk_id                         string `json:"pk_id"`
		Is_private                    bool   `json:"is_private"`
		Has_anonymous_profile_picture bool   `json:"has_anonymous_profile_picture"`
		Latest_reel_media             int    `json:"latest_reel_media"`
		Is_favorite                   bool   `json:"is_favorite"`
	} `json:"users"`
	Next_max_id string `json:"next_max_id"`
	Status      string `json:"status"`
}

func GetUserFollowing(userid string) {
	requrl := "https://i.instagram.com/api/v1/friendships/" + userid + "/following/?count=100"
	data := GetRequest(requrl)
	var followingdata Following

	err := json.Unmarshal(data, &followingdata)
	chkerr(err)

	log.Println(followingdata)
	// 	https://i.instagram.com/api/v1/friendships/54353687036/following/?count=12&max_id=12
}
