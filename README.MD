# shortURL

Hobby project written in Go that basically generates a short-ID for a given URL. Uses redis to store and retrieve data.

# Usage
- Needs a Redis server running on port 6379. If you have Docker installed, run `docker pull redis && docker run -p 6379:6379 redis"
- Runs on port 8080. Hit localhost:8080/ping to make sure the server is up and running.
- The endpoint /api/add/:url adds a particular URL to the storage and returns the associated short-ID for it.
- The endpoint /api/get/:shortID retrieves the particular URL associated with the short-iD.

# Note
- At the moment, the project does not support URLs with the "http://" or "https://" prefix. Please use URLs without the prefix. Like "google.com" instead of "https://google.com".
