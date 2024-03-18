package store

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"poll/models"
	"strconv"
	"sync"

	"github.com/go-redis/redis/v8"
)

// implements PollStore interface
type Cache struct {
	Client *redis.Client
	Mu     sync.Mutex
}

var (
	POLL = "POLL"
	VOTE = "VOTE"
	REF  = "REF"
)

func NewCache() (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})
	c := Cache{
		Client: client,
	}
	ctx := context.Background()
	pong, err := c.Client.Ping(ctx).Result()
	log.Println("Connected to Redis", pong)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
func (c *Cache) CreateMember(member models.PollMember) (uint64, error) {
	return 0, nil
}
func (c *Cache) Save(poll models.Poll) (uint64, error) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	pollId := UniqueKey(POLL, 0, poll.Id)
	json, err := json.Marshal(poll)
	if err != nil {
		return 0, fmt.Errorf("failed to marshall poll : \n %v", err.Error())
	}
	err = c.Client.Set(context.Background(), pollId, json, 0).Err()
	if err != nil {
		return 0, fmt.Errorf("failed to save poll to redis : \n %v", err.Error())
	}
	return poll.Id, nil
}

func (c *Cache) Fetch(id uint64) (models.Poll, error) {
	pollId := UniqueKey(POLL, 0, id)
	pollByte, err := c.Client.Get(context.Background(), pollId).Result()

	if err != nil {
		return models.Poll{}, fmt.Errorf("failed to fetch poll from redis :%v", err.Error())
	}
	var poll models.Poll
	err = json.Unmarshal([]byte(pollByte), &poll)
	if err != nil {
		return models.Poll{}, fmt.Errorf("failed to unmarshall poll from redis : %v", err)
	}
	return poll, nil
}
func (c *Cache) AddVote(vote models.Vote) error {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	//adds vote
	//will generate a unique key based on pollId and memberId to whom vote was given
	key := UniqueKey(VOTE, vote.PollId, vote.To)
	//will check if the key exists
	val, err := c.Client.Get(context.Background(), key).Result()
	if err != nil {
		//key does not exist
		err = c.Client.Set(context.Background(), key, 1, 0).Err()
		if err != nil {
			return fmt.Errorf("failed to add vote : \n %v", err.Error())
		}
	}
	//key exists
	voteCount, err := strconv.Atoi(val)
	if err != nil {
		return fmt.Errorf("fialed to convert votes to int : \n %v", err.Error())
	}
	//increment voteCnt
	voteCount++
	newVal := strconv.Itoa(voteCount)
	err = c.Client.Set(context.Background(), key, newVal, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to add vote : \n %v", err.Error())
	}
	return nil
}

func (c *Cache) AddReference(ref models.Reference) error {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	//adding reference will only increment the view count
	key := UniqueKey(REF, ref.Id, ref.PollId)
	val, err := c.Client.Get(context.Background(), key).Result()
	if err != nil {
		//ref does not exist
		err := c.Client.Set(context.Background(), key, 1, 0).Err()
		if err != nil {
			return fmt.Errorf("failed to add ref to redis : \n %v", err.Error())
		}
	}
	//ref exists
	viewCount, err := strconv.Atoi(val)
	if err != nil {
		return fmt.Errorf("failed to get view count int from redis : \n %v", err.Error())
	}
	viewCount++
	err = c.Client.Set(context.Background(), key, viewCount, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to add view count on redis : \n %v", err.Error())
	}
	return nil
}
func (c *Cache) GetViewCount(ref models.Reference) (uint32, error) {
	key := UniqueKey(REF, ref.Id, ref.PollId)
	val, err := c.Client.Get(context.Background(), key).Result()

	if err != nil {
		//does not exist means 0
		return 0, nil
	}
	count, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf("failed to get view count int from redis : \n %v", err.Error())
	}
	return uint32(count), nil
}

func (c *Cache) End(pollId uint64) error {
	//have to end the poll and clear all its related cache
	return nil
}

func UniqueKey(str string, x uint64, y uint64) string {
	strx := strconv.Itoa(int(x))
	stry := strconv.Itoa(int(y))
	//will have to hash it as well
	return strx + str + stry
}
