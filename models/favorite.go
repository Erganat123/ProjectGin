package models

import "time"

type Favorite struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    UserID    string    `json:"user_id"`
    BookID    uint      `json:"book_id"`
    CreatedAt time.Time `json:"created_at"`
}