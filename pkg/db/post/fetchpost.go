package post

import (
	"context"
	"database/sql"
	"liquide-assignment/pkg/dto"
)

func (obj *postDb) GetPost(ctx context.Context, postId int64) (dto.DbPostInfo, error) {

	// query := `
	// 	SELECT
	// 		p.post_id,
	// 		p.user_id,
	// 		u.user_name,
	// 		COALESCE(SUM(CASE WHEN v.vote_type THEN 1 ELSE 0 END), 0) AS upvote,
	// 		COALESCE(SUM(CASE WHEN NOT v.vote_type THEN 1 ELSE 0 END), 0) AS downvote,
	// 		p.created_at,
	// 		p.updated_at
	// 	FROM
	// 		post p
	// 	JOIN
	// 		user_detail u ON p.user_id = u.id
	// 	LEFT JOIN
	// 		vote v ON p.post_id = v.post_id
	// 	WHERE
	// 		p.post_id = $1
	// 	GROUP BY
	// 		p.post_id, u.user_name
	// `
	query := `WITH votes_cte AS (
		SELECT 
			post_id,
			SUM(CASE WHEN vote_type = TRUE THEN 1 ELSE 0 END) AS upvotes,
			SUM(CASE WHEN vote_type = FALSE THEN 1 ELSE 0 END) AS downvotes
		FROM vote
		GROUP BY post_id
		),
		comments_cte AS (
			SELECT 
				post_id,
				COUNT(1) AS comment_count
			FROM comment
			GROUP BY post_id
		)
		SELECT
			p.post_id,
			p.user_id,
			u.user_name,
			COALESCE(v.upvotes, 0) AS upvotes,
			COALESCE(v.downvotes, 0) AS downvotes,
			COALESCE(c.comment_count, 0) AS comment_count,
			p.created_at,
			p.updated_at
		FROM post p
		JOIN user_detail u ON p.user_id = u.id
		LEFT JOIN votes_cte v ON p.post_id = v.post_id
		LEFT JOIN comments_cte c ON p.post_id = c.post_id
		WHERE p.post_id = ?;		
	`

	row := obj.dbObj.WithContext(ctx).Raw(query, postId).Row()
	if row.Err() != nil {
		return dto.DbPostInfo{}, row.Err()
	}

	var postInfo dto.DbPostInfo
	err := row.Scan(
		&postInfo.PostId,
		&postInfo.UserId,
		&postInfo.UserName,
		&postInfo.UpVote,
		&postInfo.DownVote,
		&postInfo.CommentCount,
		&postInfo.CreatedAt,
		&postInfo.UpdatedAt,
	)

	if err != nil {
		return dto.DbPostInfo{}, err
	}

	return postInfo, nil
}

func (obj *postDb) GetUserPosts(ctx context.Context, userId int64, limit, offset int) ([]dto.PostInfo, int64, error) {
	query := `
		WITH votes_cte AS (
			SELECT 
				post_id,
				SUM(CASE WHEN vote_type = TRUE THEN 1 ELSE 0 END) AS upvotes,
				SUM(CASE WHEN vote_type = FALSE THEN 1 ELSE 0 END) AS downvotes
			FROM vote
			GROUP BY post_id
		),
		comments_cte AS (
			SELECT 
				post_id,
				COUNT(1) AS comment_count
			FROM comment
			GROUP BY post_id
		),
		total_count AS (
			SELECT COUNT(1) AS total
			FROM post
			WHERE user_id = ?
		)
		SELECT
			p.post_id,
			p.user_id,
			u.user_name,
			COALESCE(v.upvotes, 0) AS upvotes,
			COALESCE(v.downvotes, 0) AS downvotes,
			COALESCE(c.comment_count, 0) AS comment_count,
			p.created_at,
			p.updated_at,
			tc.total
		FROM post p
		JOIN user_detail u ON p.user_id = u.id
		LEFT OUTER JOIN votes_cte v ON p.post_id = v.post_id
		LEFT OUTER JOIN comments_cte c ON p.post_id = c.post_id
		CROSS JOIN total_count tc
		WHERE p.user_id = ?
		ORDER BY p.created_at DESC
		LIMIT ? OFFSET ?;		
	`

	rows, err := obj.dbObj.WithContext(ctx).Raw(query, userId, userId, limit, offset).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var (
		posts        []dto.PostInfo
		totalRecords sql.NullInt64
	)
	for rows.Next() {
		var (
			postInfo dto.DbPostInfo
		)
		err := rows.Scan(
			&postInfo.PostId,
			&postInfo.UserId,
			&postInfo.UserName,
			&postInfo.UpVote,
			&postInfo.DownVote,
			&postInfo.CommentCount,
			&postInfo.CreatedAt,
			&postInfo.UpdatedAt,
			&totalRecords,
		)
		if err != nil {
			return nil, 0, err
		}
		posts = append(posts, dto.PostInfo{
			PostId:        postInfo.PostId.Int64,
			UserName:      postInfo.UserName.String,
			Content:       postInfo.Content.String,
			UpVoteCount:   postInfo.UpVote.Int64,
			DownVoteCount: postInfo.DownVote.Int64,
			CommentCount:  postInfo.CommentCount.Int64,
			CreatedAt:     postInfo.CreatedAt.Time.Format("2006-01-02 15:04:05"),
			UpdatedAt:     postInfo.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		})
	}

	return posts, totalRecords.Int64, nil
}
