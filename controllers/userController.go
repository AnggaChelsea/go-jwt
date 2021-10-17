package controllers

import (
	"context"
	"fmt"
	"go-jwt/helpers"
	helper "go-jwt/helpers"
	"go-jwt/models"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = databases.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()

func VerifyPassword()

func Signup()

func Signin()

func GetUser()

func GetUserId()