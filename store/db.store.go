package store

import (
	"fmt"
	"log"
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

// var (
// 	host     = "rain.db.elephantsql.com"
// 	password = "k1HV8_hFYKkX9-ci6FzuHWUTRnAxQAFp"
// 	user     = "tgaryate"
// )

func NewDatabase(host string, password string, user string, dbname string) (*Database, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres : \n %v", err.Error())
	}
	log.Printf("Connected to database %v", db)
	Db := Database{
		Db: db,
	}
	return &Db, nil
}

func (db *Database) AutoMigrate() {
	log.Println("Running Migrations")
	err := db.Db.AutoMigrate(

		&models.Poll{},
		&models.Vote{},
		&models.Reference{},
		&models.PollMember{},
		&models.Moderator{},
	)
	if err != nil {
		log.Fatalf("failed to create tables")
		return
	}
	log.Println("tables created")
}

func (db *Database) CreateMember(member models.PollMember) (uint64, error) {
	err := db.Db.Save(&member).Error
	return member.Id, err
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
