package main 
	
	import (
	"net/http"
	"fmt"
	"net/url"
	"io/ioutil"
	"strings"
	"strconv"
	"github.com/heroku/Assignment2/CurrencyTicker"
	)


/*

type PayLoad struct {
	WebhookURL string `json:"text"`
	BaseCurrency string `json:"text"`
	TargetCurrency string `json:"text"`
	minTriggerValue string `json:"text"`
	maxTriggerValue float `json:"float"`

}
*/


func FloatToString(input_num float64) string {
    // to convert a float number to a string
    return strconv.FormatFloat(input_num, 'f', 2, 64)
}
 






func main () {



//userWebhook := PayLoad{}

fixerURL := "http://api.fixer.io/latest?symbols=NOK"

	resp, _ := http.Get(fixerURL)
	bytes, _ := ioutil.ReadAll(resp.Body)
	webData := string(bytes)
	
	defer resp.Body.Close()										// make sure we close body after
	
	parts := strings.Split(webData,"\"")	
	base := parts[3]
	TargetCurrency := parts[11]
	minTriggerValue := 1.5
	maxTriggerValue := 2.55

discordURL := "https://discordapp.com/api/webhooks/374908902032146432/lHx9nUAyDy1jmoahR8WrWnWl_y0B1WYc61LRuMPrcS9H5g9CcoVV3KVq7DzpfASIPPzP" // TODO edit this 

res, err2 := http.PostForm(discordURL, url.Values{"content": {"BaseCurrency  " + base + "    TargetCurrency  " + TargetCurrency + "   minTriggerValue	" + FloatToString(minTriggerValue) + "    maxTriggerValue		" + FloatToString(maxTriggerValue) } , "username": {"LYSA_BOT"}})

if err2 != nil {
	fmt.Errorf("Error doing post: %v", err2.Error())
}

if res.StatusCode != http.StatusOK {
	fmt.Errorf("Wrong status code: %v", res.StatusCode)
}

//GET_LATEST()

/*body ,err, := ioutil.ReadAll(res.Body)
if err != nil {
	fmt.Printf("Id:   ", res.Id) 
}*/

}




