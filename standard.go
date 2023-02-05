package mtg

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

// StandardSetNames returns map of set names in Standard.
func StandardSetNames() (map[string]string, error) {
	cards, err := StandardCards()
	if err != nil {
		return nil, err
	}
	sets := make(map[string]string)
	for _, card := range cards {
		sets[card.SetName] = card.Set.String()
	}
	return sets, nil
}
