package routes

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/look4suman/events-api/models"
)

func fetchUsers(ctx *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		slog.Error("failed to fetch users", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch users"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		slog.Error("invalid user data", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user data"})
		return
	}
	u, err := user.Save()
	if err != nil {
		slog.Error("failed to create user", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully for id: " + strconv.FormatInt(u.ID, 10)})
}

func validateUser(user models.User) (bool, error) {
	dbUser, err := models.GetUserByEmail(user)
	if err != nil {
		return false, err
	}

	if dbUser != nil {
		return true, nil
	}

	return false, nil
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		slog.Error("invalid user data", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user data"})
		return
	}

	isValid, err := validateUser(user)
	if err != nil {
		slog.Error("Error while validating user", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error while validating user"})
		return
	}

	if isValid == false {
		slog.Info("Incorrect credentials")
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Incorrect credentials"})
		return
	}

	slog.Info("Logged in successfully")
	ctx.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
}
