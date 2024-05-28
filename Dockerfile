FROM golang:1.22 AS build
WORKDIR /app
COPY go.mod /app/go.mod
RUN go mod download
COPY . /app
RUN echo "appuser:x:10001:10001:App User:/:/sbin/nologin" > /etc/minimal-passwd

ENV GOOS=linux
ENV CGO_ENABLED=0
RUN go build -ldflags="-w -s" -o /webapp cmd/server/main.go

FROM scratch
COPY --from=build /etc/minimal-passwd /etc/passwd
USER appuser
COPY --from=build --chown=10001:10001 /webapp /webapp
ENV GIN_MODE=release
EXPOSE 8080
ENTRYPOINT [ "/webapp" ]
