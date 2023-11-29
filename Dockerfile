# Build
FROM golang:1.20 as build

WORKDIR /go/src/app

ARG GIT_USER
ARG GIT_PASS
ARG GIT_HOST

ENV CGO_ENABLED=0
ENV GOPRIVATE=$GIT_HOST

COPY go.mod .
COPY go.sum .
RUN git config --global url."https://${GIT_USER}:${GIT_PASS}@${GIT_HOST}".insteadOf "https://${GIT_HOST}"
RUN go mod download
COPY . .
RUN go build -ldflags "-X main.version=$(git describe --abbrev=8 --always --tag) -s -w" -o /go/bin/app .

# Runtime
FROM alpine:3.17
COPY --from=build /go/bin/app /
CMD ["/app"]
