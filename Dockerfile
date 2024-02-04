ARG OS
ARG ARCH

FROM ghcr.io/r-dvl/golang-builder:${OS}-${ARCH}

ENV TAG=
ENV PROJECT_NAME=rdvl-cli
ENV FOLDER_NAME=${PROJECT_NAME}-${TAG}.${GOOS}-${GOARCH}

WORKDIR /home/app

COPY . .

RUN go mod download && go mod verify

RUN mkdir -p ./bin/${FOLDER_NAME}

# Compile binaries
CMD ["sh", "-c", "go build -o ./bin/${FOLDER_NAME}/${PROJECT_NAME}${EXT}"]