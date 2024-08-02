FROM golang:1.22.5 AS build-stage
WORKDIR /bot

COPY . .
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o telegram-file-explorer

# ---

FROM build-stage AS run-test-stage
RUN go test -v ./...

# ---

FROM alpine AS release-stage
WORKDIR /bot

COPY .env .env
COPY --from=build-stage /bot/telegram-file-explorer /bot/telegram-file-explorer
COPY --from=build-stage /bot/testfiles /bot/testfiles
EXPOSE 443

ENTRYPOINT ["/bot/telegram-file-explorer"]
