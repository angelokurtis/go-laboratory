FROM golang:1.17.8 as builder

WORKDIR /laboratory

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/

RUN go generate ./...
RUN GOOS=linux GOARCH=amd64 go build -a -o list-ns ./cmd/kubernetes-client

FROM amazon/aws-cli:2.4.25

WORKDIR /root
COPY kubeconfig.yaml .kube/config
COPY --from=builder /laboratory/list-ns .

ENTRYPOINT ["./list-ns"]
