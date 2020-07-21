# Go Web-Service

This is a project example build with Lancer-Kit tool set.

#### What service do?

- Listen on ports `:5000` and `:9102`

- Receive JSON logs on `:5000/logs`

- Serve Prometheus metrics on `:9102/metrics`

- Compute number of unique IP addresses in logs since service start

- Create custom Prometheus metric "unique_ip_addresses"

- Publish your result in public Git repository

#### Quick start

1. Clone this repo:

```shell script
git clone https://github.com/sheb-gregor/log_cache
cd log_cache
```

2. Prepare a local configuration:

```shell script
## here is secrets and other env vars
cp ./env/tmpl.env ./env/local.env

## here is configuration details
cp ./env/tmpl.config.yaml ./env/local.config.yaml
```

3. Build docker image:

```shell script
make build_docker image=log_cache config=local
```

4. Start all:

```shell script
docker-compose up -d
```

## Development 

- Get `forge` — a tool for code generation:

```shell script
go get -u github.com/lancer-kit/forge
```



