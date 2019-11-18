# Dexter

An attempt at tackling a classic problem with modern tools.
I will add more content once this project proves to be useful.

## Storage

The Dexter store is a content-addressable system, where everything is indexed
with a SHA-224 hash. Dexter also provides a well defined database interface,
allowing developers to easily write their own storage engine.

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

### API Reference

The Dexter Web API is a [REST](https://en.wikipedia.org/wiki/Representational_state_transfer) API that you often come across on the Web.
This section is temporary until there's a dedicated home for the reference.

#### Authentication

TBD: Likely support Basic auth and JSON Web Tokens (JWT).

#### Subscriptions

List of subscriptions that Dexter is configured to collect.

```
GET /subscriptions
```

## Dependencies

Dexter currently has no external dependencies. Let's keep it that way
unless there's a very strong reason.
