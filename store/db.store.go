package store

import (
	"fmt"
	"poll/models"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// implements PollStore interface
type Database struct {
	Db *gorm.DB
	Mu sync.Mutex
}

func NewDatabase() (*Database, error) {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres : \n %v", err.Error())
	}
	Db := Database{
		Db: db,
	}
	return &Db, nil
}

func (db *Database) CreateMember(member models.PollMember) (uint64, error) {
	db.Db.Save(&member)
	return member.Id, nil
}
func (db *Database) Save(poll models.Poll) (uint64, error) {
	db.Db.Save(&poll)
	return poll.Id, nil
}

func (db *Database) Fetch(id uint64) (models.Poll, error) {
	var poll models.Poll
	db.Db.First(&poll)
	return poll, nil
}
func (db *Database) AddVote(vote models.Vote) error {
	//not needed to add individual votes since its captured by redis
	return nil
}
func (db *Database) AddReference(ref models.Reference) error {
	db.Db.Save(&ref)
	return nil
}
func (db *Database) End(pollId uint64) error {
	return nil
}
