ARG buildenv=1.18

FROM golang:${buildenv} as builder
WORKDIR /app
COPY . /app
RUN go get -d -v

# Statically compile our app for use in a distroless container
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -v -o app .

# A distroless container image with some basics like SSL certificates
# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static
LABEL maintainer="Admir Trakic <atrakic@users.noreply.github.com>"

COPY --from=builder /app/app /app
COPY --from=busybox:1.36.0-musl /bin/wget /usr/bin/wget

ENV PORT=5000

ENTRYPOINT ["/app"]
