package openai

import (
	"github.com/radio-t/super-bot/app/bot"
	"math/rand"
	"time"
)

// ChatHistory holds the count and messages for a single chat
type ChatHistory struct {
	count    int
	messages []bot.Message
}

// LimitedMessageHistory is a limited message history for OpenAI bot
// It's using to make context answers in the chat
// This isn't thread safe structure
type LimitedMessageHistory struct {
	limit int
	chats map[string]*ChatHistory
}

// NewLimitedMessageHistory makes a new LimitedMessageHistory with limit
func NewLimitedMessageHistory(limit int) LimitedMessageHistory {
	return LimitedMessageHistory{
		limit: limit,
		chats: make(map[string]*ChatHistory),
	}
}

// Add adds a new message to the history for a specific chat
func (l *LimitedMessageHistory) Add(message bot.Message) {
	chatHistory, exists := l.chats[message.ChatID]
	if !exists {
		chatHistory = &ChatHistory{
			count:    0,
			messages: make([]bot.Message, 0, l.limit),
		}
		l.chats[message.ChatID] = chatHistory
	}

	chatHistory.count++
	chatHistory.messages = append(chatHistory.messages, message)
	if len(chatHistory.messages) > l.limit {
		chatHistory.messages = chatHistory.messages[1:]
	}
}

// GetRandomMessage returns a random message from the history for a specific chat
func (l *LimitedMessageHistory) GetRandomMessage(chatID string) *bot.Message {
	chatHistory, exists := l.chats[chatID]
	if !exists || len(chatHistory.messages) == 0 {
		return nil
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(chatHistory.messages))

	return &chatHistory.messages[randomIndex]
}

// GetMessagesByChatID returns the messages for a specific chat
func (l *LimitedMessageHistory) GetMessagesByChatID(chatID string) []bot.Message {
	chatHistory, exists := l.chats[chatID]
	if !exists {
		return nil
	}
	return chatHistory.messages
}
