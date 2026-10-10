package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Создаем новый встроенный роутер
	mux := http.NewServeMux()

	// Маршруты для сайта услуг и блога
	mux.HandleFunc("POST /api/v1/orders", createOrder)
	mux.HandleFunc("GET /api/v1/posts", listPosts)
	mux.HandleFunc("GET /api/v1/posts/{id}", getPost) // Извлечение ID из URL

	fmt.Println("Сервер (стандартная библиотека) на :8080")
	http.ListenAndServe(":8080", mux)
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Создать заказ")
}

func listPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Список статей")
}

func getPost(w http.ResponseWriter, r *http.Request) {
	// Получаем динамический параметр {id} встроенным методом
	id := r.PathValue("id")
	fmt.Fprintf(w, "Получение статьи с ID: %s", id)
}
