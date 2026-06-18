package model

import "time"

const (
	TodoStateOpen     = "open"
	TodoStateClosed   = "closed"
	TodoStateReopened = "reopened"
	TodoStateRejected = "rejected"
)

const (
	ScheduleKindUpcoming = "upcoming"
	ScheduleKindOverdue  = "overdue"
)

const (
	ScheduleStatusActive = "active"
	ScheduleStatusSent   = "sent"
)

type Todo struct {
	ID         int64      `json:"id"`
	Title      string     `json:"title"`
	DueAt      *time.Time `json:"due_at,omitempty"`
	State      string     `json:"state"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	ClosedAt   *time.Time `json:"closed_at,omitempty"`
	RejectedAt *time.Time `json:"rejected_at,omitempty"`
}

type Sink struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	Events    []string  `json:"events"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Schedule struct {
	ID         string    `json:"id"`
	TodoID     string    `json:"todo_id"`
	Kind       string    `json:"kind"`
	Before     string    `json:"before,omitempty"`
	Every      string    `json:"every,omitempty"`
	Status     string    `json:"status"`
	TargetMOTD bool      `json:"target_motd"`
	SinkIDs    []string  `json:"sink_ids,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}
