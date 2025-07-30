package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/justinas/alice"

	"github.com/gorilla/mux"
)

type RouterResponse struct {
	Message string `json:"message"`
}

const contentTypeJSON = "application/json"

const projectByIDRoute = "/project/{id}"

func main() {
	router := mux.NewRouter()
	
	// Middleware to log requests and handle authentication
	// router.Use(Middleware)
	// // Middleware to handle CORS (Cross-Origin Resource Sharing)
	// router.Use(func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		w.Header().Set("Access-Control-Allow-Origin", "*")
	// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	// 		if r.Method == http.MethodOptions {
	// 			w.WriteHeader(http.StatusOK)
	// 			return
	// 		}
	// 	})
	// })



	




	// Define routes and associate them with handlers


	
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




// LOGIN MIDDLEWARE

func LoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Placeholder for authentication logic
		// In a real application, you would check the user's credentials here
		log.Println("#### Middleware executed ####")
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr) // Log the request method and path 
		w.Header().Set("Content-Type", contentTypeJSON) // Set the content type for the response
		next.ServeHTTP(w, r) // Call the next handler in the chain
		log.Println("#### Middleware completed ####") 
	})
}





// Handlers for various routes



func homeHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, RouterResponse{Message: "Welcome To Projects Kanban API MAZIN"})
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, RouterResponse{Message: "Register Endpoint Hit"})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, RouterResponse{Message: "Login Endpoint Hit"})
}

func createProjectHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, RouterResponse{Message: "Create Project Endpoint Hit"})
}

func getProjectsHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, RouterResponse{Message: "Get All Projects Endpoint Hit"})
}

func getProjectHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, RouterResponse{Message: "Get Project By ID Endpoint Hit"})
}

func updateProjectHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, RouterResponse{Message: "Update Project By ID Endpoint Hit"})
}

func deleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, RouterResponse{Message: "Delete Project By ID Endpoint Hit"})
}

// Helper function to send JSON responses
func sendJSONResponse(w http.ResponseWriter, response RouterResponse) {
	w.Header().Set("Content-Type", contentTypeJSON)
	json.NewEncoder(w).Encode(response)
}