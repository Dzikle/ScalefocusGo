package model

type List struct {
	Id   int    `json:"id,"`
	Name string `json:"name,"`
}

type Task struct {
	Id        int    `json:"id,omitempty"`
	Text      string `json:"text,omitempty"`
	ListId    int    `json:"list_id,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
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
