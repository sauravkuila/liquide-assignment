package onboarding

import (
	"log"
	"net/http"
	"strings"
	"time"

	"liquide-assignment/pkg/auth"
	"liquide-assignment/pkg/dto"
	e "liquide-assignment/pkg/errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (obj *onboardingService) UserSignup(c *gin.Context) {
	var (
		request  dto.UserSignupRequest
		response dto.UserSignupResponse
	)
	if err := c.BindJSON(&request); err != nil {
		log.Printf("unable to marshal request. Error:%s", err.Error())
		response.Errors = append(response.Errors, *e.ErrorInfo[e.BadRequest])
		response.Message = "failed to signup user"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//hash the password
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("unable to hash password. Error:%s", err.Error())
		response.Errors = append(response.Errors, e.ErrorInfo[e.ConversionError].GetErrorDetails("failed to hash the password"))
		response.Message = "failed to signup user"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	hashedPassword := string(hashedPasswordBytes)

	//add the entry into db
	userDetail := request.ToUserDetails()
	userDetail.UserPassword = hashedPassword

	log.Println(userDetail)
	userId, err := obj.dbObj.AddUser(c, userDetail.ToDbUserDetail())
	if err != nil {
		log.Printf("failed to add user. Error: %s", err.Error())
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			response.Errors = append(response.Errors, e.ErrorInfo[e.AddDBError].GetErrorDetails("unique username needed"))
		} else {
			response.Errors = append(response.Errors, e.ErrorInfo[e.AddDBError].GetErrorDetails(err.Error()))
		}
		response.Message = "failed to signup user"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = true
	response.Data = &dto.UserSignup{
		UserName: request.UserName,
		UserId:   userId,
	}
	response.Message = "successfully signed up user"
	c.JSON(http.StatusOK, response)
}

func (obj *onboardingService) UserLogin(c *gin.Context) {
	var (
		request  dto.UserLoginRequest
		response dto.UserLoginResponse
	)
	if err := c.BindJSON(&request); err != nil {
		log.Printf("unable to marshal request. Error:%s", err.Error())
		response.Errors = append(response.Errors, *e.ErrorInfo[e.BadRequest])
		response.Message = "failed to login user"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//fetch password hash for the user
	userDbDetail, err := obj.dbObj.GetUserByUsername(c, request.UserName)
	if err != nil {
		log.Printf("failed to fetch user data. Error: %s", err.Error())
		response.Errors = append(response.Errors, e.ErrorInfo[e.GetDBError].GetErrorDetails(err.Error()))
		response.Message = "failed to login user"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	userDetail := userDbDetail.ToUserDetail()

	if userDetail.UserId == 0 {
		//invalid username sent
		log.Println("username not found")
		response.Errors = append(response.Errors, e.ErrorInfo[e.NoDataFound].GetErrorDetails("username not found"))
		response.Message = "failed to login user"
		c.JSON(http.StatusNotFound, response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDetail.UserPassword), []byte(request.Password))
	if err != nil || userDetail.UserName != request.UserName {
		log.Println("invalid password")
		response.Errors = append(response.Errors, e.ErrorInfo[e.BadRequest].GetErrorDetails("incorrect username/password"))
		response.Message = "failed to login user"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	exp := time.Now().Add(60 * time.Minute)
	payload := auth.Token{
		UserName: userDetail.UserName,
		UserId:   userDetail.UserId,
		UserType: userDetail.UserType,
		Exp:      exp,
	}

	token, err := auth.GenerateJWT(payload)
	if err != nil {
		log.Println("failed to generate JWT")
		response.Errors = append(response.Errors, e.ErrorInfo[e.DefaultError].GetErrorDetails("failed to generate JWT token"))
		response.Message = "failed to login user"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = true
	response.Data = &dto.UserLogin{
		Token:        token,
		RefreshToken: "",
		Expiry:       exp.Format("2006-01-02 15:04:05"),
	}
	response.Message = "successfully logged in user"
	c.JSON(http.StatusOK, response)
}
