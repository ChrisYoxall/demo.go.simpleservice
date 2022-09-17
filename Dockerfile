FROM index.docker.io/library/golang:1.18 AS build

WORKDIR /app
COPY ./src/* ./
RUN go test
RUN go build -o simpleservice

FROM gcr.io/distroless/base-debian11

WORKDIR /
COPY --from=build /app/simpleservice /simpleservice

EXPOSE 8080

CMD ["/simpleservice"]
