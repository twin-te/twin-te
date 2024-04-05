FROM golang:1.22 as dev

WORKDIR /usr/src/twin-te/back

RUN go install github.com/bufbuild/buf/cmd/buf@latest \
  && go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest \
  && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
  && go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest \
  && go install github.com/spf13/cobra-cli@latest \
  && go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

CMD [ "air" ]
