package service

import (
	"encoding/json"
	"fmt"
	"model"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

//定义一个路由对象
type Route struct {
	Name        string      		  // 路由名称
	Method      string				  // 路由方法
	Pattern     string				  // 路由参数
	HandlerFunc http.HandlerFunc      // 路由函数
}

type Routes []Route

// 初始化一个路由线路
var routes = Routes{
	Route{
		"GetAccount",                // Name
		"GET",                     // HTTP method
		"/accounts/{accountId}",   // Route pattern
		func (w http.ResponseWriter, r *http.Request) {
			// Read the 'accountId' path parameter from the mux map
			var accountId = mux.Vars(r)["accountId"]

			fmt.Println("======", accountId)
			// Read the account struct BoltDB
			client  := new(model.BoltClient)
			client.OpenBoltDb()
			account, err := client.QueryAccount(accountId)

			// If err, return a 404
			if err != nil {
				fmt.Println("|||||||||||||", err)
				w.WriteHeader(http.StatusNotFound)
				return
			}

			// If found, marshal into JSON, write headers and content
			data, _ := json.Marshal(account)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Length", strconv.Itoa(len(data)))
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		},
	},
	Route{
		"index",
		"GET",
		"/user/{userId}",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write([]byte("{\"result\":\"OK\"}"))
		},
	},
}



// Function that returns a pointer to a mux.Router we can use as a handler.
func NewRouter() *mux.Router {

	// Create an instance of the Gorilla router
	router := mux.NewRouter().StrictSlash(true)

	// Iterate over the routes we declared in routes.go and attach them to the router instance
	for _, route := range routes {

		// Attach each route, uses a Builder-like pattern to set each route up.
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

