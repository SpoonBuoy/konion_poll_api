package service

import (
	"poll/models"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateMember(req models.CreateMemberReq) (uint64, error)
	CreatePoll(req models.CreatePollReq) (uint64, error)
	RegisterVote(req models.RegisterVoteReq, from string) error
	CreateReference(req models.CreateReferenceReq) error
	EndPoll(pollId uint64) error
	GetPoll(c *gin.Context, req models.GetPollReq) (models.Poll, error)
}
