FROM pavhov/golang:1.12 AS build

ENV GIN_MODE=debug

USER root

WORKDIR /go/src/histories

COPY . .

RUN make dep-install
RUN make setup-env
RUN make build-app

CMD ["make", "run-app"]
