package main

import (
	"fmt"
	"net/http"

	"github.com/AanishRahmani/rssAggregator/internal/auth"
	"github.com/AanishRahmani/rssAggregator/internal/databases"
)

// Define authHandler as a function type with an additional user parameter
type authHandler func(http.ResponseWriter, *http.Request, databases.User)

// middlewareAuth wraps the handler with authentication logic
func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract API key
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithError(w, http.StatusForbidden, fmt.Sprintf("auth error: %v", err))
			return
		}

		// Get user by API key
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, http.StatusUnauthorized, fmt.Sprintf("couldn't get user: %v", err))
			return
		}

		// Call the actual handler with the authenticated user
		handler(w, r, user)
	}
}
