# Build the manager binary
FROM 172.22.50.227/system_containers/golang:1.18.0-bullseye as builder

ENV GOPROXY https://repo.huaweicloud.com/repository/goproxy,direct

COPY . /k8s-resource-operator

WORKDIR /k8s-resource-operator

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -v -mod=vendor -o manager -ldflags="-w -s" main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM 172.22.50.227/system_containers/distroless-static:nonroot
WORKDIR /
COPY --from=builder /third-party-config-source-adaptor/manager .
#COPY bin/manager /manager
USER 65532:65532

ENTRYPOINT ["/manager"]
