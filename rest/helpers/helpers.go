package helpers

import (
	"errors"
	"net/http"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	default:
		return "is not valid"
	}
}

func PascalToCamel(s string) string {
	if s == "" {
		return s
	}

	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func BodyValidation[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body T

		err := ctx.ShouldBindJSON(&body)

		validationErrors := make(map[string]string)

		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				validationErrors[PascalToCamel(e.Field())] = GetErrorMessage(e)
			}
		} else if err != nil {
			var timeErr *time.ParseError

			if errors.As(err, &timeErr) {
				validationErrors["dateTime"] = "must be ISO8601 (RFC3339), e.g. 2006-01-02T15:04:05Z"
			} else {
				validationErrors["body"] = err.Error()
			}
		}

		if err != nil {

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
