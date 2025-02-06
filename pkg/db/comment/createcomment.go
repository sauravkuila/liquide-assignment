package comment

import (
	"context"
	"database/sql"
	"fmt"
	"liquide-assignment/pkg/dto"
	"log"

	"gorm.io/gorm"
)

func (obj *commentDb) AddComment(ctx context.Context, comment dto.DbComment) (int64, error) {
	query := `
		insert into
			comment(post_id,user_id,content)
		values
			(?,?,?)
		returning comment_id;
	`

	var commentId sql.NullInt64
	insertTx := obj.dbObj.WithContext(ctx).Raw(query, comment.PostId.Int64, comment.UserId.Int64, comment.Content.String).Scan(&commentId)
	if insertTx.Error != nil {
		log.Println("error in adding comment")
		return 0, insertTx.Error
	}
	insertTx.Commit()

	return commentId.Int64, nil
}

func (obj *commentDb) AddReply(ctx context.Context, comment dto.DbComment) (int64, error) {
	query := `
		insert into
			comment(post_id,user_id,content,parent_comment_id)
		values
			(?,?,?,?)
		returning comment_id;
	`

	var commentId sql.NullInt64
	insertTx := obj.dbObj.WithContext(ctx).Raw(query, comment.PostId.Int64, comment.UserId.Int64, comment.Content.String, comment.ParentCommentId.Int64).Scan(&commentId)
	if insertTx.Error != nil {
		log.Println("error in adding comment")
		return 0, insertTx.Error
	}
	insertTx.Commit()

	return commentId.Int64, nil
}

func (obj *commentDb) UpdateComment(ctx context.Context, comment dto.DbComment) (int64, error) {
	query := `
		update comment
		set content = ?
		where comment_id = ? and user_id = ? and is_deleted = FALSE
		returning comment_id;
	`

	var commentId sql.NullInt64
	updateTx := obj.dbObj.WithContext(ctx).Raw(query, comment.Content.String, comment.CommentId.Int64, comment.UserId.Int64).Scan(&commentId)
	if updateTx.Error != nil {
		log.Println("error in updating comment")
		if updateTx.Error == gorm.ErrRecordNotFound {
			return 0, fmt.Errorf("no comment found to update")
		}
		return 0, updateTx.Error
	}
	if updateTx.RowsAffected == int64(0) {
		return 0, gorm.ErrRecordNotFound
	}
	updateTx.Commit()

	return commentId.Int64, nil
}

func (obj *commentDb) DeleteComment(ctx context.Context, commentId int64, userId int64) error {
	query := `
		update comment
		set is_deleted = TRUE
		where comment_id = ? and user_id = ?;
	`

	deleteTx := obj.dbObj.WithContext(ctx).Exec(query, commentId, userId)
	if deleteTx.Error != nil {
		log.Println("error in deleting comment")
		if deleteTx.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("no comment found to delete")
		}
		return deleteTx.Error
	}
	if deleteTx.RowsAffected == int64(0) {
		return gorm.ErrRecordNotFound
	}
	deleteTx.Commit()

	return nil
}
