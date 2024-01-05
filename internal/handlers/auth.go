package handlers

import (
	"community-qa-platform/internal/models"
	"community-qa-platform/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// 引入其他必要的包，例如用于数据库操作的包，JWT处理的包等
)

// UserCredentials 用于登录请求的结构
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthHandler struct {
	service *service.UserService
}

func NewAuthHandler(service *service.UserService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Register 处理用户注册的请求
func (a *AuthHandler) Register(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := a.service.CreateUser(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// Login 处理用户登录的请求
func (a *AuthHandler) Login(c *gin.Context) {
	var creds UserCredentials

	// 将请求的JSON绑定到UserCredentials结构体
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if ok, err := a.service.CheckUserNamePassword(creds.Username, creds.Password); ok {
		log.Print(err)
		return
	}

	// 根据creds验证用户身份
	// 如果身份验证失败，则返回错误
	// 如果成功，生成JWT并返回给客户端

	// 示例响应
	c.JSON(http.StatusOK, gin.H{"message": "登录成功", "token": "jwt_token_here"})
}

// JwtMiddleware 是一个JWT验证的中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 实现JWT验证
		// 例如：从请求头中获取token，解析和验证token

		// 如果验证失败，设置适当的响应并返回
		// c.AbortWithStatus(http.StatusUnauthorized)

		// 如果验证成功，将用户信息设置到上下文中，例如
		// c.Set("user", user)

		c.Next()
	}
}
