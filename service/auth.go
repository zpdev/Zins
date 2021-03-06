package service

import (
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	cons "github.com/zpdev/zins/common/constance"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/common/utils"
	"github.com/zpdev/zins/model"
	"time"
)

var AuthService = &authService{}

type authService struct {
}

func (sec *authService) Login(db *gorm.DB, user *model.User) (*model.User, *model.Token, *errutils.ZinError) {
	var loginUser model.User
	if db.Where("Username = ?", user.Username).First(&loginUser).RecordNotFound() {
		return nil, nil, errutils.SpecifiedUserNotFound(user.Username)
	}
	encryptPass, enErr := base64.StdEncoding.DecodeString(loginUser.Password)
	if enErr != nil {
		return nil, nil, errutils.PasswordVerifyError()
	}
	result, err := utils.VerifyPassphrase(user.Password, encryptPass)
	if err != nil {
		return nil, nil, errutils.PasswordVerifyError()
	}
	if !result {
		return nil, nil, errutils.UserPassError()
	}
	loginUser.Password = ""
	token, tErr := sec.genToken(db, &loginUser)
	if tErr != nil {
		return nil, nil, errutils.LoginFailed()
	}

	return &loginUser, token, nil
}

func (sec *authService) genToken(db *gorm.DB, user *model.User) (*model.Token, *errutils.ZinError) {
	tokenId := uuid.NewV4()
	timeNow := time.Now()
	h, err := time.ParseDuration("0.5h")
	if err != nil {
		return nil, errutils.LoginFailed()
	}
	expiredTime := timeNow.Add(h)
	token := model.Token{
		Token:       tokenId.String(),
		UserID:      user.ID,
		Status:      cons.ACTIVE,
		CreatedTime: timeNow.Unix(),
		ExpiredTime: expiredTime.Unix(),
	}
	if err := db.Create(&token).Error; err != nil {
		return nil, errutils.DBOperationsFailed(err.Error())
	}
	return &token, nil
}

func (sec *authService) Logout(db *gorm.DB, user *model.User) *errutils.ZinError {
	return nil
}
