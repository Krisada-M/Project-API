
# command
# docker build -t api-hairsalon -f Dockerfile . 
# docker run -dp 5540:80 --name api-hairsalon  api-hairsalon

# docker run -dp {PortLocal}:80 --name api-hairsalon  api-hairsalon


FROM golang:latest

RUN mkdir /app

ARG . /app

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 80

ENV TZ=Asia/Bangkok

RUN go build -o main

CMD [ "/app/main" ]