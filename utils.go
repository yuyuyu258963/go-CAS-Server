package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 登录时携带的信息
type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Service  string `json:"service"`
}

// 验证用户信息
func (p *LoginInfo) Check() (flag bool) {
	flag = false
	// 这里本该是向数据库查找比对的
	// 这里简化为设定的用户
	flag = p.Username == configData.AdminUserInfo.Name && p.Password == configData.AdminUserInfo.Password
	return
}

// 获得用户登录时的信息
func InitLoginInfo(username string, password string, service string) *LoginInfo {
	return &LoginInfo{Username: username, Password: password, Service: service}
}

// 获得请求中的各种参数
func PraseLoginInfoFromRequest(c *gin.Context) *LoginInfo {
	return InitLoginInfo(
		c.DefaultQuery("username", "none"),
		c.DefaultQuery("password", "none"),
		c.DefaultQuery("service", ""),
	)
}

// 限制service的跳转地址
func checkAllowedDomains(urlPath string) (allowed bool) {
	allowed = false
	for _, val := range configData.AllowDomain {
		if strings.Contains(urlPath, val) {
			allowed = true
			break
		}
	}
	return
}

// 获得 uuid
func getUUid() (uid string) {
	uid = uuid.NewString()
	uid = strings.ReplaceAll(uid, "-", "")
	return
}
