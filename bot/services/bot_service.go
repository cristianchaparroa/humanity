package services

import (
	"fmt"
	"strings"
)

const (
	// UnknowIntent ....
	UnknowIntent string = "unknow_intent"

	// StockIntent ...
	StockIntent string = "stock_intent"
)

// BootMessage ...
type BootMessage struct {
	Intent     string
	Message    interface{}
	RawMessage string
}

// IBotService ...
type IBotService interface {
	GetIntent(message string) string
	Process(message string) *BootMessage
}

// BotService ...
type BotService struct {
	StockService IStockService
}

// NewBotService returns a pointer to BotService
func NewBotService() *BotService {
	ss := NewStockService()
	return &BotService{StockService: ss}
}

// GetIntent ...
func (s *BotService) GetIntent(message string) string {
	if strings.HasPrefix(strings.ToLower(message), strings.ToLower("/stock")) {
		return StockIntent
	}
	return UnknowIntent
}

// Process ...
func (s *BotService) Process(message string) *BootMessage {
	intent := s.GetIntent(message)

	if intent == StockIntent {
		ss := s.StockService
		stock := ss.Parse(message)
		res, err := ss.Quote(stock)

		if err != nil {
			return &BootMessage{Intent: StockIntent}
		}
		raw := fmt.Sprintf("%s quote is %s (open) per share", stock, res.Open)
		return &BootMessage{Intent: StockIntent, Message: res, RawMessage: raw}
	}

	return &BootMessage{Intent: UnknowIntent}
}
