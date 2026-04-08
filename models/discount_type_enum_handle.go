package models

import (
	"encoding/json"
	"errors"
	"strings"
)

type DiscountType string

const (
	DiscountPercent DiscountType = "percent"
	DiscountFixed   DiscountType = "fixed"
)

// >> MarshalJSON: Will be saved as an enum value in the database
func (d DiscountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}

// >>> Convert
func (d *DiscountType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	s = strings.TrimSpace(s)
	switch s {
	case "%", "percent", "Percentage":
		*d = DiscountPercent
	case "fixed", "amount", "৳":
		*d = DiscountFixed
	default:
		return errors.New("invalid discount_type")
	}

	return nil
}