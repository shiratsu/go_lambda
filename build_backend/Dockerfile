FROM alpine as build
RUN apk add go git
RUN go env -w GOPROXY=direct
COPY go.mod go.sum ./
COPY app ./
RUN go version
RUN go mod download
ADD . .
RUN go build -o /main app/main.go
FROM alpine
COPY --from=build /main /main
ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /usr/bin/aws-lambda-rie
RUN chmod 755 /usr/bin/aws-lambda-rie
COPY build_backend/entry.sh /
RUN chmod 755 /entry.sh
ENTRYPOINT [ "/entry.sh" ]