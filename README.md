## Rate Limiter Challenge

This repo is the delivering of a Go course assignment.

**Important**

The rate limiter package repo is

[philippe-berto/rate-limiter](https://github.com/philippe-berto/rate-limiter)

And it's imported here for test purposes.

### Running and Testing

Just run

```
docker-compose up -d
```

And make requests to `http://localhost:8080/`

You can make requests with API_KEY on header and test the limiter by token, or with no api key set and test the limiter by IP.

To change configuration, just change the env vars present in the `docker-compose.yml` for the `rl-tester` service.

You can use `request.http` file to make the requests with REST Client extension on VS Code.
