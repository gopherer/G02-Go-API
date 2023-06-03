// Package user 存放用户 Model 相关逻辑
package user

import "G02-Go-API/models"

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"_"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
