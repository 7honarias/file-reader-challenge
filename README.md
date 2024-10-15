# file-reader-challenge
## Setup
Environment Variables: Define the following environment variables in a .env file or directly in your environment:
```
    HOST = HOST
    PORT = PORT
    FROM = PORT
    PASSWORD = PASSWORD
    TO = TO
    DATABASE_URL=DATABASE_URL
```

## Main Components:
- main.go: The main file that starts the Lambda application and connects to the database.
- reader.go: Reads the CSV file of transactions and stores each transaction in the database.
- db.go: Handles database initialization and closure of the PostgreSQL connection.
- report.go: Generates a report based on the total balance and a monthly summary of the transactions.
- email.go: Sends an email with the report using OAuth2 authentication.

## Execution
###  Run Locally
The project is configured with Docker and AWS Lambda Runtime Interface Emulator (RIE). To run the project:
with script

```
    ./script/up-local.sh
```
You can invoke the Lambda function locally using curl:

```
    curl -X POST http://localhost:9000/2015-03-31/functions/function/invocations -d '"transaction.csv"'
```

and finally run

```
    ./script/down-local.sh
```