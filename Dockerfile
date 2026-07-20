ARG VERSION=1.25.10
ARG VERSIONBINARIO=3.22

FROM golang:${VERSION}-alpine3.23 AS build

WORKDIR /gimroutine

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main /gimroutine/cmd/api/

FROM alpine:${VERSIONBINARIO}

WORKDIR /gimroutine

WORKDIR /web

COPY --from=build /gimroutine/main .

COPY web/ web/ 

CMD ["./main"]




