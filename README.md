# Telegram File Explorer Bot

Exposes a local directory through a Telegram bot. Only whitelisted users can see the real contents while other users are shown a fake directory.

## Deploy

Copy `env.list` into `.env` and fill the values as needed

## Docker-compose

Example docker-compose:
```
services:
  bot:
    build:
      dockerfile: Dockerfile
      context: .
    ports:
      - "443:443"
    volumes:
      - /your/directory:/bot/files

volumes:
  files:
```

## Kubernetes

Using minikube:

1. Build docker image with: `docker build . -t telegram-file-explorer`
1. Load into minikube: `minikube image load telegram-file-explorer-bot`
1. Apply with `kubectl apply -f kubernetes-file-explorer.yaml`
