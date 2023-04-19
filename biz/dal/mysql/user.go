package mysql

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	models "github.com/lyleshaw/mlops/biz/model/orm_gen"
	"github.com/lyleshaw/mlops/biz/model/query"
	"github.com/lyleshaw/mlops/pkg/utils"
	"time"
)

// CreateUser creates a new user
func CreateUser(ctx context.Context, name, email, password string) (*models.User, error) {
	hashedPassword, err := utils.HashAndSaltPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		UserName: name,
		Email:    email,
		Password: hashedPassword,
		IsActive: true,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	logger.Infof("Creating user: %+v", user)

	err = query.User.WithContext(ctx).Create(user)
	return user, err
}

// GetUserByID retrieves a user by their ID
func GetUserByID(ctx context.Context, userID int) (*models.User, error) {
	logger.Infof("Getting user by ID: %d", userID)
	user, err := query.User.WithContext(ctx).Where(query.User.UserID.Eq(int32(userID))).Where(query.User.IsDeleted.Is(false)).First()
	return user, err
}

// GetUserByEmail retrieves a user by their Email
func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	logger.Infof("Getting user by email: %s", email)
	user, err := query.User.WithContext(ctx).Where(query.User.Email.Eq(email)).Where(query.User.IsDeleted.Is(false)).First()
	return user, err
}

// GetAllUsers retrieves all users
func GetAllUsers(ctx context.Context) ([]*models.User, error) {
	logger.Info("Getting all users")
	users, err := query.User.WithContext(ctx).Where(query.User.IsDeleted.Is(false)).Find()
	logger.Infof("Retrieved all users: %+v", users)
	return users, err
}

// UpdateUserByID updates a user's information by their ID
func UpdateUserByID(ctx context.Context, userID int, name, avatar, email, password string) (*models.User, error) {
	logger.Infof("Updating user by ID: %d", userID)
	user, err := GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	hashedPassword, err := utils.HashAndSaltPassword(password)
	if err != nil {
		return nil, err
	}

	user.UserName = name
	user.Avatar = avatar
	user.Email = email
	user.Password = hashedPassword
	user.UpdateAt = time.Now()

	err = query.User.WithContext(ctx).Save(user)
	logger.Infof("Updated user: %+v", user)
	return user, err
}

// DeleteUserByID soft deletes a user by their ID
func DeleteUserByID(ctx context.Context, userID int) error {
	user, err := GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	if user == nil {
		return nil
	}
	user.IsDeleted = true
	err = query.User.WithContext(ctx).Save(user)
	logger.Infof("Soft deleted user: %d", userID)
	return err
}
