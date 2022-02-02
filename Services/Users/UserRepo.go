package users

import (
	Model "crowdfunding/Model"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user Model.User) (Model.User, error)
	FindByEmail(email string) (Model.User, error)
	FindByID(id int) (Model.User, error)
	Update(user Model.User) (Model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user Model.User) (Model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (Model.User, error) {
	var user Model.User

	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(id int) (Model.User, error) {
	var user Model.User

	err := r.db.First(&user, id).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// for update all column
func (r *repository) Update(user Model.User) (Model.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
