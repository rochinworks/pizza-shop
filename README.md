# API
## GET /
returns:
```
{
  "alive": "true"
}
```

## POST /user/create
on success returns:
```
{
  "id": uuid
}
```

### example:
```
curl -X POST -H "Content-Type": "application/json" -d '{"username": "tests"}' <url>/user/create
```
### example return:
```
{
  "id": 97ee00a2-696a-4c4c-8d1d-233ce31f7eab
}
```

## POST /order/start
on success returns:
```
{
  "orderId": uuid
}
```

### example:
```
curl -X POST -H "Content-Type:application/json" -d '{"userId": "97ee00a2-696a-4c4c-8d1d-233ce31f7eab", "style": "hawaiian"}' <url>/order/start
```
### example return:
```
{
  "orderId": 6a0a9367-a2af-4ae8-8a3a-5c50e093f758,
  "status": "starting"
}
```

## GET /order/status
on success returns:
```
{
  "orderId": uuid,
  "status": string
}
```

### example:
```
`curl <url>/order/status?orderId=6a0a9367-a2af-4ae8-8a3a-5c50e093f758\&userId=97ee00a2-696a-4c4c-8d1d-233ce31f7eab
```

### example return:
```
{
  "orderId": 6a0a9367-a2af-4ae8-8a3a-5c50e093f758,
  "status": "starting"
}
```
