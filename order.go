package sendowl

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type OrderID int

// UnmarshalJSON implements the json.Unmarshaler interface.
func (id *OrderID) UnmarshalJSON(data []byte) error {
	var i int64
	if err := json.Unmarshal(data, &i); err != nil {
		return fmt.Errorf("sendowl: OrderID should be an int64, got %T: %v", data, data)
	}
	*id = OrderIDFromInt(i)
	return nil
}

func (id OrderID) String() string {
	return strconv.Itoa(int(id))
}

func (id OrderID) Int() int64 {
	return int64(id)
}

func OrderIDFromString(s string) OrderID {
	i, _ := strconv.Atoi(s)
	return OrderID(i)
}

func OrderIDFromInt(i int64) OrderID {
	return OrderID(i)
}

type Order struct {
	ID   OrderID `json:"id"` // ID of the order.
	Cart struct {
		Items []struct {
			Item struct {
				ProductID        `json:"product_id"`
				PackageID        int        `json:"package_id"`
				Quantity         int        `json:"quantity"`
				PriceAtCheckout  Price      `json:"price_at_checkout"`
				ValidUntil       *time.Time `json:"valid_until"`
				DownloadAttempts int        `json:"download_attempts"`
			} `json:"cart_item"`
		} `json:"cart_items"`
		CompletedCheckoutAt time.Time `json:"completed_checkout_at"`
		StartedCheckoutAt   time.Time `json:"started_checkout_at"`
	} `json:"cart"`
}
