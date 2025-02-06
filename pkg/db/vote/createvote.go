package vote

import (
	"context"
	"database/sql"
	"fmt"
	"liquide-assignment/pkg/dto"
	"log"

	"gorm.io/gorm"
)

func (obj *voteDb) AddVote(ctx context.Context, vote dto.DbVote) (int64, error) {
	query := `
		insert into
			vote(user_id,post_id,vote_type)
		values
			(?,?,?)
		returning vote_id;
	`

	var voteId sql.NullInt64
	insertTx := obj.dbObj.WithContext(ctx).Raw(query, vote.UserId.Int64, vote.PostId.Int64, vote.VoteType.Bool).Scan(&voteId)
	if insertTx.Error != nil {
		log.Println("error in adding vote")
		return 0, insertTx.Error
	}

	return voteId.Int64, nil
}

func (obj *voteDb) UpdateVote(ctx context.Context, vote dto.DbVote) (int64, error) {
	query := `
		update vote
		set vote_type = ?
		where post_id = ? and user_id = ?
		returning vote_id;
	`

	var voteId sql.NullInt64
	updateTx := obj.dbObj.WithContext(ctx).Raw(query, vote.VoteType.Bool, vote.PostId.Int64, vote.UserId.Int64).Scan(&voteId)
	if updateTx.Error != nil {
		log.Println("error in updating post")
		if updateTx.Error == gorm.ErrRecordNotFound {
			return 0, fmt.Errorf("no vote found to update")
		}
		return 0, updateTx.Error
	}

	return voteId.Int64, nil
}

func (obj *voteDb) UpsertVote(ctx context.Context, vote dto.DbVote) (int64, error) {
	query := `
		insert into vote (user_id, post_id, vote_type)
		values (?, ?, ?)
		on conflict (user_id, post_id)
		do update set vote_type = excluded.vote_type
		returning vote_id;
	`

	var voteId sql.NullInt64
	upsertTx := obj.dbObj.WithContext(ctx).Raw(query, vote.UserId.Int64, vote.PostId.Int64, vote.VoteType.Bool).Scan(&voteId)
	if upsertTx.Error != nil {
		log.Println("error in upserting vote")
		return 0, upsertTx.Error
	}

	return voteId.Int64, nil
}
