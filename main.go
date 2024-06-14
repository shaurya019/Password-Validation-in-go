package main

import (
    "github.com/gin-gonic/gin"
    "gopkg.in/go-playground/validator.v9"
    "net/http"
    "regexp"
)

// User represents the structure of the request payload
type User struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,password"`
}

// Define a global validator
var validate *validator.Validate

// passwordValidation function to validate password strength
func passwordValidation(fl validator.FieldLevel) bool {
    password := fl.Field().String()
    var uppercase = regexp.MustCompile(`[A-Z]`)
    var lowercase = regexp.MustCompile(`[a-z]`)
    var number = regexp.MustCompile(`[0-9]`)
    var special = regexp.MustCompile(`[!@#~$%^&*()_+|<>?:{}]`)

    if len(password) < 8 || !uppercase.MatchString(password) || !lowercase.MatchString(password) || !number.MatchString(password) || !special.MatchString(password) {
        return false
    }
    return true
}

func main() {
    r := gin.Default()

    // Initialize the validator
    validate = validator.New()
    // Register the password validation function
    validate.RegisterValidation("password", passwordValidation)

    r.POST("/validateUser", func(c *gin.Context) {
        var user User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := validate.Struct(user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "User data is valid"})
    })

    r.Run(":8080")
}


// passwordValidation(fl validator.FieldLevel) bool: This function implements the custom logic to validate the password. It checks for a minimum length of 8 characters and ensures that the password contains at least one uppercase letter, one lowercase letter, one digit, and one special character.

// The validate.RegisterValidation("password", passwordValidation) line registers the custom password validation function.