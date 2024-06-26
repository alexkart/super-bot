package openai

import (
	"github.com/radio-t/super-bot/app/bot"
	"math/rand"
	"time"
)

// LimitedMessageHistory is a limited message history for OpenAI bot
// It's using to make context answers in the chat
// This isn't thread safe structure
type LimitedMessageHistory struct {
	limit    int
	count    int
	messages []bot.Message
}

// NewLimitedMessageHistory makes a new LimitedMessageHistory with limit
func NewLimitedMessageHistory(limit int) LimitedMessageHistory {
	return LimitedMessageHistory{
		limit:    limit,
		count:    0,
		messages: make([]bot.Message, 0, limit),
	}
}

// Add adds a new message to the history
func (l *LimitedMessageHistory) Add(message bot.Message) {
	l.count++
	l.messages = append(l.messages, message)
	if len(l.messages) > l.limit {
		l.messages = l.messages[1:]
	}
}

// GetRandomMessage returns a random message from the history
func (l *LimitedMessageHistory) GetRandomMessage() *bot.Message {
	if len(l.messages) == 0 {
		return nil
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(l.messages))

	return &l.messages[randomIndex]
}
