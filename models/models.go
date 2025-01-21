package models

import (
	"database/sql"
	"time"
)

// type User struct {
// 	ID           int       `json:"id"`
// 	Username     string    `json:"username"`
// 	Email        string    `json:"email"`
// 	PasswordHash string    `json:"-"`
// 	CreatedAt    time.Time `json:"created_at"`
// 	UpdatedAt    time.Time `json:"updated_at"`
// }

type Article struct {
	ID         int            `json:"id"`
	UserID     int            `json:"user_id"`
	Title      string         `json:"title"`
	Content    string         `json:"content"`
	CoverImage sql.NullString `json:"cover_image"`
	Category   string         `json:"category"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

// type Video struct {
// 	ID         int            `json:"id"`
// 	UserID     int            `json:"user_id"`
// 	Title      string         `json:"title"`
// 	URL        string         `json:"url"`
// 	CoverImage sql.NullString `json:"cover_image"`
// 	Category   string         `json:"category"`
// 	CreatedAt  time.Time      `json:"created_at"`
// 	UpdatedAt  time.Time      `json:"updated_at"`
// }

// type Ebook struct {
// 	ID         int            `json:"id"`
// 	UserID     int            `json:"user_id"`
// 	Title      string         `json:"title"`
// 	FileURL    string         `json:"file_url"`
// 	CoverImage sql.NullString `json:"cover_image"`
// 	Category   string         `json:"category"`
// 	CreatedAt  time.Time      `json:"created_at"`
// 	UpdatedAt  time.Time      `json:"updated_at"`
// }

type Comment struct {
	ID        int           `json:"id"`
	UserID    int           `json:"user_id"`
	Content   string        `json:"content"`
	ArticleID sql.NullInt64 `json:"article_id"`
	VideoID   sql.NullInt64 `json:"video_id"`
	EbookID   sql.NullInt64 `json:"ebook_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
