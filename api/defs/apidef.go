package defs

//reqeusts
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

type NewComment struct {
	AuthorId int `json:"author_id"`
	Content string `json:"content"`
}

type NewVideo struct {
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
}

//response
type SignedUp struct {
	Success bool `json:"user_name"`
	SessionId string `json:"session_id"`
}

type UserSession struct {
	Username string `json:"user_name"`
	SessionId string `json:"session_id"`
}

type UserInfo struct {
	Id int `json:"id"`
}

type SignedIn struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

type VideosInfo struct {
	Videos []*VideoInfo `json:"videos"`
}

type Comments struct {
	Comments []*Comment `json:"comments"`
}

//Data model
//视频结构体
type VideoInfo struct {
	Id string `json:"id"`
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

//评论结构体
type Comment struct {
	Id string `json:"id"`
	VideoId string `json:"video_id"`
	Author string `json:"author"`
	Content string `json:"content"`
}

//用户结构体
type User struct {
	Id int
	LoginName string
	Pwd string
}

//session结构体
type SimpleSession struct {
	Username string //login name
	TTL int64
}