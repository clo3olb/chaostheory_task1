FROM golang

WORKDIR /app

COPY . .
# RUN go mod download

EXPOSE 80

CMD [ "go", "run", "main.go" ]