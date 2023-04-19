package mysql

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	models "github.com/lyleshaw/mlops/biz/model/orm_gen"
	"github.com/lyleshaw/mlops/biz/model/query"
	"gorm.io/gorm"
)

// CreateApplication creates a new application
func CreateApplication(ctx context.Context, name string, port int, image string, dockerID string, userID int) (*models.Application, error) {
	logger.Infof("Creating application: %s, port: %d, image: %s, userID: %d", name, port, image, userID)
	app := &models.Application{
		AppName:   name,
		AppPort:   int32(port),
		AppImage:  image,
		DockerID:  dockerID,
		UserID:    int32(userID),
		AppStatus: "development",
		IsDeleted: false,
	}
	err := query.Application.WithContext(ctx).Create(app)
	if err != nil {
		logger.Errorf("Failed to create application: %v", err)
	}
	return app, err
}

// GetApplicationsByUserID retrieves applications by user ID
func GetApplicationsByUserID(ctx context.Context, userID int) ([]*models.Application, error) {
	logger.Infof("Getting applications for userID: %d", userID)
	apps, err := query.Application.WithContext(ctx).Where(query.Application.UserID.Eq(int32(userID))).Where(query.Application.IsDeleted.Is(false)).Find()
	if err != nil {
		logger.Errorf("Failed to get applications for userID %d: %v", userID, err)
	}
	return apps, err
}

// GetApplicationByID retrieves an application by ID
func GetApplicationByID(ctx context.Context, appID int) (*models.Application, error) {
	logger.Infof("Getting application with ID: %d", appID)
	app, err := query.Application.WithContext(ctx).Where(query.Application.AppID.Eq(int32(appID))).Where(query.Application.IsDeleted.Is(false)).First()
	if err != nil {
		logger.Errorf("Failed to get application with ID %d: %v", appID, err)
	}
	return app, err
}

// UpdateApplicationByID updates an application by ID
func UpdateApplicationByID(ctx context.Context, appID int, name *string, image *string) (*models.Application, error) {
	logger.Infof("Updating application with ID: %d, name: %s, image: %s", appID, name, image)
	app, err := GetApplicationByID(ctx, appID)
	if err != nil {
		return nil, err
	}
	if name != nil {
		app.AppName = *name
	}
	if image != nil {
		app.AppImage = *image
	}
	err = query.Application.WithContext(ctx).Save(app)
	if err != nil {
		logger.Errorf("Failed to update application with ID %d: %v", appID, err)
	}
	app, err = GetApplicationByID(ctx, appID)
	if err != nil {
		logger.Errorf("Failed to get application with ID %d: %v in UpdateApplicationByID", appID, err)
		return nil, err
	}
	return app, err
}

// UpdateApplicationStatusByID updates an application by ID
func UpdateApplicationStatusByID(ctx context.Context, appID int, status string) (*models.Application, error) {
	logger.Infof("Updating application with ID: %d, status: %s", appID, status)
	app, err := GetApplicationByID(ctx, appID)
	if err != nil {
		return nil, err
	}
	app.AppStatus = status
	err = query.Application.WithContext(ctx).Save(app)
	if err != nil {
		logger.Errorf("Failed to update application with ID %d: %v", appID, err)
	}
	return app, err
}

// DeleteApplicationByID soft deletes an application by ID
func DeleteApplicationByID(ctx context.Context, appID int) error {
	logger.Infof("Soft deleting application with ID: %d", appID)
	app, err := GetApplicationByID(ctx, appID)
	if err != nil {
		return err
	}
	app.IsDeleted = true
	err = query.Application.WithContext(ctx).Save(app)
	if err != nil {
		logger.Errorf("Failed to soft delete application with ID %d: %v", appID, err)
	}
	return err
}

// GeneratePort 生成一个未使用的端口(最小值为 10000)
func GeneratePort() (port int, err error) {
	// 找一个未使用的端口
	db := DB
	var maxPort int
	if err = db.Table("applications").Select("COALESCE(MAX(app_port), 0)").Scan(&maxPort).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 表中没有数据，maxPort 的值为 0 或者其他默认值，视情况而定
			return 10000, nil
		} else {
			logger.Errorf("get max port error: %v", err)
			return -1, err
		}
	}
	if maxPort == 0 {
		return 10000, nil
	}
	return maxPort + 1, nil
}
