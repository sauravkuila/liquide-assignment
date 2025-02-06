package dto

type Vote struct {
	VoteId    int64  `json:"voteId,omitempty"`
	PostId    int64  `json:"postId,omitempty"`
	UserId    int64  `json:"-"`
	UserName  string `json:"userName,omitempty"`
	VoteType  string `json:"voteType,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type UpVoteRequest struct {
	PostId int64 `json:"postId,omitempty"`
}

type UpVoteResponse struct {
	Data *Vote `json:"data,omitempty"`
	CommonResponse
}

type DownVoteRequest struct {
	PostId int64 `json:"postId,omitempty"`
}

type DownVoteResponse struct {
	Data *Vote `json:"data,omitempty"`
	CommonResponse
}
