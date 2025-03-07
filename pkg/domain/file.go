package domain

import "time"

type File struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    UserID    uint      `json:"user_id"`
    FileName  string    `json:"file_name"`
    FileSize  int64     `json:"file_size"`
    FileType  string    `json:"file_type"`
    FileURL   string    `json:"file_url"`
    CreatedAt time.Time `json:"created_at"`
}
