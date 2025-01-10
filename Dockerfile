FROM golang:1.23.4-alpine

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .

RUN go mod tidy

COPY . .

RUN go build -o ./tmp/main ./cmd/app/.

RUN chmod +x entrypoint.sh
ENV ENVIRONMENT=production

EXPOSE 4000

CMD ["sh", "entrypoint.sh"]
