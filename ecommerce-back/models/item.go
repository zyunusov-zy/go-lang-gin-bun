package models

import "time"

type Item struct {
	ID          int64     `bun:"id,pk,autoincrement" json:"id"`
	Name        string    `bun:"name,notnull" json:"name"`
	Description string    `bun:"description,notnull" json:"description"`
	Price       float64   `bun:"price,notnull" json:"price"`
	ImageURL    string    `bun:"image_url" json:"image_url"` // Optional
	SellerID    int64     `bun:"seller_id,notnull" json:"seller_id"`
	CreatedAt   time.Time `bun:"created_at,default:current_timestamp" json:"created_at"`
}
