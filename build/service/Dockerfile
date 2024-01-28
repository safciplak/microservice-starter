# Specify hash to prevent hijacking attacks
FROM aanzeenl/go-development:latest@sha256:592c033ba7b6af7da07ce85816dad880e073155125475938bdb1201705ea7b29 as development

# The build will be done in the build folder which isn't present in the root image
WORKDIR /app

# This is done so that if the mod or sum file are unchanged the docker cache is used to build the rest of your application
COPY ./go.mod ./go.sum /app/
RUN go mod download

# Build the migration tool and set as global
RUN go get -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate/
RUN go build -tags 'postgres' -ldflags="-X main.Version=$(git describe --tags)" -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate

# Copy the source of the application into the docker, ignoring everything else
COPY ./ /app

# Build the capila CLI
RUN go build -o $GOPATH/bin/capila ./src/cmd/capila/main.go

# Add an additional stage for the tests/lints
FROM development AS builder

# Compile the code to an statically compiled excecutable
RUN go build  -ldflags="-s -w" -o /app/build/service ./...

# Expose the port for local development
EXPOSE $APP_PORT

# Create the final image using the output of the builder step
FROM busybox:latest AS final

# Set the maintainer of the final scratch image
MAINTAINER "development@aanzee.nl"

# Import the time zone files
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# Import the user and group files
COPY --from=builder /user/group /user/passwd /etc/
# Import the CA certs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Import the compiled go executable
COPY --from=builder /app/build/service/microservice-starter /service/microservice-starter
# Import the Capila executable
COPY --from=builder /go/bin/capila /bin/capila
# Import the migrate executable
COPY --from=builder /go/bin/migrate /bin/migrate
# Import the migrations
COPY --from=builder /app/src/database/migrations /service/database/migrations

WORKDIR /

# Run as an unpriveleged user
USER nobody:nobody

# What should be excecuted when the docker initializes
ENTRYPOINT ["/service/microservice-starter"]

# What should be considered a healthy docker container
HEALTHCHECK --interval=5m --timeout=3s CMD curl -f http://localhost:$APP_PORT/ || exit 1

# Expose port on which the application will listen
EXPOSE $APP_PORT

# Add an additional stage for the tests/lints
FROM builder AS test
