package model

import "time"

type SendCodeMessage struct {
	DateCreated time.Time `json:"date_created"`
	Success     bool      `json:"success"`
}
