ARG OS
ARG ARCH
ARG VERSION

FROM ghcr.io/r-dvl/golang-builder:${OS}-${ARCH}

ENV PROJECT_NAME=rdvl-cli
ENV VERSION=VERSION
ENV FOLDER_NAME=${PROJECT_NAME}-${VERSION}.${GOOS}-${GOARCH}

WORKDIR /home/app

COPY . .

RUN go mod download && go mod verify

RUN mkdir -p ./bin/${FOLDER_NAME}

# Compile binaries
CMD ["sh", "-c", "go build -o ./bin/${FOLDER_NAME}/${PROJECT_NAME}${EXT}"]