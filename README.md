# Proxy Server Application

## Getting Started

```sh
git clone https://github.com/canyouhearthemusic/proxy-server.git

cd proxy-server

make build
make up
```

## Endpoints

`GET /health` – Health check endpoint.
\
`POST /request` – Proxy Request
\
#### Request Template
```json
{
    "method": "get",
    "url": "https://www.google.com",
    "headers": {
        "Content-Type": "application/json",
    }
}
```

#### Response
```json
{
    "id": /* uuid */,
    "status": /* status code of response */,
    "headers": {
        ...
    }
}
```