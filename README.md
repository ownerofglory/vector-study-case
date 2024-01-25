# Vector Case Study
![](https://github.com/ownerofglory/vector-study-case/actions/workflows/ecs-deploy-pipeline.yml/badge.svg)
![](https://github.com/ownerofglory/vector-study-case/actions/workflows/docker-build-pipeline.yml/badge.svg)
![](https://github.com/ownerofglory/vector-study-case/actions/workflows/test-pipeline.yml/badge.svg)


Following programming languages were considered for the implementation:
- __Java__
- __Kotlin__
- __Go__

### Java
Pros:
- Cross-platform (Runs on JVM)
- Has a variety of frameworks for web and microservice dev (Spring, SpringBoot, SpringCloud, Jakarta RS, etc)

Cons:
- Verbose
- Requires a lot of "boilerplate" code
- Requires runtime JVM (JRE), runtime dependencies thus:
  - bigger container image size, 
  - slower startup time, 
  - relatively slow*
  - memory overhead

### Kotlin
Pros:
- Compiles into JVM byte-code
- "Better", less verbose syntax comparing to Java
- Use whatever you'd use in Java and more (like KTor)

Cons:
- JVM variant same as Java with respect to memory and performance

### Go
Pros:
- Simplicity
- Great standard lib
- Built in server capabilities
- Fast and efficient
- Perfect cross-compiling model
- Easy to containerize, small image size

Cons:
- Runs natively* (no cross-platform, but due to great cross-compiling capability is still a great choice)

#### Result
> Go was chosen for the implementation

## Service architecture
> Hexagonal architecture was chosen as a common type of architecture in Go

![](./assets/hex-arch.png)


## Build and test
- Download dependencies
```shell
go mod download
```
- Test code with coverage
```shell
# Generate mocks
go generate ./...

# Run tests with coverage
go test -cover ./...
```
- Build
```shell
go build cmd/main.go
```
- Run locally:
```shell
  go run cmd/main.go
```
- Build Docker image and run container locally
```shell
## Build an image
docker build -t vector-study-case .

## Run container
docker run -p 80:80 vector-study-case:latest 
```



## Generate OpenAPI Spec
Prerequisites:
- go-swagger installed
```shell
 SWAGGER_GENERATE_EXTENSION=false swagger generate spec -o ./apidocs/swagger.json
```


## Deployment
### Deployment platform comparison

Requirements to the platform:
- Automated deployment
- Scheduling
- Scalability (or auto-scaling) -> incl. load balancing
- Preferably no infrastructure provisioning/management
- Service registry and discovery *
- On-demand option (considering the workload)

Options:
- __AWS Elastic Kubernetes Service (EKS)__
  - Managed k8s cluster
  - k8s capabilities: perfect for microservice apps
  - requires provisioned infrastructure (nodes, storage, networking)
- __AWS Elastic Container Service (ECS)__ EC2
  - Less admin overhead comparing to EKS
  - requires provisioned infrastructure (nodes, storage, networking)
  - For relatively simpler architectures
- __AWS Elastic Container Service (ECS) Fargate__ (Serverless)
  - on-demand, pay-as-you-go model
  - No provisioning and up-front costs
- __AWS Lambda (Serverless functions)__
  - on-demand, pay-as-you-go model
  - No admin overhead
  - No provisioning and up-front costs
  - For simple architectures*
  - Requires lambda handler

Result:
> AWS ECS Fargate option was chosen

### Architecture overview
Amazon ECS Fargate Architecture
![architecture](./assets/ecs-architecture.png)

#### Terms
- __ECS Fargate cluster__ - Scheduling, configuration
- __ECS Task__ - smallest computing unit, consists of one or more containers (like `Pod` in k8s)
- __ECS Task definition__ - task blueprint
- __ECS Service__ - (auto)-scaling, load balancing, service discovery

### AWS Deployment
- Build and push a docker image to Docker Hub
- Create an ECS Fargate cluster (serverless)
- Create an ECS Task definition
- Create an ECS service 
- Create a GitHub Actions Pipeline


### Deployment pipelines
- __Test__

  - Executes Go tests on each push

[Test pipeline definition](./.github/workflows/test-pipeline.yml)

Status: ![](https://github.com/ownerofglory/vector-study-case/actions/workflows/test-pipeline.yml/badge.svg)
- __Docker image build & push__
  - Runs on each merge into `master` branch
  - Builds a docker image and pushes into Docker Hub

[Docker push pipeline definition](./.github/workflows/docker-build-pipeline.yml)

Status: ![](https://github.com/ownerofglory/vector-study-case/actions/workflows/docker-build-pipeline.yml/badge.svg)

- __AWS ECS deployment__
  - Runs on a manual trigger
  - Rolls out the image from Docker Hub and run a container within an ECS service

[ECS deploy definition](./.github/workflows/ecs-deploy-pipeline.yml)

Status: ![](https://github.com/ownerofglory/vector-study-case/actions/workflows/ecs-deploy-pipeline.yml/badge.svg)

