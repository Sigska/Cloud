
package main 


import (
	"fmt"
	"net/http"
	//"strings"
	"os"
//	"encoding/json"
	"github.com/heroku/Assignment2/CurrencyTicker"
//	"github.com/heroku/Assignment2/CMD/WebHooks"
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









func main () {


	Global_db := &CurrencyTicker.CurrencyTickerDB{
		"mongodb://localhost",
		"currencyTicker_db",
		"CurrencyData",
	}

	fmt.Println(Global_db.DatabaseURL)


	port := os.Getenv("PORT")

	//port := "localhost:8080"


	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}





	//http.ListenAndServe("localhost8080", nil)


}
