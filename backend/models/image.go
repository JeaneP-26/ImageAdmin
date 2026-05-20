package models

import "time"

type Image struct {
    ID         int
    UserID     int
    Name       string
    FilePath   string
    CategoryID int
    UploadedAt time.Time
}