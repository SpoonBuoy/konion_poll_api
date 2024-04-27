package store

import (
	"poll/models"

	"github.com/gin-gonic/gin"
)

// pollstore could be cache or db
type PollStore interface {
	GetPollById(c *gin.Context, id uint64) (models.Poll, error)
	CreateMember(member models.PollMember) (uint64, error)
	Save(poll models.Poll) (uint64, error)
	Fetch(id uint64) (models.Poll, error)
	AddVote(vote models.Vote) error
	AddReference(ref models.Reference) error
	End(id uint64) error
}
