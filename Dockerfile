FROM golang:1.21.0

# Set the GOPATH to /go
ENV GOPATH=/go

RUN go install github.com/cweill/gotests/gotests@v1.6.0
RUN go install github.com/fatih/gomodifytags@v1.16.0
RUN go install github.com/josharian/impl@v1.1.0
RUN go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install golang.org/x/tools/gopls@latest


# Set the working directory inside the container
WORKDIR /go/src/app

# Expose the port that gopls uses (default is 4389)
EXPOSE 4389
EXPOSE 8080

# Define the command to run gopls
CMD ["gopls"]

ENTRYPOINT /usr/bin/bash
