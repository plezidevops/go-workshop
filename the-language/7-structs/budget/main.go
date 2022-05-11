package main

import (
	"fmt"
	"time"
)

// Budget is a budget for campaign
type Budget struct {
	CampaignID string
	Balance    float64
	Expires    time.Time
}

func (b Budget) TimeLeft() time.Duration {
	return b.Expires.Sub(time.Now().UTC())
}

func NewBudget(campaignID string, balance float64, expires time.Time) (*Budget, error) {
	if campaignID == "" {
		return nil, fmt.Errorf("empty campaignID")
	}

	if balance <= 0 {
		return nil, fmt.Errorf("balance must be greater than zero")
	}

	if expires.Before(time.Now()) {
		return nil, fmt.Errorf("bad expiration date")
	}

	b := Budget{
		campaignID,
		balance,
		expires,
	}

	return &b, nil
}

func main() {
	b1 := Budget{
		"Dog",
		54.90,
		time.Now().Add(7 * 24 * time.Hour),
	}

	fmt.Printf("%#v\n", b1)

	b2 := Budget{
		CampaignID: "Chicken",
		Balance:    234.90,
	}

	fmt.Println(b2)
}
