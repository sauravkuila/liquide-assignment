package dto

type Post struct {
	PostId    int64  `json:"postId,omitempty"`
	UserId    int64  `json:"-"`
	UserName  string `json:"userName,omitempty"`
	Content   string `json:"content,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type CreatePostRequest struct {
	Content string `json:"content" binding:"required"`
}

type CreatePostResponse struct {
	Data *Post `json:"data,omitempty"`
	CommonResponse
}

type DeletePostRequest struct {
	PostId int64 `uri:"postId" binding:"required"`
}

type DeletePostResponse struct {
	Data *Post `json:"data,omitempty"`
	CommonResponse
}
