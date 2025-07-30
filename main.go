package main

import (
	"encoding/json"
	"fmt"
	"database/sql"
	"log"
	"net/http"
	"os"
	"github.com/justinas/alice"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)



type App struct {
	DB *sql.DB
}

type Credentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserResponse struct {
	XataID string `json:"xata_id,omitempty"`
	Username string `json:"username,omitempty"`

}	

type ErrorResponse struct {
	Message string `json:"message"`
	
}



type RouterResponse struct {
	Message string `json:"message"`
}



const contentTypeJSON = "application/json"

const projectByIDRoute = "/project/{id}"





func main() {
	router := mux.NewRouter()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	
	connStr := os.Getenv("XATA_PSQL_URL")
	if len(connStr) == 0 {
		log.Fatalf("XATA_PSQL_URL environment variable is not set")
	}



	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	defer DB.Close()


	app := &App{
		DB: DB,
	}







	
	router.Handle("/", alice.New(LoginMiddleware).ThenFunc(homeHandler)).Methods("GET")
	router.Handle("/register", alice.New(LoginMiddleware).ThenFunc(registerHandler)).Methods("POST")
	router.Handle("/login", alice.New(LoginMiddleware).ThenFunc(loginHandler)).Methods("POST")
	router.Handle("/projects", alice.New(LoginMiddleware).ThenFunc(createProjectHandler)).Methods("POST")
	router.Handle("/projects", alice.New(LoginMiddleware).ThenFunc(getProjectsHandler)).Methods("GET")
	router.Handle(projectByIDRoute, alice.New(LoginMiddleware).ThenFunc(getProjectHandler)).Methods("GET")
	router.Handle(projectByIDRoute, alice.New(LoginMiddleware).ThenFunc(updateProjectHandler)).Methods("PUT")
	router.Handle(projectByIDRoute, alice.New(LoginMiddleware).ThenFunc(deleteProjectHandler)).Methods("DELETE")

	log.Println("Server is running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}





func LoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		log.Println("#### Middleware executed ####")
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr) 
		w.Header().Set("Content-Type", contentTypeJSON) 
		next.ServeHTTP(w, r) 
		log.Println("#### Middleware completed ####") 
	})
}






func homeHandler(w http.ResponseWriter, r *http.Request) {	sendJSONResponse(w, RouterResponse{Message: "Welcome To Projects Kanban API xxx"})
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    var credentials Credentials
    if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
        w.Header().Set("Content-Type", contentTypeJSON)
        json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid request payload"})
        log.Println("Error decoding request body:", err)
        return
    }

    w.Header().Set("Content-Type", contentTypeJSON)
    json.NewEncoder(w).Encode(RouterResponse{Message: "Register Endpoint Hitx"})
    log.Println("Register endpoint hit")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, RouterResponse{Message: "Login Endpoint Hit"})
}

func createProjectHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, RouterResponse{Message: "Create Project Endpoint Hitx"})
}

func getProjectsHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, RouterResponse{Message: "Get All Projects Endpoint Hitx"})
}

func getProjectHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, RouterResponse{Message: "Get Project By ID Endpoint Hitx"})
}

func updateProjectHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, RouterResponse{Message: "Update Project By ID Endpoint Hitx"})
}

func deleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, RouterResponse{Message: "Delete Project By ID Endpoint Hitx"})
}


func sendJSONResponse(w http.ResponseWriter, response RouterResponse) {
	w.Header().Set("Content-Type", contentTypeJSON)
	json.NewEncoder(w).Encode(response)
}