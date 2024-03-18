package models

import "time"

type Poll struct {
	Id          uint64       `gorm:"id;primary_key"`
	IsActive    bool         `gorm:"is_active"`
	Title       string       `gorm:"title"`
	Description string       `gorm:"description"`
	Banner      string       `gorm:"banner"`
	Members     []PollMember `gorm:"members"`
	Winner      PollMember   `gorm:"winner"`
	Mod         Moderator    `gorm:"mod"`
	CreatedAt   time.Time    `gorm:"created_at"`
	References  []Reference  `gorm:"references"`
}

type Vote struct {
	Id           uint64    `gorm:"id;primary_key"`
	PollId       uint64    `gorm:"poll_id"`
	FromIP       string    `gorm:"from_ip"`
	FromLocation string    `gorm:"from_location"`
	RegisteredAt time.Time `gorm:"registered_at"`
	To           uint64    `gorm:"to"`
	IsOrgainc    bool      `gorm:"is_organic"`
}

type Reference struct {
	Id        uint64 `gorm:"id;primary_key"`
	PollId    uint64 `gorm:"poll_id"`
	Name      string `gorm:"name"`
	Link      string `gorm:"link"`
	ViewCount uint32 `gorm:"view_count"`
}

type PollMember struct {
	Id               uint64 `gorm:"id;primary_key"`
	Name             string `gorm:"name"`
	Avatar           string `gorm:"avatar"`
	Description      string `gorm:"description"`
	Perception       string `gorm:"perception"`
	SeedVoteCount    uint32 `gorm:"seed_vote_count"`
	OrganicVoteCount uint32 `gorm:"organic_vote_count"`
	TotalVotes       uint32 `gorm:"total_votes"`
}

type Moderator struct {
	Id         uint64 `gorm:"id;primary_key"`
	Name       string `gorm:"name"`
	Email      string `gorm:"email"`
	Phone      string `gorm:"phone"`
	Password   string `gorm:"password"`
	TotalPolls uint32 `gorm:"total_polls"`
}
