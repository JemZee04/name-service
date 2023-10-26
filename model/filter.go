package model

type FilterHuman struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	MinAge      int    `json:"minAge"`
	MaxAge      int    `json:"maxAge"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
	Page        int    `json:"page" binding:"required"`
}
