package service

import (
	"strings"

	"github.com/Chihaya-Anon123/task_manager/internal/code"
	"github.com/Chihaya-Anon123/task_manager/internal/dao"
	"github.com/Chihaya-Anon123/task_manager/internal/errs"
	"github.com/Chihaya-Anon123/task_manager/internal/model"
	"github.com/Chihaya-Anon123/task_manager/internal/utils"
)

type RegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type RegisterOutput struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

func Register(input RegisterInput) (*RegisterOutput, error) {
	input.Username = strings.TrimSpace(input.Username)
	input.Nickname = strings.TrimSpace(input.Nickname)

	if input.Username == "" {
		return nil, errs.New(code.CodeInvalidParams, "username is required")
	}

	if input.Password == "" {
		return nil, errs.New(code.CodeInvalidParams, "password is required")
	}

	if len(input.Password) < 6 {
		return nil, errs.New(code.CodeInvalidParams, "password must be at least 6 characters")
	}

	existUser, err := dao.GetUserByUsername(input.Username)
	if err != nil {
		return nil, errs.ErrDBError
	}
	if existUser != nil {
		return nil, errs.New(code.CodeInvalidParams, "username already exists")
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, errs.ErrInternalServer
	}

	user := &model.User{
		Username: input.Username,
		Password: hashedPassword,
		Nickname: input.Nickname,
	}

	if err := dao.CreaterUser(user); err != nil {
		return nil, errs.ErrDBError
	}

	return &RegisterOutput{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
	}, nil
}
