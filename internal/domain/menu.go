package domain

type Menu struct {
	Items []Item `json:"items" genai:"description=Menu items list;required"`
}

type Item struct {
	OriginalItemName   string   `json:"originalItemName" genai:"description=Item name in original language;required"`
	TranslatedItemName string   `json:"translatedItemName" genai:"description=Trnalated item name;required"`
	Label              string   `json:"label" genai:"description=JSON array label;required"`
	BoundingBox        []int    `json:"box_2d" genai:"description=Item name;required"`
	Price              float32  `json:"price" genai:"description=Price of the item"`
	AvailableOptions   []string `json:"availableOptions" genai:"description=An array of translated option names"`
}

type MenuExplanation struct {
	Name            string `json:"name" genai:"description=Menu name.;required"`
	Description     string `json:"description" genai:"description=The most general description of the menu.;required"`
	Ingredients     string `json:"ingredients" genai:"description=What ingredients usually are included in.;required"`
	WhatToBeCareful string `json:"whatToBeCareful" genai:"description=What ingredients to be careful of, referring to the user dislike foods.;required"`
}
