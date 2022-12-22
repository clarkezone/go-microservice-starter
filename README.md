# Go-microservice-starter

[![License](https://img.shields.io/github/license/clarkezone/go-microservice-starter.svg)](https://github.com/clarkezone/go-microservice-starter/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/clarkezone/go-microservice-starter)](https://goreportcard.com/report/github.com/clarkezone/go-microservice-starter)
[![Build and Tests](https://github.com/clarkezone/go-microservice-starter/workflows/run%20tests/badge.svg)](https://github.com/clarkezone/go-microservice-starter/actions?query=workflow%3A%22run+tests%22) [![Coverage Status](https://coveralls.io/repos/github/clarkezone/go-microservice-starter/badge.svg?branch=main)](https://coveralls.io/github/clarkezone/go-microservice-starter?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/clarkezone/go-microservice-starter.svg)](https://pkg.go.dev/github.com/clarkezone/go-microservice-starter)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/6231/badge)](https://bestpractices.coreinfrastructure.org/projects/6231)
[![GitHub release](https://img.shields.io/github/release/clarkezone/go-microservice-starter.svg?style=flat-square)](https://github.com/clarkezone/go-microservice-starter/releases)
![Total Downloads](https://img.shields.io/github/downloads/clarkezone/go-microservice-starter/total?logo=github&logoColor=white)

# project state

Bootstrapping infra:

- [x] update k8s files to point to correct image
- [ ] create 0.0.1 release
- [ ] Dashboard for telemetry verified with loadtester
- [ ] update loadtester
- [ ] verify microservice mode locally and add commands to test to readme
- [ ] k8s local override with staging / prod
- [ ] Instructions for prerequs
- [ ] Turn on PR enforcement, protect main branch

Template backlog

- factor out common concerns
- Add k8s manifest scanner for best practices (PDB, CPU/MEM requests and limits)
- Add DT
- Add vscode devcontaienr
- complete openssf best practices
- Add minimal viable governance

# Creating a new release

```bash
git tag -a v0.0.1 -m "helloinfra"
git push origin v0.0.1
gh release create
```
