# Usage

```bash
go run ./cmd/api
curl "http://localhost:8080/sum?n=40"
```

## Build and push image
```
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -t devpies/fib:latest \
  --push .
```

## HPA Live Testing

```
curl "https://ex7.devpie.io/sum?n=40"
k top pods -n devpies
```