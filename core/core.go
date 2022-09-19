package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgraph-io/badger"
	"github.com/gorilla/mux"
	"main.go/cache"
)

var db *badger.DB

type opts struct {
	Learning bool
	TTL      string
}

var Options = opts{
	Learning: false,
}

func init() {

	db, _ = badger.Open(badger.DefaultOptions("./KVDS"))

}

func cacheKey(req *http.Request) []byte {
	/*cache key structure

	GET  -> method.host.url
	POST -> method.host.url.postparameters

	*/
	key := ""
	if req.Method == "GET" {
		key += "GET"
		key += "."
		key += req.Host
		key += "."
		key += req.URL.Path

	}
	return []byte(key)

}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Host, r.URL.Path, r.URL.Scheme, r.Method)
	key := cacheKey(r)
	fmt.Println(string(key))

	if cache.IsChached(db, key) {
		fmt.Println("return from cache")
		if !Options.Learning {
			fmt.Fprintf(w, string(cache.GetItem(db, key)))
		}

	} else { // pass to upstream

		fmt.Printf("New Request: %s \n", string(key))

		r.RequestURI = ""
		resp, err := http.DefaultClient.Do(r)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		bodyContent, _ := ioutil.ReadAll(resp.Body)
		cache.Cache(db, key, bodyContent)
		fmt.Fprintf(w, string(bodyContent))
	}

}

func GetCache() string {
	cacheContent := cache.GetCachedKeys(db)
	if len(cacheContent) == 0 {
		return "empty"

	}
	j, err := json.Marshal(cacheContent)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	return string(j)
}

func Run() {
	defer db.Close()

	router := mux.NewRouter()
	router.PathPrefix("/").HandlerFunc(handleRequest)

	fmt.Println("Listening:8080 ...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
