package mtg

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

// StandardCards returns slice of cards in Standard.
func StandardCards() ([]*Card, error) {
	// NewQuery is mtg.Query.
	query := NewQuery().Where(CardGameFormat, "Standard")
	// cards is mtg.[]*Card
	cards, err := query.Where(CardLegality, "Legal").All()
	if err != nil {
		return nil, err
	}

	return cards, nil
}

// StandardSets returns map of set names in Standard.
func StandardSets() (map[string]SetCode, error) {
	URL := "https://whatsinstandard.com/api/v6/standard.json"
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var stdResp standardResp
	if err := json.Unmarshal(body, &stdResp); err != nil {
		return nil, err
	}

	standardSets := make(map[string]SetCode)
	for _, setItem := range stdResp.Sets {
		isStandard, err := parseDates(
			setItem.EnterDate.Exact,
			setItem.ExitDate.Exact,
		)
		if err != nil {
			return nil, err
		}

		if isStandard {
			standardSets[setItem.Name] = setItem.Code
		}
	}
	return standardSets, nil
}

func parseDates(enter, exit string) (bool, error) {
	const longForm = "2006-01-02 15:04:05"
	currentDate := time.Now()
	// Parse date strings into usable time.Time format.
	formatDate := func(date string) (time.Time, error) {
		date = strings.Replace(strings.Split(date, ".")[0], "T", " ", 1)
		return time.Parse(longForm, date)
	}

	// Validate enter date.
	var enterValidated bool
	if enter != "" {
		enterDate, err := formatDate(enter)
		if err != nil {
			return false, err
		}
		enterValidated = enterDate.Local().Before(currentDate)
	}

	// If enter is empty, the set is in the future.
	if enter == "" {
		return false, nil
	}

	// Validate exit date.
	var exitValidated bool
	if exit != "" {
		exitDate, err := formatDate(exit)
		if err != nil {
			return false, err
		}
		exitValidated = exitDate.Local().After(currentDate)
	}

	// If exit is empty, the set has not yet left standard.
	if exit == "" {
		exitValidated = true
	}

	// Compare w/current date.
	if enterValidated && exitValidated {
		return true, nil
	}
	return false, nil
}

// standardResp defines the JSON response whatisinstandard.
type standardResp struct {
	Deprecated bool   `json:"deprecated"`
	Sets       []sets `json:"sets"`
}

type enterDate struct {
	Exact string `json:"exact"`
}

type exitDate struct {
	Exact string `json:"exact"`
}

type sets struct {
	Name      string    `json:"name"`
	Code      SetCode   `json:"code"`
	EnterDate enterDate `json:"enterDate"`
	ExitDate  exitDate  `json:"exitDate"`
}
