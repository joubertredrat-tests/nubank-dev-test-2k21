FROM --platform=${BUILDPLATFORM} golang:1.16.3-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/authorize .

FROM scratch AS bin-unix
COPY --from=build /out/authorize /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM bin-${TARGETOS} AS bin
