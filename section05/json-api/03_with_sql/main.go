package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var dataFile = path.Join("..", "data", "proverbs.json")

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	defer db.Close()

	h := newHandler(db)
	r := newRouter(h)

	var sigChan = make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGTERM)
	signal.Notify(sigChan, syscall.SIGINT)
	go func() {
		log.Printf("Signal received: %+v.", <-sigChan)
		// Cleanup
		db.Close()
		log.Println("Bye.")
		os.Exit(0)
	}()

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}

// newRouter returns a router to expose the following endpoints
// POST /proverbs (create proverb)
// GET /proverbs (get all proverbs)
// GET /proverbs/{id} (get a specific proverb)
// PUT /proverbs/{id} (update a specific proverb)
// DELETE /proverbs/{id} (delete a specific proverb)
func newRouter(h *handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/proverbs", h.createProverb).Methods("POST")
	r.HandleFunc("/proverbs", h.getProverbs).Methods("GET")
	r.HandleFunc("/proverbs/{id:[0-9]+}", h.getProverb).Methods("GET")
	r.HandleFunc("/proverbs/{id:[0-9]+}", h.updateProverb).Methods("PUT")
	r.HandleFunc("/proverbs/{id:[0-9]+}", h.deleteProverb).Methods("DELETE")
	return r
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@/gostdessentials?charset=utf8")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
