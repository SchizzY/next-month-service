# build the go binary
build:
  go build .

# run the go binary
run:
  go run .

# clean the go binary
clean:
  go clean .

# start project with docker compose
up:
  docker compose up -d

# stop project with docker compose
down:
  docker compose down

# start project with docker compose and watch for changes
watch:
  docker compose watch
