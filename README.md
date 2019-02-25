# GO Social Networking

This repository is used as a POC for a large scale social networking API written on top of GO.

## How to run locally

To run locally, you need to be in a linux/mac machine and have Docker and Docker Compose installed on your machine.

```sh
docker-compose up
```

## Abstract

- Showcase how to build an application that is capable of scaling to adapt to traffic generated by millions of users.
- Showcase how CQRS prevents optimistic locking and improve read operation performances.
- Showcase how event sourcing introduces auditability and higher scalability of a system.
- Showcase how to implement graceful error management that protects application from 3rd party downtimes.
- Showcase how to recover from projection directed attacks and disasters.

## Architecture Diagram

## Technology Stack

This section describes the technology stack composing the application.

### GO as core language

GO is be picked due to its capability to handle larger number of concurrent and parallel requests over NodeJS. This capability originates from its compiled nature (Meaning there is no runtime penalty / cost associated with the interpretation process of code).

### Janus as API Gateway

An API gateway will be used to aggregate downstream services (Activity Write API + Activity Read APIs) and provide the following cross cutting concerns:

- Authentication + Authorizations
- Throughput Limitations (Rate Limiting)
- Single point of contact
- Single requirement for domain and ssl certificate.
- Domain abstraction
- SSL terminations
- Distributed tracings
- Resiliency through the use of circuit breakers and retries.
- Centralized CORS management.

Janus API Gateway is selected because of the following:

- Runs of top of single GO binary
- No Dependency Hell
- Ease of API Configuration
- Hot reloading support for better development UX
- HTTP2 Support
- Awesome offical docker image

### Gin Gonic as Downstream Service APIs

A downstream service in a microservice ecosystem refers to an API that is encapsulated inside your infrastructure and only accessible from an API Gateway.

Gin Gonic is selected as the core web framework because of its superior performance over it's comparables (Revel, Gorilla) and predecessor (Martini). Aside from its speed, it is less opinionated and minimalist by design which makes it lightweight and more performant compared to its competitors.

### Write and Read Separations (CQRS)

The read and write APIs will is broken to separate deployments to allow more efficient use resources. This is motivated by the fact that the ratio of posting in social media is **1 write vs 500-1000 reads**. An example that validates this claim is that an average middle-aged person would have an average of 500-1000 friends with at least 2 devices (Laptop + Mobile Phone) causing a single activity to have higher read demand and lower write demand.

### Different Write and Read Stores (Event Stores + Projections)

The high demand for read over write nature of a social media network can directly impact how could a single write affect response time for users attempting to pull data from data storage in the form of optimistic locking.

To solve the mentioned problem, the project will be using a hash-based read-dedicated document stores (Redis + MongoDB) for pulling data. This read-dedicated stores will be called **projections** as they are the projected current state of all events that we're successfully registered in the system.

In order for us to build projections (Read-stores), we will need a data store to write data that enters the system, we will call this as an **event store** (PostgreSQL), the event store will be in-charge of recording changes in the state of data (activities / events) and wil be dedicated for writing changes to the system.

### Offloading of write operations via queues.

To enable faster response rates and higher HTTP traffic processing, write operations will be offloaded to background jobs via queueing technology. This approach gives the following advantages:

- Faster API response time since the write doesn't need to reflect quickly.
- Resiliency against write store downtimes.
- Graceful error management by either secondary event store or error store.

### Graceful Error Management

Anticipating downtimes associated with issues, deployments and 3rd party failures, the system is designed to gracefully manage any 3rd party downtimes. The following technologies are introduced to manage failures in the system:

- Redundant Queues
- Slaves Projections (Activate in-case of master projection downtime)
- Backup Event Store
- Error Logging
- Error Queues
- Circuit Breaking
- Re-queueing of Failed Messages

### Containerization and Orchestration

To ease package management, each technical service would be containerized which gives the following benefits:

- Isolation of Application
- Controlled Environment
- Enables Usage of Different Dependencies
- Faster Bootup Time Over VMs

Since this application is designed to handle millions of users, container ochestration makes more sense to be included in this stack. Container orchestrators like Kubernetes enable CPU-based automatic scaling which enables the system to scale up accordingly to current volume of HTTP traffic.
