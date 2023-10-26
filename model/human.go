package model

import "errors"

type Human struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Surname     string `json:"surname" db:"surname" binding:"required"`
	Patronymic  string `json:"patronymic" db:"patronymic"`
	Age         int    `json:"age" db:"age"`
	Gender      string `json:"gender" db:"gender"`
	Nationality string `json:"nationality" db:"nationality"`
}

type UpdateHumanInput struct {
	Name        *string `json:"name"`
	Surname     *string `json:"surname"`
	Patronymic  *string `json:"patronymic"`
	Age         *int    `json:"age"`
	Gender      *string `json:"gender"`
	Nationality *string `json:"nationality"`
}

func (i UpdateHumanInput) Validate() error {
	if i.Age == nil && i.Gender == nil && i.Name == nil && i.Nationality == nil && i.Surname == nil && i.Patronymic == nil {
		return errors.New("update struct has no values")
	}

	return nil
}
