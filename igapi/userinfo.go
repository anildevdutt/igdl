package igapi

import (
	"encoding/json"
	"log"
)

type UserData struct {
	Data struct {
		User struct {
			Biography               string `json:"biography"`
			Biography_with_entities struct {
				Raw_text string `json:"raw_text"`
				Entities []struct {
					User    string `json:"user"`
					Hashtag struct {
						Name string `json:"name"`
					} `json:"hashtag"`
				} `json:"entities"`
			} `json:"biography_with_entities"`
			Blocked_by_viewer        bool   `json:"blocked_by_viewer"`
			Restricted_by_viewer     bool   `json:"restricted_by_viewer"`
			Country_block            bool   `json:"country_block"`
			External_url             string `json:"external_url"`
			External_url_linkshimmed string `json:"external_url_linkshimmed"`
			Edge_followed_by         struct {
				Count int `json:"count"`
			} `json:"edge_followed_by"`
			Fbid               string `json:"fbid"`
			Followed_by_viewer bool   `json:"followed_by_viewer"`
			Edge_follow        struct {
				Count int `json:"count"`
			} `json:"edge_follow"`
			Follows_viewer            bool   `json:"follows_viewer"`
			Full_name                 string `json:"full_name"`
			Has_clips                 bool   `json:"has_clips"`
			Has_channel               bool   `json:"has_channel"`
			Highlight_reel_count      int    `json:"highlight_reel_count"`
			Has_requested_viewer      bool   `json:"has_requested_viewer"`
			Hide_like_and_view_counts bool   `json:"hide_like_and_view_counts"`
			Id                        string `json:"id"`
			Is_business_account       bool   `json:"is_business_account"`
			Is_professional_account   bool   `json:"is_professional_account"`
			Is_supervision_enabled    bool   `json:"is_supervision_enabled"`
			Is_joined_recently        bool   `json:"is_joined_recently"`
			Guardian_id               string `json:"guardian_id"`
			Business_email            string `json:"business_email"`
			Business_phone_number     string `json:"business_phone_number"`
			Business_contact_method   string `json:"business_contact_method"`
			Business_category_name    string `json:"business_category_name"`
			Category_name             string `json:"category_name"`
			Is_private                bool   `json:"is_private"`
			Is_verified               bool   `json:"is_verified"`
			Profile_pic_url_hd        string `json:"profile_pic_url_hd"`
			Username                  string `json:"username"`
			Connected_fb_page         string `json:"connected_fb_page"`
		} `json:"user"`
	} `json:"data"`

	Status string `json:"status"`
}

func GetUserNameInfo(username string) UserData {
	requrl := "https://i.instagram.com/api/v1/users/web_profile_info/?username=" + username
	data := GetRequest(requrl)
	log.Println(string(data))
	var userdata UserData
	err := json.Unmarshal(data, &userdata)
	chkerr(err)
	return userdata
}
