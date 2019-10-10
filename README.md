# Publish

* `docker login stor.highloadcup.ru`
* `docker build --rm -t stor.highloadcup.ru/travels/curly_leopard .`
* `docker push stor.highloadcup.ru/travels/curly_leopard`

# Profile

* env `PROFILE_PATH=./profile`
* `go tool pprof -http=:6060 profile/cpu.pprof`

# Load test
* `cd cmd/load && go build && mv load* ../../ && cd ../..`
* `./load -target http://127.0.0.1 -ammo test/data/TRAIN/ammo/phase_1_get.ammo -load "line(1, 1000, 10s)"`
* `cd test/data && docker build -t local/hlcup2017-data . && cd ../..`

## Docker machine for Windows (cygwin)
* `docker-machine.exe create --driver virtualbox --virtualbox-share-folder "d:\\full\\path\\to\\hlcup2017-travel:hlcup2017-travel" --virtualbox-cpu-count "2" --virtualbox-disk-size "20000" --virtualbox-memory "6144" hlcup2017-travel`
* `docker-compose up --build`
* `go tool pprof -source_path=./workspace -http=:6060 profile/cpu.pprof`