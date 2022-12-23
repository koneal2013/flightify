FROM golang:1.19-alpine as build
# Set the working directory
WORKDIR /go/src/flightify
# Copy and download dependencies using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download
# Copy the source files from the host
COPY . /go/src/flightify

RUN CGO_ENABLED=0 GOOS=linux go build -o flightify

FROM scratch
COPY --from=build /go/src/flightify/flightify  .
COPY --from=build /go/src/flightify/config.json  .

ENTRYPOINT ["/flightify", "--config-file=./config.json"]