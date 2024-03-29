package bot

import (
	"database/sql"
	"log"

	"github.com/anildevdutt/igdl/igapi"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var dbfile = "db/igdl.db"

func chkerr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func SaveUserData(username string) {
	userdata := igapi.GetUserNameInfo(username)
	user := userdata.Data.User
	insert := "INSERT INTO users VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	_, err := DB.Exec(insert, user.Id, user.Username, user.Full_name, user.Biography, user.External_url,
		user.Edge_followed_by.Count, user.Edge_follow.Count, user.Fbid, user.Has_clips, user.Has_channel, user.Highlight_reel_count,
		user.Hide_like_and_view_counts, user.Is_business_account, user.Is_professional_account, user.Is_supervision_enabled,
		user.Is_joined_recently, user.Guardian_id, user.Business_email, user.Business_phone_number, user.Business_contact_method,
		user.Business_category_name, user.Category_name, user.Is_private, user.Is_verified, user.Profile_pic_url_hd, user.Connected_fb_page)
	chkerr(err)
}

func initDB() {
	var err error
	DB, err = sql.Open("sqlite3", dbfile)
	chkerr(err)
}

func SaveAllUserFollowers(userid string) {

}

func Start() {
	initDB()

	SaveUserData("f1")
}
