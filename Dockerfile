# FROM golang:1.19-alpine3.17

# WORKDIR /app

# COPY ./go /app

# RUN cd /app && go build -o library

# CMD [ "/app/library" ]

# EXPOSE 8080




# This is Dockerfile for Github_action
FROM golang:1.19-alpine3.17

WORKDIR /app

COPY ./library /app

CMD [ "/app/library" ]

EXPOSE 8080