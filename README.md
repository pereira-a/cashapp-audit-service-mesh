# Cashapp with audit project (R&D)

Goals:
- Learn basics of Go
- Investigate a Java framework (TBD quarkus, micronaut)
- Trains and investigate more Istio service mesh
- Have an heterogenous "microservice environment" deployed in k8s (minikube)
- Install istio and orchestrate a service mesh, use side cars to:
  - apply trace logging
  - service discovery
  - traffic management
- Investigate flagger.app

## MS1 (Golab)

Audit service. 2 service API:
- POST to register an action
- GET to fetch actions by userid


Maybe, also consume kafka event to register action (later to investigate kafka integration in Go)

## MS2 (Java)

Cashapp service, with following services:
- Register client
- Send money to another client
- Consult movements

Audits operations made by users in MS1.