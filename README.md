# rm_scraper

A simple go app to pull the latest running man episode from [myrunningman](https://myrunningman.com)
and add it to transmission.

## Usage

    Usage:
      rm_scraper [OPTIONS]

    Application Options:
      -H, --host=     Transmission server address (default: localhost) [$RMS_HOST]
      -P, --port=     Transmission server port (default: 9091) [$RMS_PORT]
      -s, --secure    Connect to transmission using tls [$RMS_SECURE]
      -u, --user=     Transmission server user [$RMS_USER]
      -p, --password= Transmission server password [$RMS_PASSWORD]

    Help Options:
      -h, --help      Show this help message

Transmission connection can be configured using either flags or environment variables.

## Installation

Choose a method from:

- Download the latest release from github releases.
- Compile the latest master by running `go build`.
- Use the docker image found at [shanedabes/rm_scraper](https://hub.docker.com/r/shanedabes/rm_scraper).
