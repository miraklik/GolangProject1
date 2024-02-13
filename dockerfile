FROM golang:1.21.5-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev
# Зависимости 
COPY ["go.mod", "go.sum", "./"]
RUN go mod donwload

#build code 
COPY app ./
RUN go build -o ./bin/app goproject/app/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/app /
COPY configs/configs.yaml /configs.yaml

CMD ["/app"]

