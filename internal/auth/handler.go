package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"transportation/internal/config"
	"transportation/internal/database"

	"transportation/internal/utils"

	"github.com/gin-gonic/gin"
)

func GoogleLogin(c *gin.Context) {

	  if config.GoogleOAuthConfig == nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "OAuth config not initialized",
        })
        return
    }
	state := randomState()
	// Persist state in an HttpOnly cookie for callback validation
	//c.SetCookie("oauth_state", state, 600, "/", "", false, true)

	authURL := config.GoogleOAuthConfig.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, authURL)
}

func GoogleCallback(c *gin.Context) {
	// state := c.Query("state")
	// cookieState, err := c.Cookie("oauth_state")
	// if err != nil || state == "" || state != cookieState {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid state"})
	// 	return
	// }

	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code not provided"})
		return
	}

	token, err := config.GoogleOAuthConfig.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code exchange failed: " + err.Error()})
		return
	}

	userInfo, err := fetchGoogleUser(token.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userinfo fetch failed: " + err.Error()})
		return
	}

	user, err := saveOrUpdateUser(userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user save failed: " + err.Error()})
		return
	}

	jwtToken, err := utils.GenerateJWT(user.ID, user.Email, user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "jwt generation failed: " + err.Error()})
		return
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "https://mstradingcorporation.xyz"
	}
	redirectURL := fmt.Sprintf("%s/auth/success?token=%s", frontendURL, jwtToken)
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

func randomState() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func fetchGoogleUser(accessToken string) (*GoogleUserInfo, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", accessToken))
	if err != nil {
		return nil, fmt.Errorf("call google api: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read google response: %v", err)
	}

	var userInfo GoogleUserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, fmt.Errorf("parse google response: %v", err)
	}
	return &userInfo, nil
}

func saveOrUpdateUser(googleUser *GoogleUserInfo) (*User, error) {
	var user User
	result := database.DB.Where("google_id = ?", googleUser.ID).First(&user)

	if result.Error != nil {
		user = User{
			GoogleID: googleUser.ID,
			Email:    googleUser.Email,
			Name:     googleUser.Name,
			Picture:  googleUser.Picture,
		}
		if err := database.DB.Create(&user).Error; err != nil {
			return nil, fmt.Errorf("create user: %v", err)
		}
	} else {
		needsUpdate := user.Email != googleUser.Email || user.Name != googleUser.Name || user.Picture != googleUser.Picture
		if needsUpdate {
			user.Email = googleUser.Email
			user.Name = googleUser.Name
			user.Picture = googleUser.Picture
			if err := database.DB.Save(&user).Error; err != nil {
				return nil, fmt.Errorf("update user: %v", err)
			}
		}
	}
	return &user, nil
}
