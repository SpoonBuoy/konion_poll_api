package service

import "poll/models"

type Service interface {
	CreateMember(req models.CreateMemberReq) (uint64, error)
	CreatePoll(req models.CreatePollReq) (uint64, error)
	RegisterVote(req models.RegisterVoteReq, from string) error
	CreateReference(req models.CreateReferenceReq) error
	EndPoll(pollId uint64) error
	FetchPoll(pollId uint64) (*models.Poll, error)
}
