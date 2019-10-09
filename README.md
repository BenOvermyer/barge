# Barge

Barge is a command line tool for viewing multiple Docker Swarms through Portainer.

## Usage

### Get Barge

```bash
curl -L https://github.com/BenOvermyer/barge/releases/download/0.3.0/barge-linux-amd64 -o barge
chmod +x barge
```

### Config & Run

Optionally, these values can be provided in YAML notation via `barge.yaml`.

```bash
export PORTAINER_URL=https://portainer.mysite.com
export PORTAINER_USERNAME=myuser
export PORTAINER_PASSWORD=mypass
barge -h
```

## Contributing

During local development, you can use `go run` to build from source and run the target application all in one shot.

```bash
go run .
```

If you'd like to test it with arguments, you can append them. Here's an example with the container command.

```bash
go run . container --help
```
