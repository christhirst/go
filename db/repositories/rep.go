package repositories

import "database/sql"

// UserRepo implements models.UserRepository
type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// FindByID ..
//func (r *UserRepo) FindByID(ID int) (*models.User, error) {
//	return &models.User{}, nil
//}

// Save ..
//func (r *UserRepo) Save(user *models.User) error {
//	return nil
//}
