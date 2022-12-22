FROM golang:1.19-alpine as build
# Set the working directory
WORKDIR /go/src/backendify
# Copy and download dependencies using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download
# Copy the source files from the host
COPY . /go/src/backendify

RUN CGO_ENABLED=0 GOOS=linux go build -o backendify

FROM scratch
COPY --from=build /go/src/backendify/backendify  .

EXPOSE 8080
ENTRYPOINT ["/flightify"]