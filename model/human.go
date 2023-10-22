package model

type Human struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Gender      string `json:"sex"` // Посмотреть в каком виде дают пол
	Nationality string `json:"nationality"`
}
