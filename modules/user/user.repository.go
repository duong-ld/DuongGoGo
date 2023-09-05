package user

import (
	"duongGoGo/infra/database"
	"duongGoGo/models"
	"duongGoGo/modules/user/userdto"
	"time"
)

type Repository struct {
}

func (repository Repository) GetUser(query userdto.GetUserRequestDto) (user models.User, err error) {
	db := database.GetDB()
	id := query.ID

	result := db.Where("id = ?", id).First(&user)
	return user, result.Error
}

func (repository Repository) GetUserBirthdayToday() (users []models.User, err error) {
	db := database.GetDB()
	currentMonth := time.Now().Month()
	currentDay := time.Now().Day()

	result := db.Where("MONTH(birth_day) = ?", currentMonth).Where("DAY(birth_day) = ?", currentDay).Find(&users)

	return users, result.Error
}

func (repository Repository) GetLuckyUser() (user models.User, err error) {
	db := database.GetDB()
	result := db.Order("RAND()").First(&user)

	return user, result.Error

}

func (repository Repository) GetUserByEmail(email string) (user models.User, err error) {
	db := database.GetDB()

	result := db.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (repository Repository) GetAll() (users []models.User, err error) {
	db := database.GetDB()

	result := db.Find(&users)
	return users, result.Error
}

func (repository Repository) Save(newUser models.User) error {
	db := database.GetDB()

	result := db.Create(&newUser)

	return result.Error
}

func (repository Repository) Update(user models.User) error {
	db := database.GetDB()

	result := db.Save(&user)

	return result.Error
}
