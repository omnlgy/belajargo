package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func GetErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "is required"
	case "email":
		return "must be a valid email"
	case "alpha":
		return "must contain only letters"
	case "alphanum":
		return "must contain only letters and numbers"
	case "alphaspace":
		return "must contain only letters and spaces"
	case "alphanumspace":
		return "must contain only letters, numbers and spaces"
	case "min":
		return "must be at least " + e.Param() + " characters"
	case "max":
		return "must be at most " + e.Param() + " characters"
	default:
		return "invalid value"
	}
}

func PascalToCamel(s string) string {
	if s == "" {
		return s
	}

	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func handleBindingError(err error) map[string]string {
	validationErrors := make(map[string]string)

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			validationErrors[PascalToCamel(e.Field())] = GetErrorMessage(e)
		}
		return validationErrors
	}

	var timeErr *time.ParseError
	var jsonErr *json.SyntaxError

	switch {
	case errors.As(err, &timeErr):
		validationErrors["dateTime"] = "must be ISO8601 (RFC3339), e.g. 2006-01-02T15:04:05Z"
	case errors.As(err, &jsonErr):
		validationErrors["body"] = "must be valid JSON"
	case err.Error() == "EOF":
		validationErrors["body"] = "request body cannot be empty"
	default:
		validationErrors["body"] = err.Error()
	}

	return validationErrors
}

func BodyValidation[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body T
		err := ctx.ShouldBindJSON(&body)

		if err != nil {
			validationErrors := handleBindingError(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request body",
				"errors":  validationErrors,
			})
			return
		}

		ctx.Set("data", body)
		ctx.Next()
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AuthMidelware(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	// Remove "Bearer " prefix
	tokenString := authHeader
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	}

	claim, err := ValidateToken(tokenString)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	ctx.Set("claim", claim)
	ctx.Next()
}
