package auth

import (
	"context"
	"duongGoGo/infra/caching"
	"duongGoGo/models"
	"duongGoGo/modules/auth/authdto"
	"duongGoGo/modules/user/userdto"
	"duongGoGo/utils/encryptutil"
	"duongGoGo/utils/errorutil"
	"duongGoGo/utils/tokenutil"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Service struct {
}

func (s Service) SignIn(c *gin.Context) {
	ctx := context.Background()
	rds := caching.GetRedisClient()

	signInForm := authdto.SignInDto{}

	if err := c.ShouldBindJSON(&signInForm); err != nil {
		message := errorutil.ParseError(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}
	user, err := userRepository.GetUserByEmail(signInForm.Email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "email or password not valid"})
		return
	}
	isValidPassword := encryptutil.ComparePassword(user.Password, signInForm.Password)

	if isValidPassword == false {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "password not correct"})
		return
	}

	accessToken, err := tokenutil.CreateToken(user.ID, tokenutil.ACCESS)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "can't create token"})
	}
	c.SetCookie("ACCESS_TOKEN", accessToken, 3600, "", "", false, true)
	accessTokenRedisKey := tokenutil.CreateRedisKeyForToken(user.ID.String(), accessToken)
	rds.Set(ctx, accessTokenRedisKey, nil, time.Hour)

	refreshToken, err := tokenutil.CreateToken(user.ID, tokenutil.REFRESH)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "can't create token"})
	}
	c.SetCookie("REFRESH_TOKEN", refreshToken, 3600*24*30, "", "", false, true)
	refreshTokenRedisKey := tokenutil.CreateRedisKeyForToken(user.ID.String(), refreshToken)
	rds.Set(ctx, refreshTokenRedisKey, nil, time.Hour*24*30)

	csrfToken, err := tokenutil.CreateToken(user.ID, tokenutil.CSRF)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "can't create token"})
	}
	c.SetCookie("CSRF_TOKEN", csrfToken, 0, "", "", false, false)
	csrfTokenRedisKey := tokenutil.CreateRedisKeyForToken(user.ID.String(), csrfToken)
	rds.Set(ctx, csrfTokenRedisKey, nil, time.Hour*24*30)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "can't create session"})
	}
	c.JSON(http.StatusOK, gin.H{
		"csrf_token": csrfToken,
		"message":    "login success",
	})
}

func (s Service) SignUp(c *gin.Context) {
	signUpForm := new(authdto.SignUpDto)

	if err := c.ShouldBindJSON(&signUpForm); err != nil {
		message := errorutil.ParseError(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	hashedPassword, err := encryptutil.HashPassword(signUpForm.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusExpectationFailed, gin.H{"message": "failed when encrypt password"})
		return
	}

	createUserForm := userdto.CreateUserRequestDto{
		Email:    signUpForm.Email,
		Password: hashedPassword,
	}
	newUser := models.User{Email: createUserForm.Email, Password: createUserForm.Password}
	err = userRepository.Save(newUser)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": newUser})
}

func (s Service) Logout(c *gin.Context) {
	ctx := context.Background()
	rds := caching.GetRedisClient()

	userId, isExist := c.Get("userId")

	if !isExist {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
	}

	accessTokenString, err := c.Cookie("ACCESS_TOKEN")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
	}
	accessTokenRedisKey := tokenutil.CreateRedisKeyForToken(userId.(string), accessTokenString)

	rds.Del(ctx, accessTokenRedisKey)
}
