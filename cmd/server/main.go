package main

import (
	"log"
	"net/http"

	"example.com/pz3-http/internal/api"
	"example.com/pz3-http/internal/storage"
)

func main() {
	store := storage.NewMemoryStore()
	h := api.NewHandlers(store)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		api.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	mux.HandleFunc("GET /tasks", h.ListTasks)
	mux.HandleFunc("POST /tasks", h.CreateTask)
	mux.HandleFunc("GET /tasks/", h.GetTask)
	mux.HandleFunc("PATCH /tasks/", h.PatchTask)
	mux.HandleFunc("DELETE /tasks/", h.DeleteTask)

	handler := api.Logging(mux)

	addr := ":8080"
	log.Println("listening on", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
