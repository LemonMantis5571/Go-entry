package tools

import (
	"time"
)

type mockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"testuser": {
		AuthToken: "testtoken",
		Username:  "testuser",
	},

	"testuser2": {
		AuthToken: "testtoken2",
		Username:  "testuser2",
	},

	"testuser3": {
		AuthToken: "testtoken3",
		Username:  "testuser3",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"testuser": {
		Coins:    100,
		Username: "testuser",
	},

	"testuser2": {
		Coins:    200,
		Username: "testuser2",
	},

	"testuser3": {
		Coins:    300,
		Username: "testuser3",
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

func (d *mockDB) SetupDatabase() error {
	return nil
}
