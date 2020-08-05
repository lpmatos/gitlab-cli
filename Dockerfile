FROM golang:alpine as builder

WORKDIR /build

COPY [ ".", "." ]

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o ./bin/glabby main.go

FROM alpine:3.12 as release

LABEL maintainer="Cloud Team Stefanini - cloud@stefanini.com" \
      org.label-schema.name="Glabby" \
      org.label-schema.version="1.0.0" \
      org.label-schema.release-data="05-08-2020" \
      org.label-schema.description="Glabby is a CLI to interact with a pipeline in GitLab CI"

COPY --from=builder [ "/build/bin/glabby", "/usr/local/bin/glabby" ]

RUN chmod +x /usr/local/bin/glabby

RUN apk --no-cache add ca-certificates shadow git bash

RUN addgroup -S docker && adduser -S -G docker glabby && usermod -aG root glabby

USER glabby

CMD ["glabby"]
