package onboarding

import (
	"database/sql"
	"liquide-assignment/pkg/dto"
	"log"

	"github.com/gin-gonic/gin"
)

func (obj *onboardingDb) AddUser(c *gin.Context, userDetail dto.DbUserDetail) (int64, error) {
	query := `
		insert into
			user_detail(user_name,password,user_type,email,mobile)
		values
			(?,?,?,?,?)
		returning id;
	`

	var userId sql.NullInt64
	insertTx := obj.dbObj.WithContext(c).Raw(query, userDetail.UserName.String, userDetail.UserPassword.String, userDetail.UserType.String, userDetail.Email.String, userDetail.Mobile.String).Scan(&userId)
	if insertTx.Error != nil {
		log.Println("error in adding user")
		return 0, insertTx.Error
	}

	return userId.Int64, nil
}

func (obj *onboardingDb) GetUserByUsername(c *gin.Context, userName string) (dto.DbUserDetail, error) {
	query := `
		select 
			id, 
			user_name, 
			password, 
			user_type, 
			email, 
			mobile,
			created_at
		from
			user_detail
		where
			user_name=?;
	`

	var userDetail dto.DbUserDetail
	rows, err := obj.dbObj.WithContext(c).Raw(query, userName).Rows()
	if err != nil {
		log.Println("failed to fetch user detail")
		return userDetail, err
	}
	for rows.Next() {
		err := rows.Scan(&userDetail.UserId, &userDetail.UserName, &userDetail.UserPassword, &userDetail.UserType, &userDetail.Email, &userDetail.Mobile, &userDetail.CreatedAt)
		if err != nil {
			log.Println("failed to scan user detail")
			return userDetail, err
		}
	}
	return userDetail, nil
}
