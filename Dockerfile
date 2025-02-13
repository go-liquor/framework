FROM golang:1.22-alpine3.20 as build
WORKDIR /var/build
COPY . /var/build
RUN apk add make
RUN make build


FROM alpine:3.20
WORKDIR /var/app
COPY --from=build /var/build/bin/app /var/app
ENTRYPOINT ["/var/app/app"]