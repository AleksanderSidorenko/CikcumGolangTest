FROM iron/go:dev
WORKDIR /app

COPY . .

# CMD [ "RUN apt-get update && apt-get install -y bzr" ]
RUN go get labix.org/v2/mgo
RUN go get bzr
RUN go build
ENTRYPOINT ["./Processing"]