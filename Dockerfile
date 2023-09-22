FROM heroiclabs/nakama-pluginbuilder:3.17.1 AS builder

ENV GO111MODULE on
ENV CGO_ENABLED 1

WORKDIR /backend
COPY go.mod .
COPY *.go .
COPY vendor/ vendor/
COPY rpcs/*.go rpcs/
COPY helpers/*.go helpers/

RUN go build --trimpath --mod=vendor --buildmode=plugin -o ./backend.so

FROM heroiclabs/nakama:3.17.1

COPY --from=builder /backend/backend.so /nakama/data/modules