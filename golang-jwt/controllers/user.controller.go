package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/irarelycode/golang-jwt/database"
	"github.com/irarelycode/golang-jwt/helpers"
	"github.com/irarelycode/golang-jwt/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var UserCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(c *gin.Context) {}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	return true, ""
}

func Signup(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
		return
	}
	count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	if count > 0 {
		log.Panic("Email already exists")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
	}
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.UserId = user.ID.Hex()
	token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.FirstName, *user.LastName, *user.UserType, *&user.UserId)
	user.Token = &token
	user.RefreshToken = &refreshToken
	result, err := UserCollection.InsertOne(ctx, user)
	if err != nil {
		msg := fmt.Sprintf("User item was not created..")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result)
}

func Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User
	var foundUser models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Email or password is incorrect"})
	}
	// isValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
}

func GetUser(c *gin.Context) {
	userId := c.Param("user_id")
	if err := helpers.MatchUserTypeToUUid(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User
	err := UserCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {}
