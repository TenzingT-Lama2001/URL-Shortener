![gotrim](https://github.com/TenzingT-Lama2001/URL-Shortener/assets/52858291/4f9c4643-5b50-4aa0-98e1-19e84a6ff572)
<p align="center">
  <h3 align=Your URL Shortening Solution.</h3>
</p>



# Go Trim

Your URL Shortening Solution.

![live](https://github.com/TenzingT-Lama2001/URL-Shortener/assets/52858291/ca0fbae2-ae90-425b-801a-be7241811402)

## Live

The site is live on:
- Backend: https://go-trim.tenzing121.com.np/
- Frontend: https://url-shortener-next-teal.vercel.app/

## Getting started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them :

- [Git](https://git-scm.com/)
- [Go](https://go.dev/doc/install)
- [Node.js](https://nodejs.org/)

---

### Installation

1. Clone the git repository

   ```bash
   git clone https://github.com/TenzingT-Lama2001/URL-Shortener.git
   ```

1. Go into the project directory

   ```bash
   cd URL-Shortener/
   ```

1. Install  dependencies

   ```bash
   go mod tidy
   ```
   
1. Copy .env.example to .env

   ```bash
   DOMAIN = "your-backend-url"
   ```
   
1. Run the program

   ```bash
   go run main.go
   ```

1. Run the test

   ```bash
   go test -v ./...
   ```

## Endpoints

#### /shorten 
#### Description
Accepts a JSON payload with a long URL and returns a short URL.

- Method: `POST`
- Path: `/shorten`
- Body: `longURL` (string): The long URL to be shortened.
- Headers:
  - `Content-Type: application/json`

```json
{
  "longURL": "https://github.com/TenzingT-Lama2001/URL-Shortener"
}
```



### /{shortCode} 
#### Description
Redirects users to the original long URL based on the short code provided.

- Method: `GET`
- Path: `/{shortCode}`

```json
{
    "shortURL": "http://localhost:3000/cCZAGC"
}
```
