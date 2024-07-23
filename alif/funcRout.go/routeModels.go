package funcroutgo

type ResCoin map[string]map[string]any


type ResTokensInfo struct {
	ID                           string  `json:"id"`
	Symbol                       string  `json:"symbol"`
	Name                         string  `json:"name"`
	Image                        string  `json:"image"`
	CurrentPrice                 float32 `json:"current_price"`
	MarketCap                    float32 `json:"market_cap"`
	MarketCapRank                float32 `json:"market_cap_rank"`
	FullyDilutedValuation        float32 `json:"fully_diluted_valuation"`
	TotalVolume                  float32 `json:"total_volume"`
	High24h                      float32 `json:"high_24h"`
	Low24h                       float32 `json:"low_24h"`
	PriceChange24h               float32 `json:"price_change_24h"`
	PriceChangePercentage24h     float32 `json:"price_change_percentage_24h"`
	MarketCapChange24h           float32 `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h float32 `json:"market_cap_change_percentage_24h"`
	CirculatingSupply            float32 `json:"circulating_supply"`
	TotalSupply                  float32 `json:"total_supply"`
	MaxSupply                    float32 `json:"max_supply"`
	Ath                          float32 `json:"ath"`
	AthChangePercentage          float32 `json:"ath_change_percentage"`
	AthDate                      string  `json:"ath_date"`
	Atl                          float32 `json:"atl"`
	AtlChangePercentage          float32 `json:"atl_change_percentage"`
	Atl_Date                     string  `json:"atl_date"`
	Roi                          any     `json:"roi"`
	LastUpdated                  string  `json:"last_updated"`
}
