package main

type Game struct {
	Field *Field
}

func CreateGame(opts *FieldOpts) (*Game, error) {
	field, err := CreateField(opts)
	if err != nil {
		return nil, err
	}

	return &Game{Field: field}, nil
}

func (game *Game) Play() error {
	return nil
}
