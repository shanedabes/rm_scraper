FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o rm_scraper

# final stage
FROM alpine

WORKDIR /app
COPY --from=build-env /src/rm_scraper /app/
ENTRYPOINT ./rm_scraper
