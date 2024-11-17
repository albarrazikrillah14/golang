package golang_validation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		t.Error("Validate is nil")
	}
}

func TestValidationField(t *testing.T) {
	validate := validator.New()
	user := "Albarra Zikrillah"

	err := validate.Var(user, "required")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidateTwoVariable(t *testing.T) {
	validate := validator.New()

	password := "rahasia"
	confirmPassword := "rahasia"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()

	user := "1405"
	err := validate.Var(user, "required,number")
	if err != nil {
		fmt.Println(err)
	}
}

func TestTagParameter(t *testing.T) {
	validate := validator.New()

	number := 1405
	err := validate.Var(number, "required,min=0,max=2000")
	if err != nil {
		fmt.Println(err)
	}
}

func TestStruct(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	user := &LoginRequest{
		Username: "albarrazikrillah1405gmail.com",
		Password: "P_assword001",
	}

	validator := validator.New()

	err := validator.Struct(user)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationErrors(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	user := &LoginRequest{
		Username: "albarrazikrillah1405gmail.com",
		Password: "P_assword001",
	}

	validate := validator.New()

	err := validate.Struct(user)

	if err != nil {
		ve := err.(validator.ValidationErrors)
		for _, fieldError := range ve {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag())
		}
	}
}

func TestStructCrossField(t *testing.T) {
	type RegisterUser struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	user := &RegisterUser{
		Username:        "albarrazikrillah1405@gmail.com",
		Password:        "P_assword001",
		ConfirmPassword: "P_assword001",
	}

	validator := validator.New()

	err := validator.Struct(user)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
	}

	user := &User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
	}

	validate := validator.New()

	err := validate.Struct(user)

	if err != nil {
		fmt.Println(err)
	}
}
