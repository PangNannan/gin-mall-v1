package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{
		NewDBClient(ctx),
	}
}
func NewUserDaoByDb(db *gorm.DB) *UserDao {
	return &UserDao{
		db,
	}
}

// 查看是否存在该用户名
func (userDao *UserDao) ExitOrNotByUserName(userName string) (user *model.User, exit bool, err error) {
	var count int64
	err = userDao.DB.Model(&model.User{}).Where("user_name = ?", userName).First(&user).Count(&count).Error

	if count == 1 {
		return user, true, err
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, true, err
	}
	return user, false, nil
}

// 创建一个用户
func (userDao *UserDao) CreateUser(user *model.User) (err error) {
	return userDao.Model(&model.User{}).Create(user).Error
}

// 根据id 获取user
func (userDao *UserDao) GetUserByUId(id uint) (user *model.User, err error) {

	err = userDao.Model(&model.User{}).Where("id = ?", id).First(&user).Error
	return
}

// 通过id更新用户属性
// 包括零值的字段
func (userDao *UserDao) UpdateUserById(uId uint, user *model.User) error {
	return userDao.Model(&model.User{}).Select("*").Where("id = ?", uId).Updates(user).Error
}
