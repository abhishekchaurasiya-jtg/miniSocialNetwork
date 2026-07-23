package customvalidators

import (
	models "app/src/models"
	strings "strings"

	binding "github.com/gin-gonic/gin/binding"
	validator "github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("genderValidator", genderValidator)
		v.RegisterValidation("maritalStatusValidator", maritalStatusValidator)
	}
}

var genderValidator validator.Func = func(fl validator.FieldLevel) bool {
	genderPtr, ok := fl.Field().Interface().(*string)
	if !ok || genderPtr == nil {
		return true // Allow nil/missing fields to pass validation
	}
	gender := strings.ToLower(*genderPtr)
	_, exists := models.GenderChoices[gender]
	return exists
}

var maritalStatusValidator validator.Func = func(fl validator.FieldLevel) bool {
	maritalStatusPtr, ok := fl.Field().Interface().(*string)
	if !ok || maritalStatusPtr == nil {
		return true
	}
	maritalStatus := strings.ToLower(*maritalStatusPtr)
	_, exists := models.MaritalStatusChoices[maritalStatus]
	return exists
}
