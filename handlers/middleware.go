package handlers

import (
	"log"
	"net/http"
)

// Middleware для проверки авторизации пользователя
// Проверяется заголовок запроса Authorization
// Если он присутствует и соответствует заданному ключу, то произойдёт смена статуса платежа
// В противном случае будет выдана ошибка и ничего не произойдёт
func (t *Transactions) MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header["Authorization"] // получение заголовка
		t.l.Println(h)
		if h[0] != token { // проверка на соответствие заданному
			log.Println("Invalid authorization header received")
			return
		}
		t.l.Println("Token allowed, go to next handler")
		next.ServeHTTP(w, r) // переход к следующей функции
	})
}
