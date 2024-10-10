# Twitter API Go Application

This Go application allows you to post and delete tweets using the Twitter API. It uses OAuth1 for authentication and is designed to be run from the command line. You can either post a new tweet or delete an existing one by passing the relevant command-line arguments.

## Prerequisites

- Go 1.16+ installed on your machine.
- A Twitter Developer account and app to obtain API keys and tokens.
- `.env` file containing the following environment variables:
  - `API_KEY`
  - `API_SECRET_KEY`
  - `ACCESS_TOKEN`
  - `ACCESS_TOKEN_SECRET`

## Setup

1. Clone the repository and navigate to the project folder.
2. Create a `.env` file in the root directory and add your Twitter API credentials:
    ```
    API_KEY=your_api_key
    API_SECRET_KEY=your_api_secret_key
    ACCESS_TOKEN=your_access_token
    ACCESS_TOKEN_SECRET=your_access_token_secret
    ```
3. Install the necessary Go packages by running:
    ```bash
    go get github.com/dghubble/oauth1
    go get github.com/joho/godotenv
    ```

## Running the Application

### Post a Tweet

To post a new tweet, use the following command:
```bash
go run main.go -text "Hellooooo!!!!!"
```
Replace `"Hellooooo!!!!!"` with the text you want to tweet.

### Delete a Tweet

To delete a tweet, use the following command:
```bash
go run main.go -delete "TWEET_ID"
```
Replace `"TWEET_ID"` with the ID of the tweet you want to delete. For example:
```bash
go run main.go -delete "1844443726153646355"
```

## Optional Configuration

If you want to automatically delete the tweet after posting it, uncomment the relevant section in the `main.go` file and add a delay before deletion.