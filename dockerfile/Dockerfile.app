FROM alpine:latest AS build

ARG TEST_RUNNING=false
ARG LINT_RUNNING=false

RUN apk add --no-cache --update go gcc g++ golangci-lint

WORKDIR /go/src/github.com/naufalfmm/dayatani-farmer-api

COPY go.mod go.sum ./
RUN GO111MODULE=on go mod download

COPY . .

RUN if [ "${TEST_RUNNING}" = "true" ] ; \
        then go test ./... -count=1 -failfast || exit 1 ; \
    fi

RUN if [ "${LINT_RUNNING}" = "true" ] ; \
        then golangci-lint run || exit 1 ; \
    fi

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o dayatani-farmer-api


FROM alpine:latest

RUN apk update && apk add --no-cache tzdata
RUN apk --no-cache add ca-certificates

WORKDIR /usr/src
COPY --from=build /go/src/github.com/naufalfmm/dayatani-farmer-api/dayatani-farmer-api dayatani-farmer-api

CMD ["./dayatani-farmer-api"]