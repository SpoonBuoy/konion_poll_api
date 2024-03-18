package controller

import (
	"net/http"
	"poll/models"
	"poll/service"

	"github.com/gin-gonic/gin"
)

type PollController struct {
	Service *service.PollService
}

func NewPollController(svc *service.PollService) PollController {
	return PollController{
		Service: svc,
	}
}

func (pc *PollController) CreateMember(c *gin.Context) {
	var req models.CreateMemberReq
	err := c.Bind(&req)
	HandleBindingError(c, err)
	id, err := pc.Service.CreateMember(req)
	HandleServe(c, err, id)
}

func (pc *PollController) CreatePoll(c *gin.Context) {
	var req models.CreatePollReq
	err := c.Bind(&req)
	HandleBindingError(c, err)
	id, err := pc.Service.CreatePoll(req)
	HandleServe(c, err, id)
}

func (pc *PollController) RegisterVote(c *gin.Context) {
	var req models.RegisterVoteReq
	err := c.Bind(&req)
	HandleBindingError(c, err)
	clientIp := c.ClientIP()
	err = pc.Service.RegisterVote(req, clientIp)
	HandleServe(c, err, nil)
}

func (pc *PollController) CreateReference(c *gin.Context) {
	var req models.CreateReferenceReq
	err := c.Bind(&req)
	HandleBindingError(c, err)
	err = pc.Service.CreateReference(req)
	HandleServe(c, err, nil)
}

func (pc *PollController) EndPoll(c *gin.Context) {

}

func (pc *PollController) FetchPoll(c *gin.Context) {

}

func HandleServe(c *gin.Context, err error, data any) {
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": data})
}

func HandleBindingError(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request body " + err.Error()})
		return
	}
}
