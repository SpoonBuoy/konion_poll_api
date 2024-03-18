package service

import (
	"fmt"
	"poll/models"
	"poll/store"
	"time"
)

type PollService struct {
	Store store.PollStore
}

func NewPollService(store store.PollStore) (*PollService, error) {
	return &PollService{Store: store}, nil
}
func (ps *PollService) CreateMember(req models.CreateMemberReq) (uint64, error) {
	member := models.PollMember{
		Name:             req.Name,
		Avatar:           req.Avatar,
		Description:      req.Description,
		Perception:       req.Description,
		SeedVoteCount:    0,
		OrganicVoteCount: 0,
		TotalVotes:       0,
	}
	id, err := ps.Store.CreateMember(member)
	if err != nil {
		return 0, fmt.Errorf("failed to create poll member \n : %v", err.Error())
	}
	return id, nil
}
func (ps *PollService) CreatePoll(req models.CreatePollReq) (uint64, error) {
	//have to create an admin mod
	poll := models.Poll{
		IsActive:    true,
		Title:       req.Title,
		Description: req.Description,
		Banner:      req.Description,
		Members:     []models.PollMember{},
		CreatedAt:   time.Now(),
		References:  []models.Reference{},
	}
	id, err := ps.Store.Save(poll)
	return id, nil
}
func (ps *PollService) RegisterVote(req models.RegisterVoteReq, from string) error {
	vote := models.Vote{
		PollId:       req.PollId,
		To:           req.To,
		RegisteredAt: time.Now(),
		IsOrgainc:    true,
		FromIP:       from,
	}
	err := ps.Store.AddVote(vote)
	return nil
}
func (ps *PollService) CreateReference(req models.CreateReferenceReq) error {
	ref := models.Reference{
		PollId:    req.PollId,
		Name:      req.Name,
		Link:      req.Link,
		ViewCount: 0,
	}
	err := ps.Store.AddReference(ref)
	return nil
}
func (ps *PollService) EndPoll(pollId uint64) error {

	return nil
}
func (ps *PollService) FetchPoll(pollId uint64) (*models.Poll, error) {

	return nil, nil
}
