FROM alpine:latest AS build

RUN apk add --no-cache --update go gcc g++

WORKDIR /go/src/github.com/naufalfmm/dayatani-farmer-api

COPY go.mod go.sum ./
RUN GO111MODULE=on go mod download

COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o dayatani-farmer-api


FROM alpine:latest

RUN apk update && apk add --no-cache tzdata
RUN apk --no-cache add ca-certificates

WORKDIR /usr/src
COPY --from=build /go/src/github.com/naufalfmm/dayatani-farmer-api/dayatani-farmer-api dayatani-farmer-api

CMD ["./dayatani-farmer-api"]