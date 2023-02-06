package mtg

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Ruling contains additional rule information about the card.
type Ruling struct {
	// Date the information was released.
	Date string `json:"date"`
	// Text of the ruling hint.
	Text string `json:"text"`
}

// ForeignCardName represents the name of the card in an other language
type ForeignCardName struct {
	// Name is the name of the card in the given language
	Name string `json:"name"`
	// Language of the ForeignCardName
	Language string `json:"language"`
	// MultiverseID of the ForeignCardName (might be 0)
	MultiverseID uint `json:"multiverseid"`
}

// Legality stores information about legality notices for a specific format.
type Legality struct {
	// Format, such as Commander, Standard, Legacy, etc.
	Format string `json:"format"`
	// Legality for the given format such as Legal, Banned or Restricted.
	Legality string `json:"legality"`
}

// Card stores information about one single card.
type Card struct {
	// Name defines the name of the front of a card.
	// For split, double-faced and flip cards, the name of only one side.
	// Basically each ‘sub-card’ has its own record.
	Name string `json:"name"`
	// Names, only used for split, flip and dual cards.
	// Will contain all the names on this card, front or back.
	Names []string `json:"names"`
	// The ManaCost of a card. Consists of one or more mana symbols.
	// Use CMC and Colors to query.
	ManaCost string `json:"manaCost"`
	// Converted mana cost(CMC). Always a number.
	CMC float64 `json:"cmc"`
	// The card Colors. Usually derived from the casting cost.
	// Except for cards like the back of dual sided cards and Ghostfire.
	Colors []string `json:"colors"`
	// ColorIdentity defines card colors by color code.
	// Ex. [“Red”, “Blue”] becomes [“R”, “U”]
	ColorIdentity []string `json:"colorIdentity"`
	// Type defines card type. Seen in type line of printed card.
	// Note: The dash is a UTF8 "long dash" as per MTG rules.
	Type string `json:"type"`
	// Types defines multiple entries for Type
	// Seen on left of the dash in a card type. Examples: Instant, Sorcery,\
	// Artifact, Creature, Enchantment, Land, Planeswalker.
	Types []string `json:"types"`
	// Supertype. Appears to the far left of the card type.
	// Examples: Basic, Legendary, Snow, World, Ongoing.
	Supertypes []string `json:"supertypes"`
	// Subtypes. Appear after long dash following type.
	// Examples: Trap, Arcane, Equipment, Aura, Human, Rat, Squirrel.
	Subtypes []string `json:"subtypes"`
	// Rarity of card.
	// Examples: Common, Uncommon, Rare, Mythic Rare, Special, Basic Land.
	Rarity string `json:"rarity"`
	// Set defines what expansion set the card belongs to by set code.
	Set SetCode `json:"set"`
	// SetName defines name of expansion set the card belongs to.
	SetName string `json:"setName"`
	// Text defines oracle text of card.
	// May contain mana symbols and other symbols.
	// Text defines oracle text of card.
	// May contain mana symbols and other symbols.
	Text string `json:"text"`
	// Flavor defines the flavor text of card.
	Flavor string `json:"flavor"`
	// Artist defines the artist on the card.
	// This may not match the card, MTGJSON corrects card misprints.
	Artist string `json:"artist"`
	// Number defines card's set number. Appears bottom-center of the card.
	// NOTE: Set number can contain letters, strconv to int will error.
	Number string `json:"number"`
	// Power defines power of creature cards.
	// NOTE: Power can contain non-int, strconv to int will error.
	Power string `json:"power"`
	// Toughness defines toughness of creature cards.
	// NOTE: Toughness can contain non-int, strconv to int will error.
	Toughness string `json:"toughness"`
	// Loyalty defines loyalty of planeswalker cards.
	Loyalty string `json:"loyalty"`
	// Layout defines card's layout.
	// Examples: normal, split, flip, double-faced, token, plane, scheme,\
	// phenomenon, leveler, vanguard
	Layout string `json:"layout"`
	// MultiverseID defines the ID of the card on Wizard’s Gatherer web page.
	// Cards from sets that do not exist on Gatherer will NOT have a MultiverseID.
	// Sets not on Gatherer: ATH, ITP, DKM, RQS, DPA and all sets with a 4 letter\
	// code that starts with a lowercase 'p’.
	MultiverseID string `json:"multiverseid"`
	// Variations defines if a card has alternate art.
	// Examples:4 different Forests, or 2 Brothers Yamazaki.
	// Each other variation’s multiverseid will be listed.
	Variations []string `json:"variations"`
	// ImageURL defines URL for card image.
	// NOTE: Only for cards with a MultiverseID.
	ImageURL string `json:"imageUrl"`
	// Watermark defines the watermark on the card.
	// NOTE: Split cards don’t this field set.
	Watermark string `json:"watermark"`
	// Border defines the card border if it differs from top level JSON.
	// Example: Unglued silver borders, except for lands.
	Border string `json:"border"`
	// Timeshifted defines if the card was timeshifted in set.
	Timeshifted bool `json:"timeshifted"`
	// Hand defines the max hand size modifier.
	// NOTE: Only exists for Vanguard cards.
	// Hand defines the max hand size modifier.
	// NOTE: Only exists for Vanguard cards.
	Hand int `json:"hand"`
	// Life defines starting life total modifier.
	// NOTE: Only exists for Vanguard cards.
	Life int `json:"life"`
	// Reserved defines if card is reserved by Wizards Reprint Policy.
	Reserved bool `json:"reserved"`
	// ReleaseDate defines when this card was released.
	// NOTE: This is only set for promo cards. May not be accurate, some missing.
	// Only partial date may be set (YYYY-MM-DD or YYYY-MM or YYYY).
	ReleaseDate string `json:"releaseDate"`
	// Starter defines if card only released as part of core box set.
	Starter bool `json:"starter"`
	// Rulings define rulings for the card.
	Rulings []*Ruling `json:"rulings"`
	// ForeignNames defines foreign language name for the card.
	// Objects defined as "Language", "Name", and "MultiverseID" keys.
	// NOTE: Not available for all sets.
	ForeignNames []ForeignCardName `json:"foreignNames"`
	// Printings defines the sets the card was printed in (set codes).
	//Printings []SetCode `json:"printings"`
	// OriginalText defines text on card when it was first printed.
	// NOTE: Not available for promo cards.
	OriginalText string `json:"originalText"`
	// OriginalType defines type on card when it was first printed.
	// NOTE: Not available for promo cards.
	OriginalType string `json:"originalType"`
	// ID defines unique identification number of the card.
	// ID calculated by SHA1 hash of setCode + cardName + cardImageName.
	ID string `json:"id"`
	// Source defines where card was originally made available.
	// For box sets that are theme decks, it's the deck the card is from.
	Source string `json:"source"`
	// Legalities defines formats this card is legal, restricted or banned in.
	// Objects defined as "format" and "legality" keys.
	Legalities []Legality `json:"legalities"`
}

// ServerError is an error implementation for server messages.
type ServerError struct {
	// Status code given by the server
	Status string `json:"status"`
	// Message given by the server
	Message string `json:"error"`
}

// Error implements the error interface
func (s ServerError) Error() string {
	return s.Message
}

// cardResponse defines response from cards API Get request.
type cardResponse struct {
	Card  *Card   `json:"card"`
	Cards []*Card `json:"cards"`
}

func checkError(r *http.Response) error {
	if r.StatusCode == 200 {
		return nil
	}

	var sverr ServerError

	if err := json.NewDecoder(r.Body).Decode(&sverr); err != nil {
		return errors.New(r.Status)
	}
	return sverr
}

// Fetch collects card by ID or MultiverseID; retuns Card pointer.
func Fetch(filterID string) (*Card, error) {
	resp, err := http.Get(fmt.Sprintf("%scards/%s", queryURL, filterID))
	if err != nil {
		return nil, err
	}
	// body is io.Reader
	body := resp.Body
	defer body.Close()

	if err := checkError(resp); err != nil {
		return nil, err
	}
	cards, err := decodeCards(body)
	if err != nil {
		return nil, err
	}
	if len(cards) != 1 {
		return nil, fmt.Errorf("Card with ID %s not found", filterID)
	}
	return cards[0], nil
}
