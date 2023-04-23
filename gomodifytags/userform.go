package main

import "time"

// UserForm struct
type UserForm struct {
	Name     string    `validate:"required|min_len:7" message:"required:{field} is required" label:"User Name"`
	Email    string    `validate:"email" message:"email is invalid" label:"User Email"`
	Age      int       `validate:"required|int|min:1|max:99" message:"int:age must int|min:age min value is 1"`
	CreateAt int       `validate:"min:1"`
	Safe     int       `validate:"-"`
	UpdateAt time.Time `validate:"required" message:"update time is required"`
	Code     string    `validate:"customValidator"`
	// ExtInfo nested struct
	ExtInfo struct {
		Homepage string `validate:"required" label:"Home Page"`
		CityName string
	} `validate:"required" label:"Home Page"`
}

// CustomValidator custom validator in the source struct.
func (f UserForm) CustomValidator(val string) bool {
	return len(val) == 4
}
