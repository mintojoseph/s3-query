# s3-query
Query s3 buckets

## Requirements

* Go
* Docker
* kubectl
* terraform
* Jenkins with following plugins- Docker Pipeline, Terraform, Kubernetes CLI. github need to be configured as 'scm'. Dockerhub credentials need to be configured as 'docker-hub-credentials'.
* minikube

### Available parameters

``` cmdline
./s3-query --help
Usage of ./s3-query:
  -port string
    	HTTP port. HTTPPORT environment variable can also be used. (default "8080")
```

Port can be set using HTTPPORT environment variable as well. Parameter takes precedence.

### Example

Run the application.

 ``` cmdline
 ./s3-query -port 8083
```

Docker Build

``` cmdline
sudo docker build  -t  mintojoseph/s3-query:1.0 .
sudo docker container run -p 8080:8080 mintojoseph/s3-query:1.0
```

### Query the server

Server can be queried using follwouing syntax.

```cmdline
$ curl <hostname>:<port>/list?name=<bucket name>
```

```cmdline
$ curl 192.168.39.142:8080/list?name=mintos-test-bucket
Name:         sample.war
Storage class:STANDARD
Name:         testfile
Storage class:STANDARD

Found 2 items in bucket 
```

### Directories and files

* deployment/k8s/ - Kuberenetes yaml files for s3-query application.
* terraform/ - Terraform code for deploying an s3 bucket.
* main.go - main program.
* Jenkinsfile - To build and deploy.
* Dockerfile - To build images.
* Makefile - To build program.
