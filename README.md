# go-micro

A small Go microservice demo deployed to Google Cloud Run via GitHub Actions
(Workload Identity Federation) and Cloud Build.

## File structure

```text
├── .github
│   └── workflows
│       └── build-and-test.yaml # GitHub Actions (security scan, test, build, deploy)
├── .pre-commit-config.yaml     # Pre-commit hooks
├── cmd
│   └── api
│       └── main.go             # Entry point
├── Dockerfile                  # Multi-stage build on distroless
├── go.mod
├── internal
│   ├── handlers
│   │   ├── handlers.go         # Business logic (hello / health)
│   │   └── handlers_test.go    # Tests
│   └── middleware
│       └── logger.go           # Logger + status-recorder middleware
├── Makefile
├── cloudbuild.yaml             # Cloud Build pipeline
└── README.md
```

## Endpoints

| Method | Path     | Description                          |
| ------ | -------- | ------------------------------------ |
| GET    | `/hello` | Returns a JSON greeting + timestamp. |
| GET    | `/health`| Returns `OK` (200) for liveness.     |

## Deployment flow

`main` -> production Cloud Run service (`go-micro-prod-service`) using a
blue/green rollout: a new revision is deployed with `--no-traffic --tag green`,
smoke-tested via the tagged URL, then traffic is migrated with
`update-traffic --to-tags green=100`.

`dev` -> dev Cloud Run service (`go-micro-dev-service`) tagged with the
commit SHA.
