package api

type Details struct {
	Active         bool
	CurrencyName   string
	CurrencySymbol string
	// BaseCurrencyName          string
	// BaseCurrencySymbol        string
	Description               string
	HomepageURL               string
	LastUpdatedUTC            Time
	MarketCap                 float64
	Name                      string
	Ticker                    string
	WeightedSharesOutstanding int64
}
