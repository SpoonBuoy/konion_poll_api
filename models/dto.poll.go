package models

type CreateMemberReq struct {
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Perception  string `json:"perception"`
}

type CreatePollReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Banner      string `json:"banner"`
	Member1     uint64 `json:"member_1"`
	Member2     uint64 `json:"member_2"`
}

type RegisterVoteReq struct {
	PollId uint64 `json:"poll_id"`
	To     uint64 `json:"to"`
}

type CreateReferenceReq struct {
	PollId uint64 `json:"poll_id"`
	Name   string `json:"name"`
	Link   string `json:"link"`
}
