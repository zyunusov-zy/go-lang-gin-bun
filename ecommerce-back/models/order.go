package models

import "time"

type Order struct {
	ID			int64		`bun:"id,pk,autoincrement" json:"id"`
	ItemID		int64 		`bun:"item_id, notnull" json:"item_id"`
	BuyerID		int64		`bun:"buyer_id, notnull" json:"buyer_id"`
	Quantity	int			`bun:"quantity, notnull" json:"quantity"`
	TotalPrice	float64		`bun:"total_price, notnull" json:"total_price"`
	Status		string		`bun:"status, default:'pending'" json:"status"`
	CreatedAt	time.Time   `bun:"created_at, default:current_timestamp" json:"created_at"`
}