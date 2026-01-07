# 🚀 DevOps Portfolio Website

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.25-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-v2.52-00ACD7?style=for-the-badge&logo=go&logoColor=white)
![CI/CD](https://img.shields.io/badge/CI%2FCD-GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

**A high-performance, portfolio website showcasing modern DevOps practices**

[Live Demo](#) • [Features](#-features) • [Tech Stack](#-tech-stack) • [CI/CD Pipeline](#-cicd-pipeline) • [Getting Started](#-getting-started)

</div>

---

## 📋 Overview

This project is a **fully automated, portfolio website** built with Go and the Fiber framework. It serves as both a personal portfolio and a demonstration of industry-standard DevOps practices including:

- **Infrastructure as Code (IaC)**
- **Continuous Integration & Continuous Deployment (CI/CD)**
- **Automated Testing & Code Quality Gates**
- **Secure Deployment via SSH**
- **Observability Stack** *(Prometheus/Grafana - Coming Soon)*

> 💡 **This project is a living demonstration of my skills** – every commit triggers a full CI/CD pipeline that builds, tests, and deploys to production (if passing quality gates, of course..).

---

## ✨ Features

| Feature | Description |
|---------|-------------|
| ⚡ **High Performance** | Built with Go Fiber – one of the fastest web frameworks |
| 🔄 **Automated CI/CD** | GitHub Actions pipeline with lint, test, build & deploy stages |
| 📊 **Code Coverage** | Automated test coverage reporting with PR comments |
| 🔒 **Secure Deployment** | SSH-based deployment with encrypted credentials |
| 📈 **Observability Ready** | Prometheus metrics & Grafana dashboards *(Roadmap)* |
| 🏗️ **Build Traceability** | Git SHA & branch injected into binary at build time |
| 🚦 **Concurrency Control** | Smart pipeline management with auto-cancellation |

---

## 🛠 Tech Stack

### Backend
| Technology | Purpose |
|------------|---------|
| ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white) | Core application language |
| ![Fiber](https://img.shields.io/badge/Fiber-00ACD7?style=flat-square&logo=go&logoColor=white) | High-performance web framework |

### DevOps & Infrastructure
| Technology | Purpose |
|------------|---------|
| ![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-2088FF?style=flat-square&logo=github-actions&logoColor=white) | CI/CD Pipeline Orchestration |
| ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat-square&logo=docker&logoColor=white) | Containerization *(Roadmap)* |
| ![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?style=flat-square&logo=prometheus&logoColor=white) | Metrics Collection *(Roadmap)* |
| ![Grafana](https://img.shields.io/badge/Grafana-F46800?style=flat-square&logo=grafana&logoColor=white) | Monitoring & Dashboards *(Roadmap)* |

### Code Quality
| Tool | Purpose |
|------|---------|
| ![golangci-lint](https://img.shields.io/badge/golangci--lint-00ADD8?style=flat-square&logo=go&logoColor=white) | Static code analysis & linting |
| ![Go Test](https://img.shields.io/badge/Go_Test-00ADD8?style=flat-square&logo=go&logoColor=white) | Unit testing with race detection |

---

## 🔄 CI/CD Pipeline

This project implements a **robust, multi-stage CI/CD pipeline** using GitHub Actions:

```
┌─────────────────────────────────────────────────────────────────┐
│                        CI/CD Pipeline                           │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│   ┌─────────┐    ┌─────────┐    ┌─────────┐    ┌─────────────┐  │
│   │  Lint   │───▶│  Test   │───▶│  Build  │───▶│   Deploy    │  │
│   │         │    │         │    │         │    │             │  │
│   │golangci │    │ go test │    │ go build│    │ SSH Deploy  │  │
│   │  -lint  │    │ -race   │    │ -ldflags│    │ (main only) │  │
│   └─────────┘    └─────────┘    └─────────┘    └─────────────┘  │
│                        │                              │         │
│                        ▼                              │         │
│                  ┌──────────┐                         │         │
│                  │Coverage  │                         │         │
│                  │ Report   │                         │         │
│                  └──────────┘                         │         │
│                        │                              │         │
│                        ▼                              │         │
│                  ┌──────────┐     On PR Only          │         │
│                  │PR Comment│◀────────────────────────┘         │
│                  │(Summary) │                                   │
│                  └──────────┘                                   │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```
---
### Pipeline Highlights

| Stage | Description                                    |
|-------|------------------------------------------------|
| **🔍 Lint** | Static code analysis using `golangci-lint`     |
| **🧪 Test** | Runs unit tests with race condition detection & generates coverage |
| **📦 Build** | Compiles binary with Git metadata for traceability |
| **💬 PR Comment** | Automatically comments PR test results & coverage |
| **🚀 Deploy** | Secure SSH deployment to production            |

### Advanced Pipeline Features

- **Concurrency Control**: Prevents parallel runs on the same branch, auto-cancels stale jobs
- **Custom Composite Actions**: Reusable deployment action in `.github/actions/deploy/`
- **Environment Protection**: Production deployments require environment approval
- **Security hardening**: Using commit SHA as action version, no secrets exposed, following best practices
---

## 🗺️ Roadmap

| Feature | Status      |
|---------|-------------|
| 🔄 Go Fiber Web Server | In Progress |
| ✅ CI/CD Pipeline with GitHub Actions | Complete    |
| ✅ Automated Testing & Coverage | Complete    |
| ✅ Secure SSH Deployment | Complete    |
| ✅ Build Traceability (Git SHA injection) | Complete    |
| 🔄 Docker Containerization | In Progress |
| 📋 Prometheus Metrics Integration | Planned     |
| 📋 Grafana Monitoring Dashboard | Planned     |
| 📋 Log Aggregation (ELK/Loki) | Planned     |
| 📋 Kubernetes Deployment Manifests | Planned     |
| 📋 Terraform Infrastructure | Planned     |

---

## 🎯 Skills Demonstrated

This project showcases my proficiency in:

- **CI/CD**: GitHub Actions, multi-stage pipelines, artifact management.
- **Infrastructure as Code**: Reusable composite actions, YAML-based configuration.
- **Programming**: Go, clean code architecture, go idiomatic approach.
- **Testing**: Unit testing, race condition detection, code coverage analysis.
- **Security**: SSH-based deployments, secret management, minimal permissions.
- **Monitoring**: Prometheus/Grafana integration *(in progress)*.
- **Best Practices**: Security hardening, automation and code review workflows.

---

## 📫 Contact

Feel free to reach out if you'd like to discuss job opportunities!

- **Portfolio**: [Site](#)
- **LinkedIn**: [LinkedIn](#)
- **Email**: serpens_deus@outlook.com

---

<div align="center">

**Built with ❤️ and a passion for automation**

*This project is continuously improving – check back for updates!*

</div>

