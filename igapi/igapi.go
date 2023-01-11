package igapi

import (
	"log"
)

func chkerr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func GetUserIDInfo(userid string) {
	requrl := "https://i.instagram.com/api/v1/users/" + userid + "/info/"
	data := GetRequest(requrl)
	log.Println(string(data))
}

func GetPostComments() {

}

func GetPostLikes() {
	// https://i.instagram.com/api/v1/media/3012208256132025812/likers/
}

func GetPostInfo() {

}

func GetUserPosts() {
	// https://i.instagram.com/api/v1/feed/user/theamyacker/username/?count=12
	// https://www.instagram.com/graphql/query/?query_hash=69cba40317214236af40e7efa697781d&variables=%7B%22id%22%3A%224122113849%22%2C%22first%22%3A12%2C%22after%22%3A%22QVFDV3YwemdKdUg1cWY1VmNCWDByTUtWWmZkaXkwNlU1bXVNUlJKTkRUN01XdmdpTnI2X21uazBwWjNsNUpyd0dHdUhqRFU2LUxKTmhxU09qMHRUU0JRTw%3D%3D%22%7D
}
