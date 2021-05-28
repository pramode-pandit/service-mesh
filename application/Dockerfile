
FROM golang:1.15.11-alpine3.13 AS build

# Set necessary environmet variables needed for our image
# ENV GO111MODULE=on \
#     CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64

WORKDIR /my-app
COPY *.go  .

RUN ls -lha 
RUN go version
RUN go build register.go
Run ls -lha


FROM alpine:3.13

RUN apk update && apk add curl && mkdir -p  /my-app/data
COPY --from=build /my-app/register  /my-app/register 

EXPOSE 8080
ENTRYPOINT ["/my-app/register"]