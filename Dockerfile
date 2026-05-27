FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG VERSION=dev
RUN CGO_ENABLED=0 go build -ldflags="-X main.appVersion=${VERSION}" -o /pulsed .

FROM scratch
COPY --from=build /pulsed /pulsed
ENTRYPOINT ["/pulsed"]
