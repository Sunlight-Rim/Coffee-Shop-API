# Coffee Shop API

Example Go application with Clean Architecture pattern. \
Contains CRUD operations for `Users`, `Coffee` and `Orders` services.


![Scheme](https://iili.io/JLaxaUX.png)

### Run

Save **[config.example.json](configs/config.example.json)** as **config.json**:
```bash
cp configs/config.example.json configs/config.json
```

Run with docker:
```bash
docker compose up --build
```

Check health:
```bash
curl localhost:1337/api/health
```

### API Documentation

Swagger available at [localhost:8080](http://localhost:8080/)

### Profiling

Pprof available at [localhost:6060](http://localhost:6060/)

### Errors

To generate custom errors add their descriptions to the **[errors.json](configs/errors.json)** file and run:
```bash
make errors
```
