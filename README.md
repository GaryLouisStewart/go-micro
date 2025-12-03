# 

```bash
# file structure
├── .github
│   └── workflows
│       └── build-and-test.yaml # Github Actions
├── .pre-commit-config.yaml     # Pre-commit hooks
├── cmd
│   └── api
│       └── main.go             # Entry point
├── Dockerfile
├── go.mod
├── internal
│   ├── handlers
│   │   ├── handlers_test.go    # Business logic
│   │   └── handlers.go         # Tests
│   └── middleware
│       └── logger.go           # Logger setup
├── Makefile
└── README.md
```
