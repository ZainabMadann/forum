package main

import (
	// "database/sql"
	// "log"
	"log"
	"net/http"
	"os"

	"learn.reboot01.com/git/hbudalam/forum/pkg/db"
	"learn.reboot01.com/git/hbudalam/forum/pkg/server"
)

func main() {

	dbErr := db.Connect()
	if dbErr != nil {
		log.Println("Error connecting to the database:", dbErr)
		os.Exit(1)
	}
	defer db.Close()

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/login", server.LoginHandler)                        // ✅
	mux.HandleFunc("/posts/", server.PostHandler)                       // ✅
	mux.HandleFunc("/api/posts/{id}/comments", server.CommentsHandler)   // ✅
	mux.HandleFunc("/posts/{id}", server.GetPostHandler)                 // ✅
	mux.HandleFunc("/api/posts/{id}/dislike", server.AddDislikesHandler) // ✅
	mux.HandleFunc("/api/posts/{id}/like", server.AddLikesHandler)       // ✅
	mux.HandleFunc("/signup", server.SignupHandler)                      // ✅
	mux.HandleFunc("/", server.HomeHandler)                              // ✅
	mux.HandleFunc("/add-post", server.AddPostsHandler)                  // ✅
	mux.HandleFunc("/logout", server.LogoutHandler)                      // ✅
	mux.HandleFunc("/error404",server.Error404Handler)
	mux.HandleFunc("/error500",server.Error500Handler)
	mux.HandleFunc("/error400",server.Error400Handler)
	mux.HandleFunc("/api/comments/{id}/like",server.LikeCommentHandler)
	mux.HandleFunc("/api/comments/{id}/dislike",server.DislikeCommentHandler)
	mux.HandleFunc("/myPosts",server.MyPostsHandler)
	mux.HandleFunc("/filter-posts", server.FilterPostsHandler) 
	mux.HandleFunc("/most-liked", server.MostLikedPostsHandler)
	mux.HandleFunc("/newest", server.NewestPostsHandler)
	mux.HandleFunc("/Mylikedposts", server.MyLikedPostsHandler)
	
	log.Println("Serving on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println("Error starting server:", err)
		os.Exit(1)
	}
}
