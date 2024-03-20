package main

import (
	"log"
	"net/http"
	"time"
	"time-clock/api"

	_ "github.com/lib/pq"
)

//	@title			Time Clock API
//	@version		1.0
//	@description	Here you will find everything you need to have the best possible integration with our APIs.
//	@termsOfService	http://timeclock.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.timeclock.io/support
//	@contact.email	support@timeclock.io

// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	db := api.SetupDB()
	r := api.SetupRouter(db)

	server := &http.Server{
		Addr:              ":8000",
		ReadHeaderTimeout: 3 & time.Second,
		Handler:           r,
	}
	log.Println(server.ListenAndServe())
}
