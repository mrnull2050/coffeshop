package handler

import (
	repo "coffeshop/Repo"
	"coffeshop/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Repo *repo.UserRepo
}

func NewUserHanlder(repo *repo.UserRepo) *UserHandler {
	return &UserHandler{Repo: repo}
}
func (uh *UserHandler) CreatUser(c *gin.Context) {
	var user model.UserModel
	c.BindJSON(&user)

	err := uh.Repo.Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Result": "error while creating user",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Result": "User Created",
		"User":   user,
	})

}
func (uh *UserHandler) GetAll(c *gin.Context) {

	users, err := uh.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "err while getting users",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "getting all users!!",
		"users ": users,
	})

}
func (uh UserHandler) GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := uh.Repo.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "user not found",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "user found",
		"user":   user,
	})

}
func (uh *UserHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var user model.UserModel
	c.BindJSON(&user)
	user.ID = id

	err := uh.Repo.Update(user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "user Not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "user update",
	})
}

func(uh *UserHandler)DeleteUser(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	err := uh.Repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound , gin.H{
			"result" : "not found",
		})
		return
	}
	c.JSON(http.StatusOK , gin.H{
		"result" : "user deleted" , 
	})
}