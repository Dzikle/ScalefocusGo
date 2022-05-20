package entity

type List struct {
	Id   int    `gorm:"primary_key;auto_increment" json:"id"`
	Name string `json:"name" gorm:"type:varchar(200)"`
}

type Task struct {
	Id        int    `gorm:"primary_key;auto_increment" json:"id"`
	Text      string `json:"text,omitempty" gorm:"type:varchar(200)"`
	ListId    List   `json:"list_id,omitempty" gorm:"foreignkey:ListId"`
	Completed bool   `json:"completed,omitempty" gorm:"type:boolean(200)"`
}

type User struct {
	Username string `json:"username,omitempty" gorm:"type:varchar(255);unique;not null"`
	Password string `json:"password,omitempty" gorm:"type:varchar(200)"`
}
type Weather struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather,omitempty"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main,omitempty"`
	City string `json:"name,omitempty"`
}
type WeatherInfo struct {
	FormatedTemp string `json:"formated_temp,omitempty"`
	Description  string `json:"description,omitempty"`
	City         string `json:"city,omitempty"`
}
