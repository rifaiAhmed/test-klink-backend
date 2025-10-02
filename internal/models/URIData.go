package models

type UriId struct {
	ID int `uri:"id" binding:"required"`
}

type ComponentServerSide struct {
	Limit     int    `json:"limit"`
	Skip      int    `json:"skip"`
	SortType  string `json:"sort_type"`
	SortBy    string `json:"sort_by"`
	Search    string `json:"search"`
	Offset    int    `json:"offset"`
	Condition string `json:"condition"`
	From      string `json:"from"`
	To        string `json:"to"`
}

type MultipleIngredients struct {
	RecipeID int    `json:"id" binding:"required"`
	Data     string `json:"data" binding:"required"`
}

type Option struct {
	Value int    `json:"value"`
	Label string `json:"label"`
}

type Option2 struct {
	Value string `json:"value"`
	Label string `json:"label"`
}
