FROM golang:1.15 as base

#Install gomock (mockgen) for mocking interfaces
RUN go get github.com/golang/mock/mockgen
RUN go install github.com/golang/mock/mockgen

# Install the watcher
RUN go get github.com/codegangsta/gin

# Download deps
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy source as late in the process as possible (to speed up local builds)
COPY src/pg/migrations ./pg/migrations
COPY . .

#################
# BUILDER STAGE #
#################
FROM base AS builder
WORKDIR /app
COPY --from=base /app .
COPY --from=base /app/pg/migrations /app/pg/migrations
RUN GOOS=linux GOARCH=386\
				go build -v\
				-o app\
				src/*.go

#################
# FINAL STAGE #
#################
FROM alpine:3.7
WORKDIR /app
COPY --from=builder /app/pg/migrations /app/pg/migrations
COPY --from=builder /app/app .
CMD [ "./app"]
