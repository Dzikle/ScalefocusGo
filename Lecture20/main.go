package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Drink struct {
	Name            string `json:"strDrink"`
	StrInstructions string `json:"strInstructions"`
}

type Bartender struct {
	query  string
	Drinks []Drink
}

//Start
func Start() {
	var bar Bartender
	fmt.Println("Welcome to the GoLang Cocktail Bar")
	fmt.Println("What would you want to drink?")

	fmt.Scanln(&bar.query)

	if bar.query != "nothing" {
		bar.FetchDrinks()
	}
	bar.Done()
}

//Fetch Drink
func (bar *Bartender) FetchDrinks() {
	u, err := url.Parse("https://www.thecocktaildb.com/api/json/v1/1/search.php")
	q := u.Query()
	q.Add("s", bar.query)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(err)
	}
	drinks := Bartender{}
	json.NewDecoder(res.Body).Decode(&drinks)
	bar.Drinks = append(bar.Drinks, drinks.Drinks...)
	time.Sleep(time.Second)
	bar.Responce()
}

func (bar *Bartender) Responce() {
	fmt.Printf("The Drink %s is prepared on the following way:  \n", bar.Drinks[0].Name)

	responce := strings.Split(bar.Drinks[0].StrInstructions, ".")

	for i := 0; i < len(responce)-1; i++ {
		fmt.Println(i+1, responce[i])
	}
	fmt.Printf("\n And now you have  %s! \n", bar.Drinks[0].Name)
	fmt.Println("")
	Start()
}

func (bar *Bartender) Done() {
	fmt.Println("You have finished Drinking, DONT DRIVE")
	os.Stdout.Close()
}

func main() {
	Start()
}
