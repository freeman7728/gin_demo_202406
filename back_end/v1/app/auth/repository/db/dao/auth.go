package dao

import (
	"context"
	"database_lesson/app/auth/repository/db/model"
	"gorm.io/gorm"
)

type AuthDao struct {
	*gorm.DB
}

func NewAuthDao(ctx context.Context) *AuthDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &AuthDao{NewDBClient(ctx)}
}

// RegisterToken 先删除再插入新token
func (a *AuthDao) RegisterToken(auth *model.Auth) (tx *gorm.DB) {
	delToken := `delete from auth where id = ?`
	delResult := a.Exec(delToken, auth.Id).Scan(nil)
	if delResult.Error != nil {
		return delResult
	}
	insertToken := `insert into auth (id, token, level) values (?, ?, ?)`
	insertResult := a.Exec(insertToken, auth.Id, auth.Token, auth.Level).Scan(nil)
	return insertResult
}
func (a *AuthDao) GetLevel(auth *model.Auth) (tx *gorm.DB) {
	getLevel := `select level from employee where id = ?`
	getResult := a.Raw(getLevel, auth.Id).Scan(&auth.Level)
	return getResult
}
func (a *AuthDao) GetToken(auth *model.Auth, token *string) (tx *gorm.DB) {
	getToken := `select token from auth where id = ?`
	getResult := a.Raw(getToken, auth.Id).Scan(&token)
	return getResult
}
func (a *AuthDao) GetLevelByToken(auth *model.Auth, l *int) (tx *gorm.DB) {
	getLevel := `select level from employee where token = ?`
	getResult := a.Raw(getLevel, auth.Token).Scan(l)
	return getResult
}
