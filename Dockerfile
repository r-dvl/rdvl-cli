ARG OS
ARG ARCH

FROM ghcr.io/r-dvl/golang-builder:${OS}-${ARCH}

ENV PROJECT_NAME=rdvl-cli
ENV VERSION=x.x.x

WORKDIR /home/app

COPY . .

RUN go mod download && go mod verify

RUN mkdir -p ./bin/${PROJECT_NAME}-${VERSION}.${GOOS}-${GOARCH}

# Compile binaries
CMD ["sh", "-c", "go build -o ./bin/${PROJECT_NAME}-${VERSION}.${GOOS}-${GOARCH}/${PROJECT_NAME}${EXT}"]