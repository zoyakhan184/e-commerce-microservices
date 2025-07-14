package models

type ActivityLog struct {
	Type      string `json:"type"`      // e.g. "user"
	Message   string `json:"message"`   // e.g. "User john@example.com registered"
	Timestamp string `json:"timestamp"` // ISO string
}
