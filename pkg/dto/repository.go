package dto

import (
	"database/sql"
)

type DbUserDetail struct {
	UserId       sql.NullInt64
	UserName     sql.NullString
	UserPassword sql.NullString
	UserType     sql.NullString
	Email        sql.NullString
	Mobile       sql.NullString
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
}

func (u *DbUserDetail) ToUserDetail() UserDetail {
	return UserDetail{
		UserId:       u.UserId.Int64,
		UserName:     u.UserName.String,
		UserPassword: u.UserPassword.String,
		UserType:     u.UserType.String,
		Email:        u.Email.String,
		Mobile:       u.Mobile.String,
		CreatedAt:    u.CreatedAt.Time,
		UpdatedAt:    u.UpdatedAt.Time,
	}
}

type DbPost struct {
	PostId    sql.NullInt64
	UserId    sql.NullInt64
	Content   sql.NullString
	IsDeleted sql.NullBool
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type DbPostInfo struct {
	PostId       sql.NullInt64
	UserId       sql.NullInt64
	Content      sql.NullString
	UserName     sql.NullString
	UpVote       sql.NullInt64
	DownVote     sql.NullInt64
	CommentCount sql.NullInt64
	IsDeleted    sql.NullBool
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
}

func (p *DbPostInfo) ToPostInfo() PostInfo {
	return PostInfo{
		PostId:        p.PostId.Int64,
		Content:       p.Content.String,
		UserName:      p.UserName.String,
		UpVoteCount:   p.UpVote.Int64,
		DownVoteCount: p.DownVote.Int64,
		CreatedAt:     p.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:     p.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
	}
}

type DbComment struct {
	CommentId       sql.NullInt64
	PostId          sql.NullInt64
	UserId          sql.NullInt64
	ParentCommentId sql.NullInt64
	Content         sql.NullString
	IsDeleted       sql.NullBool
	CreatedAt       sql.NullTime
	UpdatedAt       sql.NullTime
}

type DbVote struct {
	VoteId    sql.NullInt64
	PostId    sql.NullInt64
	UserId    sql.NullInt64
	VoteType  sql.NullBool
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
