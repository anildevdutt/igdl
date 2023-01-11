package igapi

import (
	"encoding/json"
	"log"
)

type Followers struct {
	Users []struct {
		PK                            int    `json:"pk"`
		Username                      string `json:"username"`
		Is_verified                   bool   `json:"is_verified"`
		Profile_pic_id                string `json:"profile_pic_id"`
		Profile_pic_url               string `json:"profile_pic_url"`
		Is_private                    bool   `json:"is_private"`
		Pk_id                         string `json:"pk_id"`
		Full_name                     string `json:"full_name"`
		Has_anonymous_profile_picture bool   `json:"has_anonymous_profile_picture"`
		Latest_reel_media             int    `json:"latest_reel_media"`
	} `json:"users"`
	Next_max_id string `json:"next_max_id"`
	Status      string `json:"status"`
}

func GetUserFollowers(userid string) {
	requrl := "https://i.instagram.com/api/v1/friendships/" + userid + "/followers/?count=150&search_surface=follow_list_page"
	data := GetRequest(requrl)

	var followersdata Followers

	err := json.Unmarshal(data, &followersdata)
	chkerr(err)

	log.Println(followersdata)
	// if followersdata.Next_max_id != "" {
	// 	requrl = "https://i.instagram.com/api/v1/friendships/" + userid + "/followers/?count=150&search_surface=follow_list_page&max_id=" + nextmaxid.(string)
	// 	data = GetRequest(requrl)
	// 	log.Println(string(data))

	// }
}
