package config

import "github.com/maribowman/gin-skeleton/app/model"

type AuthenticationConfig struct {
	HmacEnabled  bool
	AllowedUsers []model.User
}
