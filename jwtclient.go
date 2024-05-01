package ApiWebSite

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	Url   string
	Token string
}

type ShopPositionAP struct {
	Name        string `json:"name"`
	Caption     string `json:"caption"`
	Discription string `json:"discription"`
	Price       string `json:"price"`
}

type UsersAP struct {
	User        string `json:"user"`
	PhoneNumber string `json:"phone"`
	FirstName   string `json:"fname"`
	LastName    string `json:"sname"`
}

type BasketAP struct {
	User        string `json:"user"`
	Name        string `json:"name"`
	Caption     string `json:"caption"`
	Discription string `json:"discription"`
	Price       string `json:"price"`
	Qty         string `json:"qty"`
}

type OrdersAP struct {
	User        string `json:"user"`
	Name        string `json:"name"`
	Caption     string `json:"caption"`
	Discription string `json:"discription"`
	Price       string `json:"price"`
	Qty         string `json:"qty"`
	Date        string `json:"date"`
}

// Method fo get user information
func (c Client) GetUser() []UsersAP {
	var resdata = make(map[string][]UsersAP)

	// Make request.  See func restrictedHandler for example request processor
	// Build the request
	req, err := http.NewRequest("GET", c.Url, nil)
	if err != nil {
		fmt.Println("Error is req: ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in send req: ", err)
	}

	// Defer the closing of the body
	defer resp.Body.Close()
	// Fill the data with the data from the JSON

	// Use json.Decode for reading streams of JSON data
	json.NewDecoder(resp.Body).Decode(&resdata)

	return resdata["user"]
}

// Method fo get shop position
func (c Client) GetShop() []ShopPositionAP {
	var resdata = make(map[string][]ShopPositionAP)

	// Make request.  See func restrictedHandler for example request processor
	// Build the request
	req, err := http.NewRequest("GET", c.Url, nil)
	if err != nil {
		fmt.Println("Error is req: ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in send req: ", err)
	}

	// Defer the closing of the body
	defer resp.Body.Close()
	// Fill the data with the data from the JSON

	// Use json.Decode for reading streams of JSON data
	json.NewDecoder(resp.Body).Decode(&resdata)

	return resdata["shop"]
}

// Method fo get basket position from all users
func (c Client) GetBasket() []BasketAP {
	var resdata = make(map[string][]BasketAP)

	// Make request.  See func restrictedHandler for example request processor
	// Build the request
	req, err := http.NewRequest("GET", c.Url, nil)
	if err != nil {
		fmt.Println("Error is req: ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in send req: ", err)
	}

	// Defer the closing of the body
	defer resp.Body.Close()
	// Fill the data with the data from the JSON

	// Use json.Decode for reading streams of JSON data
	json.NewDecoder(resp.Body).Decode(&resdata)

	return resdata["basket"]
}

// Method fo get order all users
func (c Client) GetOrders() []OrdersAP {
	var resdata = make(map[string][]OrdersAP)

	// Make request.  See func restrictedHandler for example request processor
	// Build the request
	req, err := http.NewRequest("GET", c.Url, nil)
	if err != nil {
		fmt.Println("Error is req: ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in send req: ", err)
	}

	// Defer the closing of the body
	defer resp.Body.Close()
	// Fill the data with the data from the JSON

	// Use json.Decode for reading streams of JSON data
	json.NewDecoder(resp.Body).Decode(&resdata)

	return resdata["order"]
}
