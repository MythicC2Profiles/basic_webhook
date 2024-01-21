FROM itsafeaturemythic/mythic_go_base:latest as builder

WORKDIR /Mythic/

COPY [".", "."]

RUN make build

FROM alpine

RUN apk add --update make

COPY --from=builder /main /main

WORKDIR /Mythic/

COPY [".", "."]

CMD make run