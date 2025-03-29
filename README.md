# RSS Feed Aggregator

## Overview
This is an **RSS Feed Aggregator** implemented in Go, which allows users to register RSS feeds and actively fetches new articles from registered sources. The aggregator stores feeds and their content in a **PostgreSQL database**, making it easy to track updates from multiple sources.

## Features
- **User Registration**: Users can create accounts and receive an API key for authentication.
- **Feed Registration**: Users can register RSS feeds by providing a name and URL.
- **Automatic Fetching**: The aggregator periodically fetches new articles from registered feeds.
- **PostgreSQL Storage**: Feeds and articles are stored in a structured database for retrieval and analysis.
- **Secure API Access**: API key authentication ensures secure interactions with the aggregator.

## Installation
To run the RSS Feed Aggregator, ensure you have **Docker** installed on your system.

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/RssAggregator.git
   cd RssAggregator
   ```

2. Navigate to the `dockerComposeForPGSQLandPDADMIN` directory:
   ```sh
   cd dockerComposeForPGSQLandPDADMIN
   ```

3. Set up PostgreSQL and pgAdmin using Docker Compose:
   ```sh
   docker-compose up -d
   ```

4. Then, return to the project root directory:
   ```sh
   cd ..
   ```

5. Build and run the aggregator:
   ```sh
   go run main.go
   ```

## Steps to Use

### 1. Create a User
Send the following JSON request to `http://127.0.0.1:8000/v1/users`:
```json
{
  "name": "john doe"
}
```
- This will generate an **API key** that must be used for authentication in subsequent requests.

### 2. Register an RSS Feed
Send the following JSON request to `http://127.0.0.1:8000/v1/feeds`, **__including your API key in the request headers:__**
```json
{
  "name": "example Blog",
  "url": "https://example.com/index.xml"
}
```
- The aggregator will start fetching articles from the specified RSS feed and store them in the PostgreSQL database.

## Future Enhancements
- **Image Scraping**: Implement image scraping if the user wants to.
- **User-Defined Fetch Intervals**: Allow users to customize the frequency of feed updates.
- **Web UI for Managing Feeds**: Provide a graphical interface for easier management.
- **Enhanced API Endpoints**: Include pagination, search, and filtering for stored articles.
- **Notifications**: Implement email or webhook notifications for new feed updates.

## Contribution
Contributions are welcome! Feel free to fork the repository and submit pull requests.

## License
This project is licensed under the **MIT License**.

