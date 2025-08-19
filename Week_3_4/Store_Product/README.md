.
├── README.MD
├── src
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── main.go
└── terraform
    ├── main.tf
    ├── modules
    │   ├── ecr
    │   │   ├── main.tf
    │   │   ├── outputs.tf
    │   │   └── variables.tf
    │   ├── ecs
    │   │   ├── main.tf
    │   │   ├── outputs.tf
    │   │   └── variables.tf
    │   ├── logging
    │   │   ├── main.tf
    │   │   ├── outputs.tf
    │   │   └── variables.tf
    │   └── network
    │       ├── main.tf
    │       ├── outputs.tf
    │       └── variables.tf
    ├── outputs.tf
    ├── provider.tf
    └── variables.tf