package infrastructure

import (
	"errors"

	"github.com/yuuki0310/reservation_api/domain/model"
	"github.com/yuuki0310/reservation_api/domain/repository"
	"github.com/yuuki0310/reservation_api/infrastructure/mysql"
	"gorm.io/gorm"
)

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository() repository.UserRepository {
    return &userRepository{db: mysql.DB}
}

func (r userRepository) GetUserIDByUUID(uuid string) (uint, error) {
    var user model.User
    if err := r.db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return 0, errors.New("ユーザーが見つかりません")
        }
        return 0, err
    }
    return user.ID, nil
}
