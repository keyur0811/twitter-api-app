# Twitter API App

## Introduction

This project is a simple Go application that interacts with Twitter’s API v2 to post and delete tweets. The purpose of this assignment is to provide hands-on experience in using APIs, OAuth1 authentication, and error handling in Go. By working with this app, you will learn how to authenticate requests with OAuth, make API requests to post tweets, and delete tweets programmatically.

## Setup Instructions

### 1. Set Up a Twitter Developer Account

To interact with the Twitter API, you need to create a Twitter Developer account:

1. Go to [Twitter Developer Portal](https://developer.twitter.com/).
2. Create a new project and Twitter App.
3. Once your app is created, navigate to the **Keys and Tokens** section to generate your API credentials.

### 2. Generate API Keys

You will need the following credentials from your Twitter Developer app:
- **API Key**
- **API Secret Key**
- **Access Token**
- **Access Token Secret**

These credentials are required to authenticate API requests from your application.

### 3. Create a `.env` File

In the root directory of the project, create a `.env` file to store your Twitter API credentials:

```
API_KEY=your_api_key_here
API_SECRET_KEY=your_api_secret_key_here
ACCESS_TOKEN=your_access_token_here
ACCESS_TOKEN_SECRET=your_access_token_secret_here
```

Replace the placeholder values with the actual credentials from your Twitter Developer account.

### 4. Run the Program

To run the application, you need to use the Go runtime. You can either run the program directly or build an executable.

- **Post a Tweet**:

  Use the following command to post a tweet with the text of your choice:

  ```bash
  go run main.go -text "Hellooooo!!!!!"
  ```

- **Delete a Tweet**:

  If you want to delete a tweet, you need to provide the tweet ID. Use the following command to delete a tweet:

  ```bash
  go run main.go -delete "1844443726153646355"
  ```

The tweet ID is returned when you post a tweet and is unique for each tweet.

## Program Details

### Posting a New Tweet

When you post a new tweet, the program sends an HTTP POST request to the Twitter API v2 endpoint:

- **Endpoint**: `https://api.twitter.com/2/tweets`
- **Request Body**: The tweet text is sent as JSON in the format `{"text": "Your tweet message"}`.
- **Response**: If the tweet is successfully posted, the API responds with a `201 Created` status and returns the tweet ID in the response body.

Example of a successful response:
```json
{
  "data": {
    "id": "1234567890123456789",
    "text": "Hellooooo!!!!!"
  }
}
```

### Deleting an Existing Tweet

To delete a tweet, the program sends an HTTP DELETE request to the following endpoint:

- **Endpoint**: `https://api.twitter.com/2/tweets/{tweetID}`
- **Response**: If the tweet is successfully deleted, the API responds with a `200 OK` status.

Example of a successful delete response:
```json
{
  "data": {
    "deleted": true
  }
}
```

### Example Commands

- **Posting a Tweet**:
  ```bash
  go run main.go -text "Hello World!"
  ```
- **Deleting a Tweet**:
  ```bash
  go run main.go -delete "1234567890123456789"
  ```

### Error Handling

The program includes robust error handling to manage various scenarios:

1. **Invalid API Credentials**: If your API credentials are invalid or missing, the program will fail to authenticate with Twitter and return an error.
2. **Tweet Not Found (404)**: If you attempt to delete a tweet that doesn’t exist, the program will handle the 404 response and notify you that the tweet was not found.
3. **Rate Limiting (429)**: If you exceed Twitter’s rate limits, the program will return a message explaining that you’ve exceeded the rate limit and need to wait before making additional requests.
4. **Unexpected Errors**: The program handles unexpected API errors and prints relevant details to the console for troubleshooting.

Example error messages:
- **Unauthorized**: "Unauthorized. Please check your API keys and tokens."
- **Tweet Not Found**: "Tweet not found. Invalid tweet ID."
- **Rate Limit Exceeded**: "Rate limit exceeded. Try again later."