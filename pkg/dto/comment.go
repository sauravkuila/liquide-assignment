package dto

type Comment struct {
	CommentId int64     `json:"commentId,omitempty"`
	PostId    int64     `json:"postId,omitempty"`
	UserId    int64     `json:"-"`
	UserName  string    `json:"userName,omitempty"`
	Content   string    `json:"content,omitempty"`
	Replies   []Comment `json:"replies,omitempty"`
	CreatedAt string    `json:"createdAt,omitempty"`
	UpdatedAt string    `json:"updatedAt,omitempty"`
}

type CreateCommentRequest struct {
	PostId  int64  `json:"postId" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type CreateCommentResponse struct {
	Data *Comment `json:"data,omitempty"`
	CommonResponse
}

type CreateReplyRequest struct {
	PostId          int64  `json:"postId" binding:"required"`
	ParentCommentId int64  `json:"parentCommentId" binding:"required"`
	Content         string `json:"content" binding:"required"`
}

type CreateReplyResponse struct {
	Data *Comment `json:"data,omitempty"`
	CommonResponse
}

type EditCommentRequest struct {
	CommentId int64  `json:"commentId" binding:"required"`
	Content   string `json:"content" binding:"required"`
}

type EditCommentResponse struct {
	Data *Comment `json:"data,omitempty"`
	CommonResponse
}

type DeleteCommentRequest struct {
	CommentId int64 `uri:"commentId" binding:"required"`
}

type DeleteCommentResponse struct {
	Data *Comment `json:"data,omitempty"`
	CommonResponse
}

type GetAllCommentRequest struct {
	PostId int64 `uri:"postId" binding:"required"`
}

type GetAllCommentResponse struct {
	Data []Comment `json:"data,omitempty"`
	CommonResponse
}

type GetCommentRequest struct {
	CommentId int64 `uri:"commentId,omitempty"`
}

type GetCommentResponse struct {
	Data *Comment `json:"data,omitempty"`
	CommonResponse
}
