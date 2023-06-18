# icetrap

`icetrap` is a custom bingo card service

## Requirements

The only requirements to run `icetrap` are docker and docker-compose.

## Usage

Start the development server (accessible on port 8000):
```
docker-compose up --build
```

Run tests:
```
docker-compose run backend go test
docker-compose run frontend npm run test
```

Database console:
```
docker-compose run --rm postgres psql -h postgres -U postgres
```

Run database migrations:
```
docker-compose run backend sh migrate.sh
```
