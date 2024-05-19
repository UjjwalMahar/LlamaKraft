# FROM golang:alpine3.19

# WORKDIR /app

# ENV DB_USER=authdb_owner
# ENV DB_PASSWORD=6JAk5xOzeVWa
# ENV DB_HOST=ep-raspy-shadow-a2cvb3uv.eu-central-1.aws.neon.tech
# ENV DB_NAME=authdb

# COPY go.mod go.sum ./

# RUN go mod download

# COPY . ./

# RUN CGO_ENABLED=0 GOOS=linux go build -o /llamakraft

# EXPOSE 8080

# CMD ["/llamakraft"]



FROM golang:alpine3.19 AS build

WORKDIR /src

COPY go.mod go.sum ./

COPY . ./

RUN set -xe; \
    CGO_ENABLED=0 \
    go build \
      -buildmode=pie \
      -tags netgo \
      -o /llamakraft/auth main.go\
    ;

FROM scratch

COPY ./.env .env
COPY --from=build /llamakraft /llamakraft
COPY --from=build /lib/ld-musl-x86_64.so.1 /lib/ld-musl-x86_64.so.1