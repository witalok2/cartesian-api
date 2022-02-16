## CARTESIAN-API
### How to run
```bash
first run 'make run-ecosystem' to run the ecosystem

then 'make run' or 'go run main.go -port=3005 -debug=true' to run the application
```

```bash
"args": ["-port=9000", "-debug=true"]
```

## Technology
```bash
GoLang
Echo
Docker
Goose
Postgres
```
Endpoint:
    
    GET {{URL}}api/v1/points?x={number}&y={number}&distance={number}
    exemple:
    GET localhost/api/v1/points?x=-3&y=4&distance=20   
