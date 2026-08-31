# DevOps Learning Project

This repository is a small educational example for learning CI/CD and
infrastructure automation around a simple web application.

The application itself is intentionally lightweight: a Vue frontend talks to a Go
HTTP API that serves fake dumpling store data. The main purpose of the project is
to demonstrate how application code can be packaged, built in CI, deployed with
Helm, and supported by cloud infrastructure provisioned with Terraform.

> This is training code. It is designed for learning and experimentation, not as
> a production-ready reference implementation.

## What Is Included

- Go backend API with health and metrics endpoints.
- Vue frontend served by Nginx.
- Dockerfiles for backend and frontend images.
- `docker-compose.yml` for local container-based startup.
- GitLab CI examples for triggering downstream backend, frontend, and
  infrastructure pipelines.
- Kaniko-based image build jobs.
- Helm chart for deploying frontend and backend to Kubernetes.
- Terraform configuration for creating a Yandex Cloud Managed Kubernetes
  cluster and node group.
- Extra Kubernetes manifests for ingress, cert-manager, and Nexus examples.

## Repository Structure

```text
.
|-- backend/              # Go API service
|-- frontend/             # Vue frontend application
|-- chart/                # Parent Helm chart with backend/frontend subcharts
|-- ci/                   # Reusable GitLab CI snippets
|-- infra/
|   |-- terraform/        # Yandex Cloud Kubernetes infrastructure
|   |-- certsmanager/     # cert-manager issuer example
|   `-- nexus/            # Nexus ingress example
|-- docker-compose.yml    # Local container runtime
`-- .gitlab-ci.yml        # Root pipeline with downstream triggers
```

## Application Overview

### Backend

The backend is a Go service built with:

- `chi` for HTTP routing.
- `zap` for logging.
- Prometheus client libraries for metrics.
- A fake in-memory dumplings store for demo data.

Main endpoints:

```text
GET  /health
GET  /metrics
GET  /api/products
GET  /api/categories
POST /api/orders
GET  /api/csrf
GET  /api/auth/whoami
```

The backend listens on port `8081`.

### Frontend

The frontend is a Vue application. It is built into static files and served by
Nginx. API requests are sent to `/api`, and the frontend Nginx config proxies
those requests to the backend service.

The frontend container listens on port `80`.

## Local Run

The easiest way to run the whole application locally is Docker Compose.

```bash
docker compose up --build
```

After startup:

- Frontend: `http://localhost`
- Backend health check: `http://localhost:8081/health`
- Backend metrics: `http://localhost:8081/metrics`

To stop the containers:

```bash
docker compose down
```

## Backend Development

From the backend directory:

```bash
cd backend
go test ./...
go run ./cmd/api
```

The API starts on `:8081`.

To build the backend Docker image manually:

```bash
docker build -t momo-store-backend ./backend
```

## Frontend Development

From the frontend directory:

```bash
cd frontend
npm install
npm run serve
```

Useful frontend commands:

```bash
npm run build
npm run typecheck
```

To build the frontend Docker image manually:

```bash
docker build -t momo-store-frontend ./frontend
```

## CI/CD Flow

The root `.gitlab-ci.yml` defines a parent pipeline and triggers downstream
pipelines based on changed paths:

- Changes in `backend/**/*` trigger the backend pipeline.
- Changes in `frontend/**/*` trigger the frontend pipeline.
- Changes in `infra/**/*` trigger the infrastructure pipeline.

The backend and frontend pipeline examples use Kaniko to build container images
inside GitLab CI without requiring a Docker daemon.

The image tags are based on GitLab metadata, for example:

```text
$CI_REGISTRY_IMAGE/momo-store-backend:$CI_COMMIT_SHA
$CI_REGISTRY_IMAGE/momo-store-frontend:$CI_COMMIT_SHA
```

The root pipeline also defines:

```yaml
VERSION: 0.0.${CI_PIPELINE_ID}
YC_REGISTRY: cr.yandex/$YC_REGISTRY_ID
```

## CI Variables

The CI snippets expect several variables to be available in GitLab CI/CD
settings. Exact values depend on your Yandex Cloud account and deployment
environment.

Common variables:

```text
YC_SERVICE_ACCOUNT_KEY_JSON_B64
YC_REGISTRY_ID
YC_CLOUD_ID
YC_FOLDER_ID
YC_CLUSTER_ID
YC_S3_ACCESS_KEY
YC_S3_SECRET_KEY
YC_NETWORK_ID
YC_SUBNET_ID
SSH_PRIVATE_KEY_B64
SSH_KNOWN_HOSTS
DEV_USER
DEV_HOST
```

## Terraform Infrastructure

Terraform configuration is stored in `infra/terraform`.

It demonstrates:

- Yandex Cloud provider configuration.
- S3-compatible remote Terraform state in Yandex Object Storage.
- Managed Kubernetes cluster creation.
- Kubernetes node group creation.
- Service account based authentication.

Typical Terraform workflow:

```bash
cd infra/terraform
terraform init
terraform plan
terraform apply
```

The configuration expects cloud, folder, network, subnet, service account, and
Object Storage credentials to be provided as Terraform variables or environment
variables.

## Helm Deployment

The Helm chart is stored in `chart`.

It contains a parent chart named `momo-store` and two subcharts:

- `momo-backend`
- `momo-frontend`

Default values are configured in `chart/values.yaml`.

Example commands:

```bash
helm dependency update ./chart
helm upgrade --install momo-store ./chart
```

The chart demonstrates:

- Kubernetes Deployments.
- ClusterIP Services.
- Image pull secrets.
- ConfigMaps and Secrets.
- Nginx config injection for the frontend.
- Prometheus scrape annotations for the backend.
- Ingress with TLS through cert-manager.

## Kubernetes Routing

The frontend Ingress example routes:

- `/` to the frontend service.
- `/api/` to the backend service.
You should add your own hosts!
Change this value before deploying to your own environment.

## Learning Notes

Because this project is meant for education, some configuration values are
hard-coded examples and should be adjusted before real deployment:

- Container registry repository paths.
- Domain names.
- Kubernetes cluster IDs.
- Yandex Cloud folder/cloud/network/subnet IDs.
- Secrets and service account credentials.
- Terraform remote state bucket and key.

Before using this repository in a real CI/CD environment, review pipeline include
paths, variable names, Helm values, and Terraform resource references for your
own setup.

## Suggested Practice Tasks

You can use this repository to practice:

- Building Docker images locally and in CI.
- Splitting GitLab pipelines by changed paths.
- Authenticating Kaniko to a container registry.
- Deploying a multi-service app with Helm.
- Managing Kubernetes ingress and TLS.
- Provisioning Kubernetes infrastructure with Terraform.
- Connecting application metrics to Prometheus scraping.

## Disclaimer

This repository is a compact DevOps learning sandbox. Security hardening,
production observability, database persistence, rollback strategy, autoscaling,
secret management, and full release promotion workflows are intentionally kept
minimal so the core CI/CD and infrastructure concepts stay easy to study.
