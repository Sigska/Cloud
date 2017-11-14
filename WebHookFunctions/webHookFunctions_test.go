
package WebHookFunctions

import (
	"testing"
	//"time"
	//"log"
	"fmt"
//	"strings"
	"net/http"
	"net/url"
//	"net/http/httptest"
//	"encoding/json"
//	"bytes"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/heroku/Assignment2/WebHookFunctions"
	"github.com/heroku/Assignment2/utils"
	"github.com/heroku/Assignment2/CurrencyTicker"
)

//
//db := CurrencyTicker.CurrencyTickerDB{"mongodb://Siggy:Siggy@ds145275.mlab.com:45275/currency_db", "currency_db", "currency"}

var test_id = bson.NewObjectId()

var testData = WebHookFunctions.WebHook{ test_id, "URL_test", "EURO", "NOK", 2.4, 8.6}
var testDataDB = CurrencyTicker.CurrencyTickerDB{"user: sdsds", "cloudtest", "webooks"}

type Test struct {
	Base string `json:"base"`
	target string `json: "target"` 
}

func TestWebhookFunctions_add(t* testing.T) {
	err := WebHookFunctions.Insert_Webhook(&testDataDB)
	if err !=  nil {
		t.Error("Adding failed")
	}

}



func Test_WebhookFunctions_remove(t *testing.T) {
	Id := string(testData.ID)
	err := WebHookFunctions.Remove_Webhook_byId(&testDataDB, Id)
	if err !=  nil {
		t.Error("Deleting failed")
	}
}


func test_WebHookFunctions_Invoke_Webhooks(t* testing.T) {

testDatas := make([](WebHookFunctions.WebHook), 0, 10)
session, err := mgo.Dial(testDataDB.DatabaseURL)
if err != nil {
	t.Error("error dialing")
}
defer session.Close()

dbSize, err := session.DB(testDataDB.DatabaseName).C("webhooks").Count()
if err != nil {
     t.Error("error counting")
}



err = session.DB(testDataDB.DatabaseName).C("webhooks").Find(nil).All(&testDatas)
if err != nil {
    t.Error("error getting / find all webhooks")
}

	for i := 0; i < dbSize; i++ {
		_, err := http.PostForm(testDatas[i].WebhookURL, url.Values{"content": {"Webhook ID: " + testDatas[i].ID.Hex() + "		BaseCurrency  " + testDatas[i].Base + "    TargetCurrency  " + testDatas[i].Target + "   minTriggerValue	" + utils.FloatToString(testDatas[i].Min) + "    maxTriggerValue		" + utils.FloatToString(testDatas[i].Max) } , "username": {"IAMBOT"}})
		if err != nil {
			t.Error("error posting all webhooks")
		}
	}
}



func test_WebHookFunctions_Get_Last_Webhook(t* testing.T) {
	session, err := mgo.Dial(testDataDB.DatabaseURL)
if err != nil {
	t.Error("error creating session")
}
	defer session.Close()


dbSize, err := session.DB(testDataDB.DatabaseName).C("webhooks").Count()
if err != nil {
     fmt.Println("error counting collection :( :", err.Error())
	}

		err = session.DB(testDataDB.DatabaseName).C("webhooks").Find(nil).Skip(dbSize-1).One(&testData)
		if err != nil {
		    t.Error("Error finding last webhook")
		}

	// post to database
	res, err := http.PostForm(testData.WebhookURL, url.Values{"content": {"Webhook ID: " + testData.ID.Hex() + "		BaseCurrency  " + testData.Base + "    TargetCurrency  " + testData.Target + "   minTriggerValue	" + utils.FloatToString(testData.Min) + "    maxTriggerValue		" + utils.FloatToString(testData.Max) } , "username": {"IAMBOT"}})
		if err != nil {
			t.Error("error doing post")
		}

	if res.StatusCode != http.StatusOK {
		t.Error("wrong statuscode")
	}

}

