## Go server for docket app

For dev install [this](https://github.com/codegangsta/gin) and start the server as

```
gin --appPort 8000 -i
```

The api would be available at [http://localhost:3000](localhost:3000)

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

The swagger docs would be available at [localhost:3000/swagger/index.html](http://localhost:3000/swagger/index.html)
