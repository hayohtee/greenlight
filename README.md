# greenlight

## Description
A REST API for retrieving and managing information about movies with core functionality similar to (Open Movie Database API)

## How to run
1. Clone the repository:

  ```bash
  git clone https://github.com/hayohtee/greenlight.git
  cd greenlight
  ```
2. Install dependencies:

  ```bash
  go get ./...
  ```
3. Ensure PostgreSQL is installed
   
4. Create DATABASE greenlight and ROLE greenlight using PostgreSQL (psql)

   ```bash
   // First connect as postgres user
   psql -U postgres -d postgres

   // Inside psql shell run the following
   CREATE DATABASE greenlight;
   CREATE ROLE greenlight WITH ENCRYPTED PASSWORD 'your password';
   GRANT ALL PRIVILEGES ON DATABASE greenlight TO greenlight;

   \c greenlight postgres
   // You are now connected to database "greenlight" as user "postgres".
   GRANT ALL ON SCHEMA public TO greenlight;
   ```
5. Set up environment variables:
  Create enivronment variable for GREENLIGHT_DB_DSN

  ```bash
  export GREENLIGHT_DB_DSN='postgres://greenlight:yourpassword@localhost/greenlight?sslmode=disable';
  ```
  Ensure the environemt variable was setup properply (you might need to close and open the terminal again).

6. Run the command
   
  ```bash
  go run ./cmd/api -db-dsn=$GREENLIGHT_DB_DSN
  ```
  If you want to customize the default configurations, there is commandline flags option. Run the following
  to show the list of all commamdline options.
  
  ```bash
  go run ./cmd/api -help 
  ```

  **For Makefile users**

  Make sure you create .envrc in the project directory.
  <br>
  inside .envrc
  
  ```
  GREENLIGHT_DB_DSN='postgres://greenlight:yourpassword@localhost/greenlight?sslmode=disable';
  ```

  Then run:
  ```bash
  make run/api
  ```
  You can also run the following to see available commands
  ```bash
  make
  ```

## Dependencies
- github.com/go-mail/mail
- github.com/julienschmidt/httprouter
- github.com/lib/pq
- github.com/tomasen/realip
- golang.org/x/crypto
