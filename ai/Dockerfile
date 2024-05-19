FROM golang:alpine3.19 AS build

WORKDIR /src

COPY go.mod go.sum ./

COPY . ./

RUN set -xe; \
    CGO_ENABLED=0 \
    go build \
      -buildmode=pie \
      -tags netgo \
      -o /llamakraft/ai main.go\
    ;

FROM scratch

COPY ./.env .env
COPY --from=build /llamakraft /llamakraft
COPY --from=build /lib/ld-musl-x86_64.so.1 /lib/ld-musl-x86_64.so.1