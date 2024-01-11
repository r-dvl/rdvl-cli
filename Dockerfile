ARG OS

FROM ghcr.io/r-dvl/golang-builder:${OS}

ENV TAG=

WORKDIR /home/app

COPY . .

RUN go mod download && go mod verify

# Compile binaries
CMD ["sh", "-c", "go build -o ./bin/rdvl-cli.${EXT}"]