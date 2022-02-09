
## Intoduction
This application continuously pulls the events in the events collections at certain intervals and sends them to the relevant exchange on the data, then deletes the sent record if the sending operation is successful.


```
├── cmd
│   └── producer/*
│   └── main.go
├── configs
│   └── producer.yaml
├── internal
│   └── producer/*
├── pkg/*
└── main.go
```