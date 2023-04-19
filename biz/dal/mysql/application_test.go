package mysql

import (
	"context"
	"database/sql"
	"github.com/bytedance/gopkg/util/logger"
	models "github.com/lyleshaw/mlops/biz/model/orm_gen"
	"github.com/lyleshaw/mlops/biz/model/query"
	"log"
	"testing"
)

func clearApplicationDatabase() {
	db, err := sql.Open("mysql", "test:test@tcp(localhost:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Errorf("close db failed: %v", err)
		}
	}(db)

	// 清空 users 表
	_, err = db.Exec("TRUNCATE applications;")
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateApplication(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()
	name := "App Name"
	port := 8080
	image := "app/image:v1"
	dockerID := "1"
	userID := 1

	app, err := CreateApplication(ctx, name, port, image, dockerID, userID)
	if err != nil {
		t.Fatalf("CreateApplication() failed: %v", err)
	}

	if app.AppName != name {
		t.Errorf("Expected app name %q, got %q", name, app.AppName)
	}
	if app.AppPort != int32(port) {
		t.Errorf("Expected port %d, got %d", port, app.AppPort)
	}
	if app.AppImage != image {
		t.Errorf("Expected image %q, got %q", image, app.AppImage)
	}
	if app.UserID != int32(userID) {
		t.Errorf("Expected userID %d, got %d", userID, app.UserID)
	}

	clearApplicationDatabase()
}

func TestGetApplicationsByUserID(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()
	userID := 1
	apps := []*models.Application{
		{AppName: "App 1", AppPort: 8080, AppImage: "app/image:v1", UserID: 1},
		{AppName: "App 2", AppPort: 8081, AppImage: "app/image:v2", UserID: 1},
	}

	for _, app := range apps {
		_, err := CreateApplication(ctx, app.AppName, int(app.AppPort), app.AppImage, app.DockerID, int(app.UserID))
		if err != nil {
			t.Fatalf("CreateApplication() failed: %v", err)
		}
	}

	retrievedApps, err := GetApplicationsByUserID(ctx, userID)
	if err != nil {
		t.Fatalf("GetApplicationsByUserID() failed: %v", err)
	}

	if len(retrievedApps) != len(apps) {
		t.Errorf("Expected %d apps, got %d", len(apps), len(retrievedApps))
	}

	for i, app := range apps {
		retrievedApp := retrievedApps[i]
		if app.AppName != retrievedApp.AppName {
			t.Errorf("Expected app name %q, got %q", app.AppName, retrievedApp.AppName)
		}
		if app.AppPort != retrievedApp.AppPort {
			t.Errorf("Expected port %d, got %d", app.AppPort, retrievedApp.AppPort)
		}
		if app.AppImage != retrievedApp.AppImage {
			t.Errorf("Expected image %q, got %q", app.AppImage, retrievedApp.AppImage)
		}
	}

	clearApplicationDatabase()
}

func TestGetApplicationByID(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()
	name := "App Name"
	port := 8080
	image := "app/image:v1"
	dockerID := "1"
	userID := 1

	app, err := CreateApplication(ctx, name, port, image, dockerID, userID)
	if err != nil {
		t.Fatalf("CreateApplication() failed: %v", err)
	}

	retrievedApp, err := GetApplicationByID(ctx, int(app.AppID))
	if err != nil {
		t.Fatalf("GetApplicationByID() failed: %v", err)
	}

	if retrievedApp.AppName != name {
		t.Errorf("Expected app name %q, got %q", name, retrievedApp.AppName)
	}
	if retrievedApp.AppPort != int32(port) {
		t.Errorf("Expected port %d, got %d", port, retrievedApp.AppPort)
	}
	if retrievedApp.AppImage != image {
		t.Errorf("Expected image %q, got %q", image, retrievedApp.AppImage)
	}
	if retrievedApp.UserID != int32(userID) {
		t.Errorf("Expected userID %d, got %d", userID, retrievedApp.UserID)
	}

	clearApplicationDatabase()
}

func TestUpdateApplicationByID(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()
	name := "App Name"
	port := 8080
	image := "app/image:v1"
	dockerID := "1"
	userID := 1

	app, err := CreateApplication(ctx, name, port, image, dockerID, userID)
	if err != nil {
		t.Fatalf("CreateApplication() failed: %v", err)
	}

	updatedName := "New Name"
	updatedImage := "app/image:v2"

	updatedApp, err := UpdateApplicationByID(ctx, int(app.AppID), &updatedName, &updatedImage)
	if err != nil {
		t.Fatalf("UpdateApplicationByID() failed: %v", err)
	}

	if updatedApp.AppName != updatedName {
		t.Errorf("Expected updated name %q, got %q", updatedName, updatedApp.AppName)
	}
	if updatedApp.AppImage != updatedImage {
		t.Errorf("Expected updated image %q, got %q", updatedImage, updatedApp.AppImage)
	}

	clearApplicationDatabase()
}

func TestUpdateApplicationStatusByID(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()
	name := "App Name"
	port := 8080
	image := "app/image:v1"
	dockerID := "1"
	userID := 1

	app, err := CreateApplication(ctx, name, port, image, dockerID, userID)
	if err != nil {
		t.Fatalf("CreateApplication() failed: %v", err)
	}

	updatedStatus := "closed"
	updatedApp, err := UpdateApplicationStatusByID(ctx, int(app.AppID), updatedStatus)
	if err != nil {
		t.Fatalf("UpdateApplicationStatusByID() failed: %v", err)
	}

	if updatedApp.AppStatus != updatedStatus {
		t.Errorf("Expected updated status %q, got %q", updatedStatus, updatedApp.AppStatus)
	}

	clearApplicationDatabase()
}

func TestDeleteApplicationByID(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()
	name := "App Name"
	port := 8080
	image := "app/image:v1"
	dockerID := "1"
	userID := 1

	app, err := CreateApplication(ctx, name, port, image, dockerID, userID)
	if err != nil {
		t.Fatalf("CreateApplication() failed: %v", err)
	}

	err = DeleteApplicationByID(ctx, int(app.AppID))
	if err != nil {
		t.Fatalf("DeleteApplicationByID() failed: %v", err)
	}

	deletedApp, err := query.Application.WithContext(ctx).Where(query.Application.AppID.Eq(app.AppID)).First()
	if err != nil {
		t.Fatalf("GetApplicationByID() failed: %v", err)
	}
	if deletedApp.IsDeleted != true {
		t.Error("Expected deleted application to have IsDeleted = true")
	}

	clearApplicationDatabase()
}

func TestGeneratePort(t *testing.T) {
	ctx := context.Background()
	name := "App Name"
	image := "app/image:v1"
	dockerID := "1"
	userID := 1

	// 第一次调用,返回 10000
	port, err := GeneratePort()
	if err != nil {
		t.Fatalf("GeneratePort() failed: %v", err)
	}
	if port != 10000 {
		t.Errorf("Expected port 10000, got %d", port)
	}

	// 创建一个应用,端口为 10000
	_, err = CreateApplication(ctx, name, port, image, dockerID, userID)
	if err != nil {
		t.Fatalf("CreateApplication() failed: %v", err)
	}

	// 第二次调用,返回 10001
	port, err = GeneratePort()
	if err != nil {
		t.Fatalf("GeneratePort() failed: %v", err)
	}
	if port != 10001 {
		t.Errorf("Expected port 10001, got %d", port)
	}

	clearApplicationDatabase()
}
