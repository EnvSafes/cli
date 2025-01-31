package utils

import (
	"context"
	"envsafes/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pkg/browser"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	githubOauthConfig = &oauth2.Config{
		ClientID:     "Ov23lirk5sN1oOeVenm0",
		ClientSecret: "0d6206ab2f01ff56de00edbe50896afac8bb46b8",
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}
)

func Authenticate() {

	url := githubOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	// fmt.Printf("Visit the URL to login: %v", url)
	fmt.Println("Opening browser for GitHub OAuth login...")
	browser.OpenURL(url)

	http.HandleFunc("/callback", handleGitHubCallback)
	http.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Closing the browser...")
		os.Exit(0)
	})
	log.Fatal(http.ListenAndServe("localhost:30235", nil))
}

func handleGitHubCallback(w http.ResponseWriter, r *http.Request) {
	// defer os.Exit(0)
	code := r.FormValue("code")
	state := r.FormValue("state")
	if state != "state" {
		http.Error(w, "State did not match", http.StatusBadRequest)
		return
	}
	// fmt.Fprintf(w, "Your GitHub OAuth code is: %v %v", code, state)

	token, err := githubOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Login successful! You can close this tab now.")

	// Save the token to a file
	err = saveTokenToFile(token)
	if err != nil {
		fmt.Printf("Error saving token to file: %v\n", err)
	}

	go func() {
		_, err = http.Get("http://localhost:30235/close")
	}()
}

func saveTokenToFile(token *oauth2.Token) error {
	execpath := utils.GetExecutablePath()
	// Save the token to a file
	file, err := os.Create(execpath + "/auth.token")
	if err != nil {
		fmt.Printf("Error creating token file: %v\n", err)
		return err
	}
	defer file.Close()
	_, err = file.WriteString(token.AccessToken)
	if err != nil {
		fmt.Printf("Error writing token to file: %v\n", err)
		return err
	}
	return nil
}
