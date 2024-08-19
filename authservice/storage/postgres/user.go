package postgres

import (
	pb "auth_service/genproto/user"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type UserManagementRepository interface {
	Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.Message, error)
	GetUserProfile(ctx context.Context, req *pb.Void) (*pb.ProfileResponse, error)
	UpdateUserProfile(ctx context.Context, req *pb.ProfileUpdateRequest) (*pb.ProfileUpdateResponse, error)
	ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.Message, error)
	ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.Message, error)
	ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.Message, error)
	GetByUserId(ctx context.Context, req *pb.GetByUserIdRequest) (*pb.GetByUserIdResponse, error)
	DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
}
type authRepo struct {
	Db *sql.DB
}

func NewAuthRepo(db *sql.DB) UserManagementRepository {
	return &authRepo{Db: db}
}

func (a *authRepo) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	passwordHash := HashPassword(req.Password)

	req.Password = passwordHash
	resp := pb.RegisterResponse{}
	q := `
		insert into users(username, email, password_hash, full_name, native_language, created_at, role)
		values($1, $2, $3, $4, $5, $6, $7) 
		returning id, username, email, password_hash, full_name, native_language, created_at
	`

	err := a.Db.QueryRow(q, req.Username, req.Email, passwordHash, req.Fullname, req.Nativelanguage, time.Now(), req.Role).Scan(
		&resp.Id,
		&resp.Username,
		&resp.Email,
		&resp.Password,
		&resp.Fullname,
		&resp.Nativelanguage,
		&resp.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (a *authRepo) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Println(req.Email, req.Password)
	resp := pb.LoginResponse{}
	passwordHash := ""

	pass := HashPassword(req.Password)
	fmt.Println(pass)
	q := `
	select 
		id,
		username,
		native_language,
		password_hash,
		role
	from 
		users
	where 
		email = $1 
	and 
		deleted_at is null
	`
	err := a.Db.QueryRow(q, req.Email).Scan(
		&resp.UserId,
		&resp.UserName,
		&resp.NativeLanguage,
		&passwordHash,
		&resp.Role,
	)

	if err != nil {
		return nil, err
	}

	if pass != passwordHash {
		return nil, fmt.Errorf("invalid password")
	}

	return &resp, nil
}

func (a *authRepo) Logout(ctx context.Context, token *pb.LogoutRequest) (*pb.Message, error) {
	_, err := a.Db.Exec(`
	update 
		reflesh_tokens 
	set 
		deleted_at=$1 
	where 
		token=$2`, time.Now(), token.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &pb.Message{Message: "Logout successfully"}, nil
}

func (a *authRepo) GetUserProfile(ctx context.Context, req *pb.Void) (*pb.ProfileResponse, error) {
	resp := &pb.ProfileResponse{}

	q := `
	select 
		id, 
		username, 
		email, 
		full_name, 
		native_language, 
		created_at 
	from 
		users
	where 
		deleted_at is null
	`

	rows, err := a.Db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&resp.Id,
			&resp.Username,
			&resp.Email,
			&resp.Fullname,
			&resp.Nativelanguage,
			&resp.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func (a *authRepo) UpdateUserProfile(ctx context.Context, req *pb.ProfileUpdateRequest) (*pb.ProfileUpdateResponse, error) {
	resp := pb.ProfileUpdateResponse{}

	q := `
	update 
		users 
	set 
		full_name=$1, 
		native_language=$2,
		updated_at=$3 
	where 
		id=$4
	returning
		id, 
		username, 
		email, 
		full_name, 
		native_language, 
		updated_at 
	`

	err := a.Db.QueryRow(q, req.Fullname, req.Nativelanguage, time.Now(), req.Id).Scan(
		&resp.Id,
		&resp.Username,
		&resp.Email,
		&resp.Fullname,
		&resp.Nativelanguage,
		&resp.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (a *authRepo) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.Message, error) {
	fmt.Println(req.CurrentPassword, req.NewPassword)
	passwordHash := HashPassword(req.CurrentPassword)
	newPassword := HashPassword(req.NewPassword)
	fmt.Println(passwordHash, newPassword)
	q := `
	update 
		users 
	set 
		password_hash=$1, 
		updated_at=$2 
	where 
		password_hash=$3
	`
	_, err := a.Db.Exec(q, newPassword, time.Now(), passwordHash)
	if err != nil {
		return nil, err
	}
	return &pb.Message{Message: "Password changed successfully"}, nil
}

func (a *authRepo) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.Message, error) {
	return nil, nil
}

func (a *authRepo) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.Message, error) {
	fmt.Println(req.Email, req.NewPassword, "-----------------")
	q := `
	update 
		users
	set 
		password_hash=$1,
		updated_at=$2
	where 
		email=$3
	AND
		deleted_at is null
	`
	_, err := a.Db.Exec(q, req.NewPassword, time.Now(), req.Email)
	if err != nil {
		fmt.Println(err, "11111111111111111")
		return nil, err
	}

	return &pb.Message{Message: "Password reset successfully"}, nil
}

func (a *authRepo) GetByUserId(ctx context.Context, req *pb.GetByUserIdRequest) (*pb.GetByUserIdResponse, error) {
	resp := &pb.GetByUserIdResponse{}
	q := `
	select 
		id,
		username,
		email, 
		password_hash,
		full_name, 
		native_language,	 
		created_at,
		updated_at
	from 
		users
	where 
		id = $1
	and 
		deleted_at is null
	`

	err := a.Db.QueryRow(q, req.Id).Scan(
		&resp.Id,
		&resp.Username,
		&resp.Email,
		&resp.Password,
		&resp.Fullname,
		&resp.Nativelanguage,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *authRepo) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	q := `
	update 
		users 
	set 
		deleted_at=$1 
	where 
		id=$2
	`
	_, err := a.Db.Exec(q, time.Now(), req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{Message: "User deleted successfully"}, nil
}
