package services

import "testing"

// StockServiceMock is used to mock a process stock intent
type StockServiceMock struct {
}

// Parse retrieves the stock symbol to mock
func (sm *StockServiceMock) Parse(message string) string {
	return "AMD"
}

// ReadCSV mocks the readcsv behavior
func (sm *StockServiceMock) ReadCSV(url string) ([]*StockResponse, error) {
	res := make([]*StockResponse, 0)

	header := &StockResponse{}
	result := &StockResponse{
		Symbol: "AMD",
		Open:   "20",
	}
	res = append(res, header)
	res = append(res, result)
	return res, nil
}

// Quote retrieves the mock for Stock response pointer.
func (sm *StockServiceMock) Quote(stock string) (*StockResponse, error) {
	rows, _ := sm.ReadCSV("")
	return rows[1], nil
}

func TestNewBotService(t *testing.T) {
	bs := NewBotService()

	if bs == nil {
		t.Errorf("Expected a pointer to BotService, but get nil")
	}
}

func TestBotServiceGetIntet(t *testing.T) {

	var test = []struct {
		Message        string
		ExpectedIntent string
	}{
		{"/stock=APP", StockIntent},
		{"/sck", UnknowIntent},
		{"This is normal message", UnknowIntent},
	}

	bs := NewBotService()

	for _, tc := range test {
		result := bs.GetIntent(tc.Message)

		if result != tc.ExpectedIntent {
			t.Errorf("Expected %v, but get:%v", tc.ExpectedIntent, result)
		}
	}
}

// Test a bost service when intent is stock
func TestBotServiceProcessStockIntent(t *testing.T) {
	ss := &StockServiceMock{}
	bs := NewBotService()
	bs.StockService = ss

	message := bs.Process("/stock=AMD")

	if message == nil {
		t.Errorf("Expected a message but get nil")
	}
}
