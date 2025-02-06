package post

import (
	"context"
	"database/sql"
	"fmt"
	"liquide-assignment/pkg/dto"
	"log"

	"gorm.io/gorm"
)

func (obj *postDb) CreatePost(ctx context.Context, post dto.DbPost) (int64, error) {
	query := `
		insert into
			post(user_id,content)
		values
			(?,?)
		returning post_id;
	`

	var postId sql.NullInt64
	insertTx := obj.dbObj.WithContext(ctx).Raw(query, post.UserId.Int64, post.Content.String).Scan(&postId)
	if insertTx.Error != nil {
		log.Println("error in creating post")
		return 0, insertTx.Error
	}

	return postId.Int64, nil
}

func (obj *postDb) UpdatePost(ctx context.Context, post dto.DbPost) (int64, error) {
	query := `
		update post
		set content = ?
		where post_id = ? and user_id = ? and is_deleted = FALSE
		returning post_id;
	`

	var postId sql.NullInt64
	updateTx := obj.dbObj.WithContext(ctx).Raw(query, post.Content.String, post.PostId.Int64, post.UserId.Int64).Scan(&postId)
	if updateTx.Error != nil {
		log.Println("error in updating post")
		if updateTx.Error == gorm.ErrRecordNotFound {
			return 0, fmt.Errorf("no post found to udpate")
		}
		return 0, updateTx.Error
	}
	if updateTx.RowsAffected == int64(0) {
		return 0, gorm.ErrRecordNotFound
	}
	updateTx.Commit()

	return postId.Int64, nil
}

func (obj *postDb) DeletePost(ctx context.Context, postId int64, userId int64) error {
	query := `
		update post
		set is_deleted = TRUE
		where post_id = ? and user_id = ?;
	`

	deleteTx := obj.dbObj.WithContext(ctx).Exec(query, postId, userId)
	if deleteTx.Error != nil {
		log.Println("error in deleting post")
		if deleteTx.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("no post found to delete")
		}
		return deleteTx.Error
	}
	if deleteTx.RowsAffected == int64(0) {
		return gorm.ErrRecordNotFound
	}
	deleteTx.Commit()

	return nil
}
