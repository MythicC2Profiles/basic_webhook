FROM golang:1.21 as builder

WORKDIR /Mythic/

COPY [".", "."]

RUN make build

FROM alpine

RUN apk add --update make

COPY --from=builder /main /main

WORKDIR /Mythic/

COPY [".", "."]

CMD make run