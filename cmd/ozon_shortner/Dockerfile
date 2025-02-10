FROM golang:1.23.2-alpine AS build

RUN apk add --no-cache git

WORKDIR /src
COPY go.mod go.sum ./

COPY vendor vendor ./
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.Version=$(git rev-parse --short HEAD)'" -o /bin/ozon_shortner ./cmd/ozon_shortner

FROM alpine

COPY --from=build /bin/ozon_shortner /bin/ozon_shortner
COPY --from=build /src/internal/database/postgresql/migrations /migrations

ENTRYPOINT ["/bin/ozon_shortner"]