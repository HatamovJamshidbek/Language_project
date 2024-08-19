package postgres

import (
	pb "auth_service/genproto/user"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	db, err := ConnectDB()

	if err != nil {
		panic(err)
	}

	a := NewAuthRepo(db)

	req := pb.RegisterRequest{
		Username:       "aaa",
		Email:          "aaa",
		Password:       "aaa",
		Fullname:       "aaa",
		Nativelanguage: "aaa",
	}
	ctx := context.Background()

	resp, err := a.Register(ctx, &req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(resp)
}

func TestLogin(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	a := NewAuthRepo(db)

	req := pb.LoginRequest{
		Email:    "aaa",
		Password: "aaa",
	}
	ctx := context.Background()

	resp, err := a.Login(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func TestLogout(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	a := NewAuthRepo(db)

	req := pb.LogoutRequest{
		RefreshToken: "Votlp9rgIy1DOPsApVlQQvahiJGWIBaJtWxvpTEVZCgRz499MAjdIgjB0Y44yUgd",
	}
	ctx := context.Background()

	res, err := a.Logout(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestGetUserProfile(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	a := NewAuthRepo(db)

	ctx := context.Background()
	req := &pb.Void{}

	resp, err := a.GetUserProfile(ctx, req)
	if err != nil {
		panic(err)
	}

	assert.NotNil(t, resp)
	assert.Equal(t, "aaa", resp.Username)
	assert.Equal(t, "aaa", resp.Email)
	assert.Equal(t, "aaa", resp.Fullname)
	assert.Equal(t, "aaa", resp.Nativelanguage)
	assert.NotNil(t, resp.CreatedAt)
}

func TestGetByUserId(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	a := NewAuthRepo(db)

	ctx := context.Background()
	req := &pb.GetByUserIdRequest{Id: "eed9ef76-a2d3-40a3-a685-26cd0bb267bc"}

	resp, err := a.GetByUserId(ctx, req)
	if err != nil {
		panic(err)
	}

	assert.NotNil(t, resp)
	assert.Equal(t, "eed9ef76-a2d3-40a3-a685-26cd0bb267bc", resp.Id)
	assert.Equal(t, "aaa", resp.Username)
	assert.Equal(t, "aaa", resp.Email)
	assert.Equal(t, "$2a$10$w5T/nI/lgzwZjWY/U5wj/.RdIn7usNLrNvBmibZwekkpth.6dEBXm", resp.Password)
	assert.Equal(t, "aaa", resp.Fullname)
	assert.Equal(t, "aaa", resp.Nativelanguage)
	assert.NotNil(t, resp.CreatedAt)
	assert.NotNil(t, resp.UpdatedAt)
}

func TestDeleteUser(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	a := NewAuthRepo(db)

	req := pb.DeleteUserRequest{
		UserId: "eed9ef76-a2d3-40a3-a685-26cd0bb267bc",
	}
	ctx := context.Background()

	res, err := a.DeleteUser(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
