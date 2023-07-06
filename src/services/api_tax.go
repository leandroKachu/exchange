package services

import (
	"conversion-currency/src/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTaxCoin(from, to string) (float64, error) {
	var data map[string]interface{}

	url := fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=%s&api_key=%s", from, to, config.ApiKey)

	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal([]byte(body), &data); err != nil {
		return 0, err
	}
	btcValue, _ := data[to].(float64)
	// rateData, _ := data["Realtime Currency Exchange Rate"].(map[string]interface{})
	// bidPrice, _ := rateData["8. Bid Price"].(string)

	// float, err := strconv.ParseFloat(btcValue, 64)
	return btcValue, nil
}
