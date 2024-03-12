FROM golang:1.20
WORKDIR /app
COPY ./conf/conf.yaml conf/
COPY ./script/cred/*.pem script/cred/
COPY ./migrations migrations/
COPY ./build/app .

EXPOSE 8081

# Run
CMD ["./app"]