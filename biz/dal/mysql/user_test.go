package mysql

import (
	"context"
	"database/sql"
	"github.com/bytedance/gopkg/util/logger"
	model "github.com/lyleshaw/mlops/biz/model/orm_gen"
	"github.com/lyleshaw/mlops/biz/model/query"
	"log"
	"testing"
)

func clearUserDatabase() {
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
	_, err = db.Exec("TRUNCATE users;")
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateUser(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()
	name := "John Doe"
	email := "john@example.com"
	password := "password"

	user, err := CreateUser(ctx, name, email, password)
	if err != nil {
		t.Fatalf("CreateUser() failed: %v", err)
	}

	if user.UserName != name {
		t.Errorf("Expected user name %q, got %q", name, user.UserName)
	}
	if user.Email != email {
		t.Errorf("Expected email %q, got %q", email, user.Email)
	}
	if user.Password == password {
		t.Error("Password should be hashed")
	}

	clearUserDatabase()
}

func TestGetUserByID(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()
	name := "John Doe"
	email := "john@example.com"
	password := "password"

	user, err := CreateUser(ctx, name, email, password)
	if err != nil {
		t.Fatalf("CreateUser() failed: %v", err)
	}

	retrievedUser, err := GetUserByID(ctx, int(user.UserID))
	if err != nil {
		t.Fatalf("GetUserByID() failed: %v", err)
	}

	if retrievedUser.UserName != name {
		t.Errorf("Expected user name %q, got %q", name, retrievedUser.UserName)
	}
	if retrievedUser.Email != email {
		t.Errorf("Expected email %q, got %q", email, retrievedUser.Email)
	}

	clearUserDatabase()
}

func TestGetUserByEmail(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()
	name := "John Doe"
	email := "john@example.com"
	password := "password"

	_, err := CreateUser(ctx, name, email, password)
	if err != nil {
		t.Fatalf("CreateUser() failed: %v", err)
	}

	retrievedUser, err := GetUserByEmail(ctx, email)
	if err != nil {
		t.Fatalf("GetUserByEmail() failed: %v", err)
	}

	if retrievedUser.UserName != name {
		t.Errorf("Expected user name %q, got %q", name, retrievedUser.UserName)
	}
	if retrievedUser.Email != email {
		t.Errorf("Expected email %q, got %q", email, retrievedUser.Email)
	}

	clearUserDatabase()
}

func TestGetAllUsers(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()
	users := []model.User{
		{UserName: "John", Email: "john@example.com"},
		{UserName: "Jane", Email: "jane@example.com"},
	}

	for _, user := range users {
		_, err := CreateUser(ctx, user.UserName, user.Email, "password")
		if err != nil {
			t.Fatalf("CreateUser() failed: %v", err)
		}
	}

	retrievedUsers, err := GetAllUsers(ctx)
	if err != nil {
		t.Fatalf("GetAllUsers() failed: %v", err)
	}

	if len(retrievedUsers) != len(users) {
		t.Errorf("Expected %d users, got %d", len(users), len(retrievedUsers))
	}

	for i, user := range users {
		retrievedUser := retrievedUsers[i]
		if user.UserName != retrievedUser.UserName {
			t.Errorf("Expected user name %q, got %q", user.UserName, retrievedUser.UserName)
		}
		if user.Email != retrievedUser.Email {
			t.Errorf("Expected email %q, got %q", user.Email, retrievedUser.Email)
		}
	}

	clearUserDatabase()
}

func TestUpdateUserByID(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()
	name := "John Doe"
	email := "john@example.com"
	password := "password"

	user, err := CreateUser(ctx, name, email, password)
	if err != nil {
		t.Fatalf("CreateUser() failed: %v", err)
	}

	updatedName := "Jane Doe"
	updatedEmail := "jane@example.com"
	updatedPassword := "newpassword"

	updatedUser, err := UpdateUserByID(ctx, int(user.UserID), updatedName, "", updatedEmail, updatedPassword)
	if err != nil {
		t.Fatalf("UpdateUserByID() failed: %v", err)
	}

	if updatedUser.UserName != updatedName {
		t.Errorf("Expected updated name %q, got %q", updatedName, updatedUser.UserName)
	}
	if updatedUser.Email != updatedEmail {
		t.Errorf("Expected updated email %q, got %q", updatedEmail, updatedUser.Email)
	}
	if updatedUser.Password == updatedPassword {
		t.Error("Updated password should be hashed")
	}

	clearUserDatabase()
}

func TestDeleteUserByID(t *testing.T) {
	Init()
	query.SetDefault(DB)
	ctx := context.Background()
	name := "John Doe"
	email := "john@example.com"
	password := "password"

	user, err := CreateUser(ctx, name, email, password)
	if err != nil {
		t.Fatalf("CreateUser() failed: %v", err)
	}

	err = DeleteUserByID(ctx, int(user.UserID))
	if err != nil {
		t.Fatalf("DeleteUserByID() failed: %v", err)
	}

	deletedUser, err := query.User.WithContext(ctx).Where(query.User.UserID.Eq(user.UserID)).First()
	if err != nil {
		t.Fatalf("GetUserByID() failed: %v", err)
	}
	if deletedUser.IsDeleted != true {
		t.Error("Expected deleted user to have IsDeleted = true")
	}

	clearUserDatabase()
}
