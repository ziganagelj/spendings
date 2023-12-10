package domain

import "errors"

type Category string

const (
	GROCERIES      Category = "Groceries"
	SHOPPING       Category = "Shopping"
	RESTAURANT     Category = "Restaurant"
	LOAN           Category = "Loan"
	SUBSCRIPTION   Category = "Subscription"
	UTILITIES      Category = "Utilities"
	TAX            Category = "Tax"
	SPORTS         Category = "Sports"
	ECHO           Category = "Echo"
	HEALTH         Category = "Health"
	TRANSPORTATION Category = "Transportation"
	RENT           Category = "Rent"
)

func GetAllCategories() []Category {
	return []Category{GROCERIES, SHOPPING, RESTAURANT, LOAN, SUBSCRIPTION, UTILITIES, TAX, SPORTS, ECHO, HEALTH, TRANSPORTATION, RENT}
}

func ParseCategory(s string) (Category, error) {
	switch Category(s) {
	case GROCERIES, SHOPPING, RESTAURANT, LOAN, SUBSCRIPTION, UTILITIES, TAX, SPORTS, ECHO, HEALTH, TRANSPORTATION, RENT:
		return Category(s), nil
	default:
		return "", errors.New("invalid category")
	}
}
