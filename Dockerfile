# get golang image for build as workspace
FROM golang:1.20 AS build

ENV PROJECT="bot"
# make build dir
RUN mkdir /${PROJECT}
WORKDIR /${PROJECT}
COPY go.mod go.sum ./

# download dependencies if go.sum changed
RUN go mod download
COPY . .

RUN make build

# create image with new binary
FROM scratch AS deploy

ENV PROJECT="bot"
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /${PROJECT}/migrations /migrations
COPY --from=build /${PROJECT}/bin/${PROJECT} /${PROJECT}

CMD ["./bot"]