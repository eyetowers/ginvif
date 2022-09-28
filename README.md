# Ginvif

HTTP Server translating JSON requests into Onvif (onvif.org) WS/SOAP protocol based on the Gin web
framework.

## Usage

Run the `ginvif` binary supplying:

- `--bind`, `-b`: Address and port to serve requests on (optional).
- `--url`, `-a`: Base URL of the Onvif device.
- `--username`, `-u`: Username for authentication with the Onvif device.
- `--password`, `-p`: Password for authentication with the Onvif device.
- `--verbose`, `-v`: Print sent and received requests.

## Swagger

To regenerate the Swagger definitions, run the following in the root of the project:

```bash
swag init -d pkg/ginvif -g server.go -o docs -pd
```

To format the Swaggo doc comments, run:

```bash
swag fmt
```

## See Also

Ginvif is based on the [eyetowers/Gonvif](https://github.com/eyetowers/gonvif) Onvif client library.

## License

Ginvif is open source, released under the [MIT license](./LICENSE).
