package domain

import "time"

type Customer struct {
	Name         string    `json:"name" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	CreationTime time.Time `json:"creation_time"`
}
