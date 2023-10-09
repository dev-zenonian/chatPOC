package repository

import (
	"common/config"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type SessionRepository interface {
	SaveClientSession(clientID string, handlerID string, handlerAddress string) error
	DeleteClientSession(clientID string, handlerID string) error
	GetCurrentSessionStatus() (map[string][]string, error)
	GetSessionWithClientID(clientID string) (map[string][]string, error)
}

type sessionRepositoryImpl struct {
	redisClient *redis.Client
	redisCfg    *config.RedisConfig
}

func NewSessionRepositoryImpl(cfg *config.RedisConfig) (SessionRepository, error) {
	repo := &sessionRepositoryImpl{
		redisCfg: cfg,
	}

	if err := repo.initConnection(); err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *sessionRepositoryImpl) initConnection() error {
	opt, err := redis.ParseURL(r.redisCfg.DSN)
	if err != nil {
		return err
	}
	r.redisClient = redis.NewClient(opt)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	pong, err := r.redisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}
	log.Printf("Pong Message::%v\n", pong)
	return nil
}

func (r *sessionRepositoryImpl) SaveClientSession(clientID string, handlerID string, handlerAddress string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	w, err := r.redisClient.HSet(ctx, clientID, handlerID, handlerAddress).Result()
	if err != nil {
		return err
	}
	if w != 1 {
		// Entry already include in redis
		fmt.Printf("w:%v get:%v\n", 1, w)
	}
	return nil
}
func (r *sessionRepositoryImpl) DeleteClientSession(clientID string, handlerID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	w, err := r.redisClient.HDel(ctx, clientID, handlerID).Result()
	if err != nil {
		return err
	}
	if w != 1 {
		// Entry not exists in redis
		fmt.Printf("w:%v get:%v\n", 1, w)
	}
	return nil
}

func (r *sessionRepositoryImpl) GetCurrentSessionStatus() (map[string][]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	clientIDs, err := r.redisClient.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}
	res := map[string][]string{}
	for _, clientID := range clientIDs {
		handlers, err := r.redisClient.HGetAll(ctx, clientID).Result()
		if err != nil {
			return nil, err
		}
		for _, handlerAddress := range handlers {
			_, ok := res[handlerAddress]
			if !ok {
				res[handlerAddress] = []string{clientID}
				continue
			}
			res[handlerAddress] = append(res[handlerAddress], clientID)
		}
	}
	return res, nil
}

func (r *sessionRepositoryImpl) GetSessionWithClientID(clientID string) (map[string][]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	w, err := r.redisClient.HGetAll(ctx, clientID).Result()
	if err != nil {
		return nil, err
	}
	res := map[string][]string{}
	for _, handlerAddress := range w {
		res[handlerAddress] = []string{clientID}
	}
	return res, nil
}
