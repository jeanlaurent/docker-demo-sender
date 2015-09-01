# sender

## Build
need go >= 1.5

```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sender .
docker build -t jeanlaurent/sender .
```

## Run
`docker run -ti jeanlaurent/sender /sender $(docker-machine ip default) 9000`

## Run in mayhem mode

`for i in {1..5}; do ( docker run jeanlaurent/sender /sender $(docker-machine ip default) 9000 &);  done; ``
