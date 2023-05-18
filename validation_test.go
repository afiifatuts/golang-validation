package golangvalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	// sama dengan var validate *validator.Validate = validator.New()
	validate := validator.New()

	if validate == nil {
		t.Error("Validate is nil")
	}
}

func TestValidationVariable(t *testing.T) {
	validate := validator.New()

	users := "tsaani"
	err := validate.Var(users, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidateTwoVariables(t *testing.T) {
	validate := validator.New()

	pw := "rahasia"
	pw2 := "rahasia"

	err := validate.VarWithValue(pw, pw2, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}
