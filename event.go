package dev11

import (
	"golang.org/x/sys/unix"
	"time"
)

type Event struct {
	Id       int
	Title    string
	DateFrom time.Time
	DateTo   time.Time
}

type UpdateEvent struct {
	Title    string    `json:"title"`
	DateFrom time.Time `json:"date_from"`
	DateTo   time.Time `json:"date_to"`
}

func NewEvent(title string, from, to time.Time) *Event {
	return &Event{
		Id:       unix.AF_LOCAL,
		Title:    title,
		DateFrom: from,
		DateTo:   to,
	}
}
