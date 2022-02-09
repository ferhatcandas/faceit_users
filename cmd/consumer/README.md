
## Intoduction
This application runs 3 different rabbit consumers and they consume the incoming messages by connecting to the relevant exchanges.

These messages are User:Created, User:Updated, User:Deleted


```
├── cmd
│   └── consumer/* 
│   └── main.go
├── configs
│   └── consumer.yaml
├── internal
│   └── consumer/*
├── pkg/*
└── main.go
```