FROM golang:1.12
ARG BRANCH
ARG GIT
WORKDIR /
RUN git clone -b $BRANCH $GIT
WORKDIR /wails/cmd/wails
RUN go install
