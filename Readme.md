## Go server for docket app

For dev install [this](https://github.com/codegangsta/gin) and start the server as

```
gin --appPort 8000 -i
```

The api would be available at [localhost:8000](http://localhost:3000)

## Using docker

If you want to integrate docket api and do not have mysql running in your local, install docket from [here](https://docs.docker.com/desktop/mac/install/) and run below

```
docker-compose up
```

## Create Oauth client and superuser

Run below to create a oauth client:

```
go run tools/*.go oauthClient --name=web_client
```

Run below to create superuser:

```
go run tools/*.go superuser --first-name=Nayan --last-name=Das --username=dastms@gmail.com --password=password
```

If you are using docker, you need to ssh into the api server container like below and then run above commands:

```
docker exec -it docket_apiserver /bin/sh
```

### Generating swagger docs

This project is using [swaggo](https://github.com/swaggo/swag) to generate api docs

Install swag cli using below:

```
go get -u github.com/swaggo/swag/cmd/swag
```

After you add swaggo code annotations to new controllers or modify them, run below:

```
swag init
```

The swagger docs would be available at [localhost:8000/swagger/index.html](http://localhost:3000/swagger/index.html)
