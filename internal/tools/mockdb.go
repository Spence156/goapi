package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
		Currency: "EUR",
	},
	"jason": {
		Coins:    200,
		Username: "jason",
		Currency: "GBP",
	},
	"marie": {
		Coins:    300,
		Username: "marie",
		Currency: "USD",
	},
}

var mockUserDetails = map[string]UserDetails{
	"alex": {
		Username: "alex",
		FullName: "Alex Fulller",
		Country:  "United Kingdom",
		Email:    "Alex.Fuller@gmail.com",
	},
	"jason": {
		Username: "jason",
		FullName: "Jason Bourne",
		Country:  "United States",
		Email:    "Jason.Bourne@gmail.com",
	},
	"marie": {
		Username: "marie",
		FullName: "Marie Verns",
		Country:  "France",
		Email:    "Marie.Verns@gmail.com",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserDetails(username string) *UserDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = UserDetails{}
	clientData, ok := mockUserDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
