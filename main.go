package main

import (
	"fmt"
	"html/template"
	"net/http"
	"models"
)

var posts map[string]*models.Post // хранение в памяти

func indexHandler(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html") // возвращает template и ошибку
		if err != nil {
			fmt.Fprintf(w, err.Error()) // вывод в браузер
		}

		t.ExecuteTemplate(w, "index", nil)


}

func writeHandler(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html") // возвращает template и ошибку
		if err != nil {
			fmt.Fprintf(w, err.Error()) // вывод в браузер
		}

		t.ExecuteTemplate(w, "write", nil)
}


func savePostHandler(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")
		title := r.FormValue("title")
		content := r.FormValue("content")

		post := models.newPost(id, title, content)
		posts[post.Id] = post
}


func main(){
	fmt.Println("Listen on port:3000")

	posts = make(map[string]*models.Post, 0)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))) // правильный путь до assets 
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/SavePost", savePostHandler)

	http.ListenAndServe(":3000", nil)
}