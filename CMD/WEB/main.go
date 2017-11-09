
package main 


import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"os"
	"encoding/json"
//	"github.com/heroku/Assignment2/CurrencyTicker"
//	"github.com/heroku/Assignment2/CMD/WebHooks"
	service "github.com/heroku/Assignment2/CurreencyTicker"
	)


/*

func (db *CurrencyTickerDB) Get_WebHooks(w http.ResponseWriter, db    , id string) {
	webhook, ok := db.Get(id)
	if !ok {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
	}

	json.NewEncoder(w).Encode(webhook)
}
*/



/*



func HandlerWebhook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	
	case "POST":
	//	var webhook WebHook{}
		err := json.NewDecoder(r.Body).Decode(webhook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return 
		}
		webhook := insert_webhook() // ...

		return 

	case "GET":
		http.Header.Add(w.Header(), "content-type", "application/json")

		return


	
		}
	}
*/


type SubscriberHandler struct {
	db 			subscriberDB
	Monitor 	CurrencyMonitor
}

//TODO funk to copy data into this struct or something fk? 






func (handler *SubscriberHandler) HandleSubRequestPOST(r http.ResponseWriter, w *http.Request) {
	
	err:= json.NewDecoder(w.Body).Decode(&s)

	if err != nil {
		fmt.Println("error:  %v", http.StatusBadRequest)
		return
	}

	if (!ValidateSub(s)) {
		fmt.Println("error: %v", http.StatusBadRequest)
		return
	}
	/*
	// check validity of URL in posted json
	_, err = url.ParseRequestURI(*s.WebhookURL)
	if err != nil {
		respWithCode(&res, http.StatusBadRequest)
		return
}	*/



	id, err2 := handler.db.Add_New_Ticker()
	if err2 != nil {
		fmt.Println("error:   %v", http.StatusInternalServerError)
	}

	// if all works prints to resposnewriter
	fmt.Fprint(r, id)

}


func (handler *SubscriberHandler) HandleRequestGET(r http.ResponseWriter, w *http.Request) {

	parts := strings.Split(req.URL.String(), "/")
	if len(parts) < 2 {
		fmt.Println("error:  %v", http.StatusBadRequest )
		return
	}



	sub, err :=  handler.db.Get(parts[1]) 
	if err != nil {
		fmt.Println("error:  %v", http.StatusNotFound) 
		return 
	}




	http.Header.Add(res.Header(), "content-type", "application/json")

	err := json.Encoder(r,).Encode(sub)
	if err != nil {
		fmt.Println("error:  %v", http.StatusInternalServerError)
		return
	}


}

func (handler *SubscriberHandler) HandleRequestDELETE(r http.ResponseWriter, w *http.Request) {


	parts := strings.Split(w.URL.String(), "/")
	if len(parts) < 2 {
		fmt.Println("error    %v", StatusBadRequest)
		return
	}
	err := handler.Remove_Webhook_byId(parts[1])
	if err != nil {
		fmt.Println("error     %v", StatusNotFound )
		return
	}

fmt.Fprint(r, http.StatusOK)

}

func (handler *SubscriberHandler) HandleLatest(r ResponseWriter, w *http.Request) {


		var currReq UserData

		err := json.NewDecoder(r.Body).Decode(&currReq)
		if err!= nil {
			fmt.Println("error:  	%v", StatusBadRequest)
			return
		}

	//  TODO  add more checks
	//

	rate, err := handler.Get_Last_Webhook (*UserData.BaseCurrency, *UserData.TargetCurrency)

	if err != nil {
		Fprint("error  	%v", http.StatusInternalServerError)
		return 
	}


Fprint(r, rate)

}

func (handler *SubscriberHandler) HandleAverage(r ResponseWriter, w *http.Request) {
if w.Method != "POST" {
	fmt.Println("error: only post is implemented")
	return
}

var currReq CurrencyData
err := json.NewDecoder(w.Body).Decode(%currReq)
if err != nil {
	fmt.Println("error:		%v", StatusBadRequest)
}



rate, err := handler.Get_Average()



}


	
func (handler *SubscriberHandler) HandleSubscriberRequest(r http.ResponseWriter, w *http.Request) {

	// switch on the method of the request
switch w.Method {
	case "POST": 
			handler.HandleRequestPOST(r, w)
	case "GET":
			handler.HandleRequestGET(r, w) 
	case "DELETE"
			handler.HandleRequestDELETE(r, w)

}


func main () {

/*
	Global_db := &CurrencyTicker.CurrencyTickerDB{
		"mongodb://localhost",
		"currencyTicker_db",
		"CurrencyData",
	}

	fmt.Println(Global_db.DatabaseURL)
*/

	port := os.Getenv("PORT")
	fixerIO_url :=service.GetENV("FIXER_IO_URL")
	mongodb_url := service.GetENV("MONGO_DB_URL")
	mongoDBDatabaseName := service.GetENV("MONGO_DB_DATABASE_NAME")

	db, err := service.SubscriberMongo

	//port := "localhost:8080"




	// set up handlers
	http.HandleFunc("/", HandleSubRequest)
	http.Handlefunc("/latest", HandleLatest)
	http.HandleFunc("/average", HandleAverage)
	http.HandleFunc("/triggerall", HandleTriggerAll)


	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}





	//http.ListenAndServe("localhost8080", nil)


}
