package store

import "poll/models"

// pollstore could be cache or db
type PollStore interface {
	CreateMember(member models.PollMember) (uint64, error)
	Save(poll models.Poll) (uint64, error)
	Fetch(id uint64) (models.Poll, error)
	AddVote(vote models.Vote) error
	AddReference(ref models.Reference) error
	End(id uint64) error
}
