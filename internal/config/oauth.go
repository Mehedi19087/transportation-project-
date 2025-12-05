package config

import (
    "fmt"
    "log"
    "os"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

var GoogleOAuthConfig *oauth2.Config

func InitGoogleOAuthConfig() error {
    clientID := os.Getenv("GOOGLE_CLIENT_ID")
    clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
    redirectURL := os.Getenv("GOOGLE_REDIRECT_URL")

    // Debug logs
    log.Printf("GOOGLE_CLIENT_ID: %s", clientID)
    log.Printf("GOOGLE_CLIENT_SECRET: %s", clientSecret)
    log.Printf("GOOGLE_REDIRECT_URL: %s", redirectURL)

    if clientID == "" {
        return fmt.Errorf("GOOGLE_CLIENT_ID is empty")
    }
    if clientSecret == "" {
        return fmt.Errorf("GOOGLE_CLIENT_SECRET is empty")
    }
    if redirectURL == "" {
        return fmt.Errorf("GOOGLE_REDIRECT_URL is empty")
    }

    GoogleOAuthConfig = &oauth2.Config{
        ClientID:     clientID,
        ClientSecret: clientSecret,
        RedirectURL:  redirectURL,
        Scopes: []string{
            "https://www.googleapis.com/auth/userinfo.email",
            "https://www.googleapis.com/auth/userinfo.profile",
        },
        Endpoint: google.Endpoint,
    }

    log.Println("✅ Google OAuth Config initialized successfully")
    return nil
}