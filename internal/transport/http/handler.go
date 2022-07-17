package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yousafshah1214/go-comments-microservice/internal/comment"
)

const (
	CommentRetriveDBError string = "Error retriving comment from database"
	CommentInsertDBError  string = "Error inserting comment to database"
)

type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (handler *Handler) SetupRoutes() error {
	fmt.Println("Setting up routes")

	handler.Router = mux.NewRouter()
	handler.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	})

	handler.Router.HandleFunc("/api/comments", handler.GetAllComments).Methods("GET")
	handler.Router.HandleFunc("/api/comments", handler.CreateComment).Methods("POST")
	handler.Router.HandleFunc("/api/comments/{id}", handler.GetComment).Methods("GET")
	// handler.Router.HandleFunc("/api/comments/{slug}", handler.GetCommentBySlug).Methods("GET")
	handler.Router.HandleFunc("/api/comments/{id}", handler.UpdateComment).Methods("PATCH")
	handler.Router.HandleFunc("/api/comments/{id}", handler.DeleteComment).Methods("DELETE")

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (handler *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := handler.Service.GetAllComments()
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, CommentRetriveDBError)
		return
	}

	fmt.Fprintf(w, "%+v", comments)
}

func (handler *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "Error parsing string to uint")
		return
	}

	comment, err := handler.Service.GetComment(uint(commentID))
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, CommentRetriveDBError)
		return
	}

	fmt.Fprintf(w, "%+v", comment)
}

// func (handler *Handler) GetCommentBySlug(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	slug := vars["slug"]

// 	comment, err := handler.Service.GetCommentBySlug(slug)
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Fprintln(w, CommentRetriveDBError)
// 		return
// 	}

// 	fmt.Fprintf(w, "%+v", comment)
// }

func (handler *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	comment := comment.Comment{
		Slug:   "I am from post request",
		Body:   "Body of comment from post request",
		Author: "Yousaf",
	}

	comment, err := handler.Service.CreateComment(comment)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, CommentInsertDBError)
		return
	}

	fmt.Fprintf(w, "%+v", comment)
}

func (handler *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, CommentRetriveDBError)
		return
	}
	newComment := comment.Comment{
		Slug: "/new",
	}

	comment, err := handler.Service.UpdateComment(uint(commentID), newComment)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, CommentRetriveDBError)
		return
	}

	fmt.Fprintf(w, "%+v", comment)
}

func (handler *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Error parsing id to uint")
		return
	}

	comment, err := handler.Service.DeleteComment(uint(commentID))
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Error deleting comment")
		return
	}

	fmt.Fprintf(w, "%+v", comment)
}
