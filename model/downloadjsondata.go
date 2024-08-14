package model

import "github.com/google/uuid"

type DownloadJsonData struct {
	Id       uuid.UUID
	TeamName string
}
