# JKZDB

## Running the Database

- To start the different shards of JKZDB, run the following in `/scripts`:

```bash
go run run_servers.go
```

- To start the API request listener, run the following in `/coordinator`:

```bash
go run server.go
```

If the server does not start up properly (i.e. you don't see the Go Fiber output), you may need to change the last line in `server.go` from `app.Listen(":5000")` to `app.Listen("localhost:5000")`.

- To clean up JKZDB after you finish testing, run the following in `/scripts`:

```bash
go run stop_servers.go
```

## Testing the API

To test the API, you can use an application like Postman to send API requests to the proper URL.

## API Routes

The schema for our user model can be found in `/models/user.go`. Here are the available API routes.

- GET /api/user?index=id&key=\<pk>

- GET /api/user?index=email&key=\<email>

- POST /api/user

- POST /transaction?index=id&from=\<from_pk>&to=\<to_pk>&amount=\<amount>

- POST /transaction?index=email&from=\<from_email>&to=\<to_email>&amount=\<amount>

- PUT /api/email?old-email=\<old_email>&new-email=\<new_email>

- PUT /api/withdraw?amount=\<amount>&index=id&key=\<pk>

- PUT /api/withdraw?amount=\<amount>&index=email&key=\<email>

- PUT /api/deposit?amount=\<amount>&index=id&key=\<pk>

- PUT /api/deposit?amount=\<amount>&index=email&key=\<email>

- DELETE /api?index=id&key=\<pk>

- DELETE /api?index=email&key=\<email>

## Group Work

Matthew -
Surtaz -
Ross -
