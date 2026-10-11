package main

import (
	"app/internal/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := routes.New()

	// Маршруты для сайта услуг и блога
	// mux.HandleFunc("GET /api/v1/posts", listPosts)
	// mux.HandleFunc("GET /api/v1/posts/{id}", getPost) // Извлечение ID из URL

	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}

// func listPosts(w http.ResponseWriter, r *http.Request) {
// 	fmt.Print("Список статей")
// }

// func getPost(w http.ResponseWriter, r *http.Request) {
// 	// Получаем динамический параметр {id} встроенным методом
// 	id := r.PathValue("id")
// 	fmt.Fprintf(w, "Получение статьи с ID: %s", id)
// }
