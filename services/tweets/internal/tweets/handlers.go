package tweets

import (
	"encoding/json"
	"net/http"
	"twitter-go/services/common/amqp"
	"twitter-go/services/tweets/internal/core"
)

// CreateHandler handles creating a tweet record
func CreateHandler(t *core.TweetsService) func([]byte) interface{} {
	return func(msg []byte) interface{} {

		var tweet Tweet

		if err := json.Unmarshal(msg, &tweet); err != nil {
			return amqp.RPCError{Message: err.Error(), Status: http.StatusInternalServerError}
		}

		tweet.prepareForInsert()

		repo := NewRepository(t.Cassandra)
		if err := repo.Insert(tweet); err != nil {
			return err
		}

		// TODO-16: broadcast that a tweet was created

		return tweet
	}
}

// GetAllHandler handles returning all tweets for a given username
func GetAllHandler(t *core.TweetsService) func([]byte) interface{} {
	return func(msg []byte) interface{} {
		var req GetAllUserTweets

		if err := json.Unmarshal(msg, &req); err != nil {
			return amqp.RPCError{Message: err.Error(), Status: http.StatusInternalServerError}
		}

		repo := NewRepository(t.Cassandra)
		tweets, err := repo.GetAll(req.Username)
		if err != nil {
			return err
		}

		return tweets
	}
}