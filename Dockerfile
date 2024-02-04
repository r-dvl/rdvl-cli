ARG OS
ARG ARCH

FROM ghcr.io/r-dvl/golang-builder:${OS}-${ARCH}

ENV TAG=

WORKDIR /home/app

COPY . .

RUN go mod download && go mod verify

RUN mkdir ./bin/rdvl-cli.${GOOS}-${GOARCH}

# Compile binaries
CMD ["sh", "-c", "go build -o ./bin/rdvl-cli.${GOOS}-${GOARCH}/rdvl-cli${EXT}"]