# OAuth2 Proxy Development Environment

Lightweight OAuth2 Proxy container for local development of authentication functionality.

# TODO: add setup instructions

```sh
python3 -c 'import os,base64; print(base64.urlsafe_b64encode(os.urandom(32)).decode())' > ./secrets/cookie-secret.txt && truncate -s -1 ./secrets/cookie-secret.txt
```