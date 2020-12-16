node {
    def app

    stage('Clone repository') {

        checkout scm
    }

    stage('Build image') {

        app = docker.build("mintojoseph/s3-query")
    }

    stage('Push image') {
        docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
            app.push("${env.BUILD_NUMBER}")
            app.push("latest")
        }
    }
 
    stage('k8s secrets') {
   
       sh "cd deployment/k8s; kubectl apply -f secrets.yml"

    }

    stage('k8s deployment') {
   
       sh "cd deployment/k8s; kubectl apply -f deployment.yml"

    }

    stage('k8s metallb') {

       sh "cd deployment/k8s; kubectl apply -f metallb.yml"

    }

    stage('k8s service') {

       sh "cd deployment/k8s; kubectl apply -f service.yml"

    }
}