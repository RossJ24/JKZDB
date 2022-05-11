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

## API Endpoints

The schema for our user model can be found in `/models/user.go`. Here are the available API endpoints.

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

- DELETE /api/user?index=id&key=\<pk>

- DELETE /api/user?index=email&key=\<email>

## Link to Demo Video

We recorded a demo video of how we tested JKZDB. The Zoom link contains two recordings. Please refer to the **2nd recording** in the link, as the first recording was a trial run that got restarted. The video should be at least open to those with Yale emails. Please let us know by email if you cannot access the videos.

[Demo Video](https://yale.zoom.us/rec/share/-gEcVIAw_wcCCZ002k7yWK0iNI7iuwDwX6rXKUFA7Ep70CPqtU-UFPAjJ1XcQghD.6pZaM3Gwc5MScC51)

## Group Work

Matthew - Debugged and wrote API handlers, wrote batch 2PC RPCs, worked on report

Surtaz - Worked on 2PC RPCs, API handlers, worked on report

Ross - Worked on creating JKZDB over BoltDB, 2PC RPCs, API handlers, worked on report

As a group, we all discussed major design choices such as the database schema before implementing. We also constantly worked together and collaborated on code together over Zoom.
