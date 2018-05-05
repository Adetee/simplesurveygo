package main

import (
	"net/http"
	"simplesurveygo/servicehandlers"
)

func maxClients(h http.Handler, n int) http.Handler {
     sema := make(chan struct{}, n)

     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
         sema <- struct{}{}
         defer func() { <-sema }()

         h.ServeHTTP(w, r)
     })
}

func main() {

	// Serves the html pages
	http.Handle("/", http.FileServer(http.Dir("./static")))

	pingHandler := servicehandlers.PingHandler{}
	authHandler := servicehandlers.UserValidationHandler{}
	sessionHandler := servicehandlers.SessionHandler{}
	surveyHandler := servicehandlers.SurveyHandler{}
	userSurveyHandler := servicehandlers.UserSurveyHandler{}
	signupHandler := servicehandlers.SignupHandler{}

	// Serves the API content
	http.Handle("/api/v1/ping/", pingHandler)

	http.Handle("/api/v1/signup/", signupHandler)
	http.Handle("/api/v1/authenticate/", authHandler)
	http.Handle("/api/v1/validate/", sessionHandler)

	http.Handle("/api/v1/survey/{surveyname}", surveyHandler)
	http.Handle("/api/v1/survey/", surveyHandler)

	http.Handle("/api/v1/usersurvey/", userSurveyHandler)
		
	http.Handle("/",maxClients(handler,10))
	// Start Server
	http.ListenAndServe(":3000", nil)
}
