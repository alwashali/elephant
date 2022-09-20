package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/alwashali/elephant/cache"

	"github.com/dgraph-io/badger"
	"github.com/gorilla/mux"
)

var db *badger.DB

func init() {

	badgerOpts := badger.Options{}
	badgerOpts.WithEventLogging(false)
	badgerOpts.WithDir("./KVDS")
	db, _ = badger.Open(badgerOpts)

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
	fmt.Println("cache key: ", string(key))

	if cache.IsChached(db, key) {
		fmt.Println("return from cache")
		if true {
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

	fmt.Println("checking cache")
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
