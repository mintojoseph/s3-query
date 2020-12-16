# s3-query

Query s3 buckets.

## Requirements

* Go
* Docker
* kubectl
* terraform
* Jenkins with following plugins- Docker Pipeline, Terraform, Kubernetes CLI. github need to be configured as 'scm'. Dockerhub credentials need to be configured as 'docker-hub-credentials'.
* minikube
* kubeseal

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

### Credentials

kubeseal is used to manage the credentials.

Encode the aws creds.

```cmdline
echo -n '<aws creds>' |base64
```

Update credentails in secret.yml file like below.

```yml
apiVersion: v1
kind: Secret
metadata:
  name: secret-basic-auth 
data:
  username: <aws_access_key_id>
  password: <aws_secret_access_key>
```

Create sealed credetails using kubeseal.

```cmdline
 kubeseal --format yaml <secret.yml >sealedsecret.yml
```

### Query the server

Server can be queried using follwouing syntax.

```cmdline
curl <hostname>:<port>/list?name=<bucket name>
```

```cmdline
$ curl 192.168.39.142:8080/list?name=mintos-test-bucket
Name:         sample.war
Storage class:STANDARD
Name:         testfile
Storage class:STANDARD

Found 2 items in bucket 
```

Use EXTERNAL-IP from following command as hostname.

```cmdline
 kubectl get svc
```

### Directories and files

* deployment/k8s/ - Kuberenetes yaml files for s3-query application.
* terraform/ - Terraform code for deploying an s3 bucket.
* main.go - main program.
* Jenkinsfile - To build and deploy.
* Dockerfile - To build images.
* Makefile - To build program.
