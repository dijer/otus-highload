package databus_feed

type PostCreatedEvent struct {
	PostID       int64  `json:"postId"`
	PostText     string `json:"postText"`
	AuthorUserID int64  `json:"author_user_id"`
}
