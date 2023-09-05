package user

import (
	"duongGoGo/models"
	"duongGoGo/modules/user/userdto"
	"duongGoGo/utils/errorutil"
	"duongGoGo/utils/typeutil"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Service struct {
}

func (service Service) GetUser(c *gin.Context) {
	var query userdto.GetUserRequestDto

	if err := c.ShouldBindUri(&query); err != nil {
		message := errorutil.ParseError(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	user, err := userRepository.GetUser(query)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "not found user with this id"})
		return
	}

	userResponse, err := typeutil.TypeConverter[userdto.UserResponse, models.User](user)
	c.JSON(http.StatusOK, gin.H{"user": userResponse})
}

func (service Service) GetAll(c *gin.Context) {
	users, err := userRepository.GetAll()

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	usersResponse, err := typeutil.ArrayTypeConverter[userdto.UserResponse, models.User](users)

	c.JSON(http.StatusOK, gin.H{"users": usersResponse})
}

func (service Service) CreateUser(c *gin.Context) {
	var createForm userdto.CreateUserRequestDto

	if err := c.BindJSON(&createForm); err != nil {
		message := errorutil.ParseError(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	newUser := models.User{Email: createForm.Email, Password: createForm.Password}
	err := userRepository.Save(newUser)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": newUser})
}

func (service Service) UpdateUser(c *gin.Context) {
	var updateForm userdto.UpdateUserRequestDto
	var query userdto.GetUserRequestDto

	if err := c.BindUri(&query); err != nil {
		message := errorutil.ParseError(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	if err := c.BindJSON(&updateForm); err != nil {
		message := errorutil.ParseError(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	user, err := userRepository.GetUser(query)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "can't find user with this id"})
		return
	}

	user.BirthDay, err = time.Parse("2006-01-02", updateForm.BirthDay)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid birthday"})
		return
	}
	err = userRepository.Update(user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
