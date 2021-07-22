# Go Cambridge dictionary API

This is an unofficial cambridge dictionary API server implemented with Golang and also provide telegram bot feature with gitlab CI/CD pipeline on kubernetes.

## Get started !
```
git clone https://github.com/BWbwchen/Go_Cambridge_dictionary_API.git
cd Go_Cambridge_dictionary_API/
go run .
curl http://localhost:8080/api/dictionary 
```

## Dictionary API
```
GET localhost:8080/api/<word you what to search>
```

## Telegram bot API
```
POST localhost:8080/api/tg
```

## Combine with your gitlab CI/CD
1. Make sure your gitlab-runner is docker executor.
2. Make sure your kubernetes cluster have a deployment workload named `telegram-bot`
3. Fill some necessary information in [`.gitlab-ci.yml`](.gitlab-ci.yml) and [`telegram.go`](telegram.go)
4. Run on gitlab 😂 and set the environment variable `KUBERNETES_KUBE_CONFIG` in gitlab with your base64-encoded ~/.kube/config

### TODO
- [ ] Input filter
- [ ] Word pronunciation
