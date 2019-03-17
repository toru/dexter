# Dexter

An attempt at tackling a classic problem with modern tools.
I will add more content once this project proves to be useful.

## The Dexter Web API

You may choose to run the built-in Web API application in front of the collected
content. The application is enabled by providing the configuration for the API
server. Dexter otherwise runs only as a data collector.

Here's a minimal configuration to enable Dexter's Web API.

```
[web]
port = 8081
```

### TLS

TBD: Likely support TLS via the [tls] configuration.

### Authentication

TBD: Likely support Basic auth and JSON Web Tokens (JWT).

## Dependencies

Dexter currently has no external dependencies. Let's keep it that way
unless there's a very strong reason.
