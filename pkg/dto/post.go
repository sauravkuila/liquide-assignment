package dto

type Post struct {
	PostId    int64  `json:"postId,omitempty"`
	UserId    int64  `json:"-"`
	UserName  string `json:"userName,omitempty"`
	Content   string `json:"content,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type PostInfo struct {
	PostId        int64  `json:"postId,omitempty"`
	UserName      string `json:"userName,omitempty"`
	Content       string `json:"content,omitempty"`
	UpVoteCount   int64  `json:"upVoteCount,omitempty"`
	DownVoteCount int64  `json:"downVoteCount,omitempty"`
	CommentCount  int64  `json:"commentCount,omitempty"`
	CreatedAt     string `json:"createdAt,omitempty"`
	UpdatedAt     string `json:"updatedAt,omitempty"`
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

type EditPostRequest struct {
	PostId  int64  `json:"postId" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type EditPostResponse struct {
	Data *Post `json:"data,omitempty"`
	CommonResponse
}

type GetPostRequest struct {
	PostId int64 `uri:"postId" binding:"required"`
}

type GetPostResponse struct {
	Data *PostInfo `json:"data,omitempty"`
	CommonResponse
}

type GetAllUserPostRequest struct {
	Page     int `form:"page" binding:"required,gte=0"`
	PageSize int `form:"pageSize" binding:"required,gte=0,min=10"`
}

type GetAllUserPostResponse struct {
	Data       []PostInfo `json:"data,omitempty"`
	Page       int        `json:"page,omitempty"`
	PageSize   int        `json:"pageSize,omitempty"`
	TotalCount int64      `json:"totalCount,omitempty"`
	CommonResponse
}
