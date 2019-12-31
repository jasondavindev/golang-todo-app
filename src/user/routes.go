package user

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jasondavindev/golang-todo-app/src/common"
)

func UserRegister(r *gin.RouterGroup) {
	r.POST("/login", UserLogin)
	r.POST("/", UserRegistration)
	r.Use(AuthMiddleware())
	r.GET("/:name/tasks", UserTasks)
	r.GET("/:name", UserFind)
}

func UserFind(c *gin.Context) {
	userSearch := c.Param("name")

	user, err := FindOne(&User{Name: userSearch})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "ERROR"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user.Response()})
}

func UserRegistration(c *gin.Context) {
	userValidator := NewUserValidator()

	if err := userValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userValidator.UserModel.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userValidator.UserModel.Response()})
}

func UserLogin(c *gin.Context) {
	var cred User

	if err := common.Bind(c, &cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := FindOne(User{Email: cred.Email})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect credentials"})
		return
	}

	if err := user.checkPassword(cred.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect credentials"})
		return
	}

	expTime := time.Now().Add(time.Hour * 24)

	claims := Claims{
		Email: cred.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(common.GetJwtKey())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", tokenString, int(expTime.Unix()), "/", "localhost", true, true)
}

func UserTasks(c *gin.Context) {
	name := c.Param("name")
	user, err := FindOne(&User{Name: name})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks := user.GetTasks()
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
