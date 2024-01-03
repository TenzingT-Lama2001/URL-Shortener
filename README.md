

## Getting started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them :

- [Git](https://git-scm.com/)
- [Go](https://go.dev/doc/install)

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

1. Run the program

   ```bash
   go run main.go
   ```

1. Run the test

   ```bash
   go test -v ./...
   ```
