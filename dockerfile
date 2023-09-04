FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . . 
RUN go build -o main .
# EXPOSE 8089 
# ENTRYPOINT ["./main"]


FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/main .
EXPOSE 8089
CMD ["/main"]