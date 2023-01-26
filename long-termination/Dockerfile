ARG REPO_NAME="long-termination"
ARG APP_NAME="long-termination"

FROM golang:1.19 as build
ARG REPO_NAME
ARG APP_NAME
COPY . /go/src/${REPO_NAME}
WORKDIR /go/src/${REPO_NAME}
RUN go build -o /go/bin/${APP_NAME} ./

FROM gcr.io/distroless/base
ARG APP_NAME
COPY --from=build /go/bin/${APP_NAME} /${APP_NAME}
CMD ["/long-termination"]