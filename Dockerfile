# Build
FROM golang:1.20 as build
COPY . /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -a -o kvs

# Publish
FROM "scratch"
COPY --from=build /src/kvs .
EXPOSE 8080
CMD ["/kvs"]
