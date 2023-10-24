FROM golang:1.20 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /parcel-master

FROM build AS run-tests

RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build /parcel-master /parcel-master

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/parcel-master"]
