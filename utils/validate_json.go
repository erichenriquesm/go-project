package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ValidateJSON faz o bind + validação automática
func ValidateJSON(ctx *gin.Context, obj interface{}) map[string]string {
	if err := ctx.ShouldBindJSON(obj); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make(map[string]string)
			for _, fe := range ve {
				out[fe.Field()] = validationMessage(fe)
			}
			return out
		}

		// Se for erro de parse do JSON
		return map[string]string{"json": "JSON inválido"}
	}
	return nil
}

func validationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Este campo é obrigatório"
	case "min":
		return "Valor muito curto"
	case "max":
		return "Valor muito longo"
	case "gt":
		return "Deve ser maior que 0"
	}
	return "Valor inválido"
}
