package dashboard

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// DashboardHandler expose dashboard routes
type DashboardHandler struct {
	Assets http.FileSystem
}

// Append add dashboard routes on a router
func (g DashboardHandler) Append(router *mux.Router) {
	if g.Assets == nil {
		log.Printf("No assets for dashboard")
		return
	}

	// Expose dashboard
	router.Methods(http.MethodGet).
		Path("/").
		HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
			http.Redirect(response, request, request.Header.Get("X-Forwarded-Prefix")+"/dashboard/", http.StatusFound)
		})
	router.Handle("/dashboard", http.RedirectHandler("/dashboard/", http.StatusFound))
	if g.Assets != nil {
		router.Methods(http.MethodGet).
			PathPrefix("/dashboard/").
			Handler(http.StripPrefix("/dashboard/", http.FileServer(g.Assets)))
	}
}
