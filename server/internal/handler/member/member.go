package member

import (
	"time"

	"github.com/google/uuid"
)

type MemberRole string

const (
	Participant MemberRole    = "participant"
	Admin       MemberRole    = "admin"
	Moderator   MemberRole    = "moderator"
)

type Member struct {
	UserId     uuid.UUID  `json:"userId" db:"user_id"`
	Username   string     `json:"username" db:"username"`
	Role       MemberRole `json:"role" db:"role"`
	ForumId    uuid.UUID  `json:"forumId" db:"forum_id"`
	ForumTitle string     `json:"title" db:"title"`
	Active     bool       `json:"active" db:"active"`
	JoinedAt   time.Time  `json:"joinedAt" db:"joined_at"`
}