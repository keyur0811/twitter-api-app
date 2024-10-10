package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	// "time"

	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

// Function to load environment variables from .env file
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// Function to create a new OAuth1 client
func getOAuth1Client() *http.Client {
	apiKey := os.Getenv("API_KEY")
	apiSecretKey := os.Getenv("API_SECRET_KEY")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	config := oauth1.NewConfig(apiKey, apiSecretKey)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	return httpClient
}

// Function to post a tweet
func postTweet(httpClient *http.Client, tweetText string) (string, error) {
	// Twitter API v2 endpoint for posting a new tweet
	endpoint := "https://api.twitter.com/2/tweets"

	// Prepare the tweet data
	tweetData := map[string]string{"text": tweetText}
	tweetBody, _ := json.Marshal(tweetData)

	// Create a new request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(tweetBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	response, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to post tweet: %v", err)
	}
	defer response.Body.Close()

	// Handle response
	if response.StatusCode == http.StatusCreated {
		// Successfully posted the tweet
		var result struct {
			Data struct {
				ID string `json:"id"`
			} `json:"data"`
		}
		if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
			return "", fmt.Errorf("failed to decode response: %v", err)
		}
		return result.Data.ID, nil
	}

	// Handle other unexpected responses
	bodyBytes, _ := io.ReadAll(response.Body)
	log.Printf("Response body: %s", string(bodyBytes))
	return "", fmt.Errorf("failed to post tweet: unexpected status %s", response.Status)
}

// Function to delete a tweet
func deleteTweet(httpClient *http.Client, tweetID string) {
	endpoint := fmt.Sprintf("https://api.twitter.com/2/tweets/%s", tweetID)

	// Make a DELETE request to delete the tweet
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Make the request
	response, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to delete tweet: %v", err)
	}
	defer response.Body.Close()

	// Handle response errors
	handleResponseErrors(response)
}

// Function to handle response errors
func handleResponseErrors(response *http.Response) {
	if response.StatusCode == http.StatusOK {
		fmt.Println("Operation successful!")
	} else if response.StatusCode == http.StatusUnauthorized {
		fmt.Println("Unauthorized. Please check your API keys and tokens.")
	} else if response.StatusCode == http.StatusNotFound {
		fmt.Println("Tweet not found. Invalid tweet ID.")
	} else if response.StatusCode == http.StatusTooManyRequests {
		fmt.Println("Rate limit exceeded. Try again later.")
	} else {
		fmt.Printf("Unexpected error. Status Code: %d\n", response.StatusCode)
	}
}

func main() {
	// Load environment variables
	loadEnv()

	// Initialize OAuth1 client
	httpClient := getOAuth1Client()

	// Parse command-line flags
	tweetText := flag.String("text", "", "Text of the tweet to post")
	deleteID := flag.String("delete", "", "ID of the tweet to delete")
	flag.Parse()

	// If the delete flag is provided, delete the specified tweet
	if *deleteID != "" {
		fmt.Printf("Attempting to delete tweet with ID: %s\n", *deleteID)
		deleteTweet(httpClient, *deleteID)
		return
	}

	// Ensure tweet text is provided
	if *tweetText == "" {
		fmt.Println("Please provide tweet text using the -text flag.")
		return
	}

	// Post a new tweet
	tweetID, err := postTweet(httpClient, *tweetText)
	if err != nil {
		log.Fatalf("Error posting tweet: %v", err)
	}

	// Log the posted tweet ID
	fmt.Printf("Posted tweet with ID: %s\n", tweetID)

	// Uncomment this code to automatically delete the post you created while posting new tweet
	// Wait for 2 seconds before deleting the tweet to ensure it's fully registered
	// time.Sleep(2 * time.Second)

	// Delete the posted tweet (optional, you can comment this out if you don't want to delete immediately)
	// deleteTweet(httpClient, tweetID)
}
