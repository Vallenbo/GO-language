package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

type User struct {
	Name string `validate:"required,min=5,max=10"`
	Age  int    `validate:"gte=0,lte=150"`
}

func Test_StructValidation(*testing.T) {
	user := User{Name: "John", Age: 30}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println("未通过err :", err)
	} else {
		fmt.Println("Validation passed.")
	}
}
