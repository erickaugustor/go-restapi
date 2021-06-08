package types

type Deck struct {
	deck_id   int
	shuffled  bool
	remaining string
	cards     []Card
}
