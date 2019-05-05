package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// StockResponse represents the
// information returned from the external api
type StockResponse struct {
	Symbol string
	Date   string
	Time   string
	Open   string
	High   string
	Low    string
	Close  string
	Volume string
}

// IStockService ...
type IStockService interface {
	Parse(message string) string
	ReadCSV(url string) ([]*StockResponse, error)
	Quote(stock string) (*StockResponse, error)
}

// StockService ...
type StockService struct {
	Client *http.Client
}

// NewStockService retrieves a pointer to StockService
func NewStockService() *StockService {

	c := http.DefaultClient
	return &StockService{Client: c}
}

// Parse retrieves the stock symbol related to stock command
func (s *StockService) Parse(message string) string {

	command := strings.ReplaceAll(message, " ", "")
	results := strings.Split(command, "=")
	stock := results[1]
	return stock
}

// ReadCSV from url
func (s *StockService) ReadCSV(url string) ([]*StockResponse, error) {

	fmt.Printf("--> StockService:ReadCSV url:%s\n", url)

	resp, err := s.Client.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)

	data := make([]*StockResponse, 0)
	for {
		record, err := reader.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		symbol := record[0]
		date := record[1]
		time := record[2]
		open := record[3]
		high := record[4]
		low := record[5]
		close := record[6]
		volume := record[7]

		stock := &StockResponse{
			Symbol: symbol,
			Date:   date,
			Time:   time,
			Open:   open,
			High:   high,
			Low:    low,
			Close:  close,
			Volume: volume,
		}
		data = append(data, stock)
	}

	fmt.Printf("<-- StockService:ReadCSV\n")

	return data, nil
}

// Quote consume the external api and retrieve the values
// related to specific stock
func (s *StockService) Quote(stock string) (*StockResponse, error) {

	fmt.Printf("--> StockService:Quote:%s\n", stock)

	url := fmt.Sprintf("https://stooq.com/q/l/?s=%s.us&f=sd2t2ohlcv&h&e=csv.", stock)

	rows, err := s.ReadCSV(url)

	if err != nil {
		fmt.Printf("<-- StockService:Quote Error:%v\n", err)
		return nil, err
	}

	fmt.Printf("<-- StockService\n")
	return rows[1], nil
}
