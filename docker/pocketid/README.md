# PocketID Development Environment

Lightweight PocketID container for local development of authentication functionality.

# TODO: add setup instructions

## devcontainer setup

When running the MCP server, any locally generated certificates won't be trusted by default within the devcontainer. Copy the cert into the default TLS trust store:

```sh
sudo cp certs/auth.localhost.pem /etc/ssl/certs/
```