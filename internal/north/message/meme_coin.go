package message

type CreateMemeCoinCommand struct {
	Name        string
	Description string
}

type GetMemeCoinQuery struct {
	ID string
}

type UpdateMemeCoinCommand struct {
	ID          string
	Description string
}

type DeleteMemeCoinCommand struct {
	ID string
}

type PokeMemeCoinCommand struct {
	ID string
}
