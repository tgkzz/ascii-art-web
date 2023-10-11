FROM golang:1.20-alpine

LABEL version="1.0" \
      description="my first project which is containarized" \
      maintainer="kamamedov" \
      maintainer="akalbiat" 

WORKDIR /app

COPY go.mod .
COPY main.go .
COPY templates/ templates/
COPY internal/ internal/

RUN go build -o ascii-web .
EXPOSE 8080

ENTRYPOINT ["/app/ascii-web"]