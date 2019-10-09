# Publish

* `docker login stor.highloadcup.ru`
* `docker build -t stor.highloadcup.ru/travels/curly_leopard .`
* `docker push stor.highloadcup.ru/travels/curly_leopard`

# Profile

* env `PROFILE_PATH=./profile`
* `go tool pprof -http=:6060 profile/cpu.pprof`

# Load
* `cd cmd/load && go build && mv load* ../../ && cd ../..`
* `./load -target http://127.0.0.1 -ammo test/data/TRAIN/ammo/phase_1_get.ammo -load "line(1, 1000, 10s)"`