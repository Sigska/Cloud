
package WebHookFunctions

import (
	"testing"
	//"net/http"
	//"net/http/httptest"
	//"time"
	//"log"
	"fmt"
	//"strings"
	"net/http"
//	"net/http/httptest"
//	"encoding/json"
//	"bytes"
	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	"github.com/heroku/Assignment2/WebHookFunctions"
)

var testData = WebHook( "URL_test", "EURO", "NOK", 2.4, 8.6)
var testDataDB = webHookDB("user: sdsds", "cloudtest", "webooks")

type Test struct {
	Base string `json:"base"`
	target string `json: "target"` 
}

func TestWebhookFunctions_add(t* testing.T) {
	statusCheck := testDataDB.Insert_Webhook(testData)
	if statusCheck == 0 {
		t.Error("Adding failed")
	}

}



func Test_WebhookFunctions_remove(t *testing.T) {
	statusCheck := testDataDB.Remove_Webhook_byId(testData.ID)
	if statusCheck == 0 {
		t.Error("Deleting failed")
	}
}


func test_WebHookFunctions_Invoke_Webhooks(t* testing.T) {

testDatas := make([]WebHook, 0, 10)
session, err := mgo.Dial(testDataDB.DatabaseURL)
if err != nil {
	t.Error("error dialing")
}
	defer session.Close()

dbSize, err := session.DB(db.DatabaseName).C("webhooks").Count()
if err != nil {
     t.Error("error counting")
}



err = session.DB(db.DatabaseName).C("webhooks").Find(nil).All(&testDatas)
if err != nil {
    t.Error("error getting / find all webhooks")
}

	for i := 0; i < dbSize; i++ {
		_, err := http.PostForm(testDatas[i].WebhookURL, url.Values{"content": {"Webhook ID: " + testDatas[i].ID.Hex() + "		BaseCurrency  " + testDatas[i].Base + "    TargetCurrency  " + testDatas[i].Target + "   minTriggerValue	" + FloatToString(testDatas[i].Min) + "    maxTriggerValue		" + FloatToString(testDatas[i].Max) } , "username": {"IAMBOT"}})
		if err != nil {
			t.Error("error posting all webhooks")
		}
	}
}



func test_WebHookFunctions_Get_Last_Webhook(t* testing.T) {
	session, err := mgo.Dial(db.DatabaseURL)
if err != nil {
	t.Error("error creating session")
}
	defer session.Close()


dbSize, err := session.DB(db.DatabaseName).C("webhooks").Count()
if err != nil {
     fmt.Println("error counting collection :( :", err.Error())
	}

		err = session.DB(db.DatabaseName).C("webhooks").Find(nil).Skip(dbSize-1).One(&testData)
		if err != nil {
		    t.Error("Error finding last webhook")
		}

	// post to database
	res, err := http.PostForm(testData.WebhookURL, url.Values{"content": {"Webhook ID: " + testData.ID.Hex() + "		BaseCurrency  " + testData.Base + "    TargetCurrency  " + testData.Target + "   minTriggerValue	" + FloatToString(testData.Min) + "    maxTriggerValue		" + FloatToString(testData.Max) } , "username": {"IAMBOT"}})
		if err != nil {
			t.Error("error doing post")
		}

	if res.StatusCode != http.StatusOK {
		t.Error("wrong statuscode")
	}

}

