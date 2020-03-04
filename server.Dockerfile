# build stage
FROM golang:alpine AS build-env
ENV GOPROXY=direct
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD . /src
RUN cd /src && go build -o service server/main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/service /app/
COPY scripts/entrypoint.sh /app/
RUN chmod +x ./entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]
CMD []
