# Publish

* `docker login stor.highloadcup.ru`
* `docker build -t stor.highloadcup.ru/travels/curly_leopard .`
* `docker push stor.highloadcup.ru/travels/curly_leopard`

# Profile

* env `PROFILE_PATH=./profile`
* `go tool pprof -http=:6060 profile/cpu.pprof`