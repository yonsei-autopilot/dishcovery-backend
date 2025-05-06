package domain

type Menu struct {
	Items []Item `json:"items" genai:"description=Menu items list;required"`
}

type Item struct {
	Name        string  `json:"name" genai:"description=Item name;required"`
	Description string  `json:"description" genai:"description=Description of the item. Do not include any item name or price info.;required"`
	Price       float32 `json:"price" genai:"description=Price of item"`
}

type MenuExplanation struct {
	Name            string `json:"name" genai:"description=Menu name.;required"`
	Description     string `json:"description" genai:"description=The most general description of the menu.;required"`
	Ingredients     string `json:"ingredients" genai:"description=What ingredients usually are included in.;required"`
	WhatToBeCareful string `json:"whatToBeCareful" genai:"description=What ingredients to be careful of, referring to the user dislike foods.;required"`
}
