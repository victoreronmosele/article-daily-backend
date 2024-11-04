# Article Daily Backend

A Go backend service that fetches news articles and sends notifications via Firebase Cloud Messaging (FCM).

## Features

- Fetches articles from NewsData API
- Sends push notifications using Firebase Cloud Messaging
- RESTful API endpoints for article management

## Prerequisites

- Go 1.18 or higher
- Firebase account and credentials (for FCM)
- NewsData API key
- Google Cloud Platform account (for deployment)

## Deployment

### Google Cloud Run

The server can be deployed to Google Cloud Run using the following command:

```bash
gcloud alpha run deploy article-daily-backend --source . --project article-daily --set-build-env-vars GOOGLE_BUILDABLE=./server
```
