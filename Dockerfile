FROM golang:1.17.6-alpine
RUN apk add --no-cache gcc musl-dev bash

WORKDIR /app
COPY . .

CMD ["go", "test", "./..."]
