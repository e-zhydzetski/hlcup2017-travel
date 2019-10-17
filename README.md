# Publish

* `docker login stor.highloadcup.ru`
* `docker build --rm -t stor.highloadcup.ru/travels/curly_leopard .`
* `docker push stor.highloadcup.ru/travels/curly_leopard`

# Profile

* env `PROFILE_PATH=./profile`
* `go tool pprof -http=:6060 profile/cpu.pprof`

## Docker machine for Windows (cygwin)
* `docker-machine.exe create --driver virtualbox --virtualbox-share-folder "d:\\full\\path\\to\\hlcup2017-travel:hlcup2017-travel" --virtualbox-cpu-count "3" --virtualbox-disk-size "20000" --virtualbox-memory "6144" hlcup2017-travel`
* `docker-compose up --build`
* `go tool pprof -source_path=./workspace -http=:6060 profile/cpu.pprof`