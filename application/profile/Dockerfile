FROM golang:1.14-alpine as build

RUN apk add --no-cache git

WORKDIR /profile-store 

COPY go.sum /profile-store/
COPY go.mod /profile-store/

COPY *.go /profile-store 

RUN go build .

RUN ls -ltr 

FROM alpine as runtime

COPY --from=build /profile-store/profile-store /app/profile-store
CMD [ "/app/profile-store" ]