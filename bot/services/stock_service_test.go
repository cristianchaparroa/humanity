package services

import (
	"net/http"
	"testing"
)

type HttpClientMock struct {
}

func (c *HttpClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{}, nil
}

func TestNewStockService(t *testing.T) {
	ss := NewStockService()

	if ss == nil {
		t.Errorf("The Stock service is null, expected a pointer ")
	}

	if ss.Client == nil {
		t.Errorf("The http client of Stock service is null, but expected pointer")
	}
}

func TestStockServiceParse(t *testing.T) {
	var test = []struct {
		Message  string
		Expected string
	}{
		{"/stock=AMD", "AMD"},
		{"/ stock = AMD", "AMD"},
		{"/ stock= AMD", "AMD"},
	}

	ss := &StockService{}

	for _, tc := range test {
		result := ss.Parse(tc.Message)

		if result != tc.Expected {
			t.Errorf("Expected %#v, but get:%#v", tc.Expected, result)
		}
	}
}
