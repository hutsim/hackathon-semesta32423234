pipeline {
    agent any
        environment {
                PROJECT_ID = "semesta-390801"
                CLUSTER_NAME = "semesta-cluster"
                LOCATION = "asia-southeast1-a"
                CREDENTIALS_ID = "kubernetes"
        }

    stages {
            stage("Scm Checkout") {
                    steps {
                            checkout scm
                    }
            }

            stage("test semesta-app1"){
                steps {
                        sh "cd semesta-app1 && go test"

                }
            }

            stage("test semesta-app2"){
                steps {
                        sh "cd semesta-app2 && go test"

                }
            }
            
           
            stage("Build Docker Image semesta-app1") {
                    steps {
                            script {
                                 def imageName = "unvizy/s-app1:${env.BUILD_ID}"
                                 docker.build(imageName, "-f semesta-app1/Dockerfile .")
                            }
                    }
            }

            stage("Build Docker Image semesta-app2") {
                    steps {
                            script {
                                 def imageName = "unvizy/s-app2:${env.BUILD_ID}"
                                 docker.build(imageName, "-f semesta-app2/Dockerfile .")
                            }
                    }
            }
            
            stage("Push Docker Image semesta-app1") {
                    steps {
                            script {
                                    echo "Push Docker Image"
                                    withCredentials([string(credentialsId: "dockerhub", variable: "dockerhub")]) {
                                        sh "docker login -u unvizy -p ${dockerhub}"
                                    }
                                    def imageName = "unvizy/s-app1:${env.BUILD_ID}"
                                        docker.image(imageName).push()

                                    
                            }
                    }
            }

            stage("Push Docker Image semesta-app2") {
                    steps {
                            script {
                                    echo "Push Docker Image"
                                    withCredentials([string(credentialsId: "dockerhub", variable: "dockerhub")]) {
                                        sh "docker login -u unvizy -p ${dockerhub}"
                                    }
                                    def imageName = "unvizy/s-app2:${env.BUILD_ID}"
                                        docker.image(imageName).push()

                                    
                            }
                    }
            }
            
            stage("Deploy semesta-app1 to K8s") {
                    steps{
                            echo "Deployment started ..."
                            sh "ls -ltr"
                            sh "pwd"
                            sh "sed -i 's/tagversion/${env.BUILD_ID}/g' semesta-app1/deployment.yaml"
                            echo "Start deployment of deployment.yaml"
                            step([$class: "KubernetesEngineBuilder", projectId: env.PROJECT_ID, clusterName: env.CLUSTER_NAME, location: env.LOCATION, manifestPattern: "semesta-app1/deployment.yaml", credentialsId: env.CREDENTIALS_ID])
                            echo "Deployment Finished ..."
                    }
            }


            stage("Deploy semesta-app2 to K8s") {
                    steps{
                            echo "Deployment started ..."
                            sh "ls -ltr"
                            sh "pwd"
                            sh "sed -i 's/tagversion/${env.BUILD_ID}/g' semesta-app2/deployment.yaml"
                            echo "Start deployment of deployment.yaml"
                            step([$class: "KubernetesEngineBuilder", projectId: env.PROJECT_ID, clusterName: env.CLUSTER_NAME, location: env.LOCATION, manifestPattern: "semesta-app2/deployment.yaml", credentialsId: env.CREDENTIALS_ID])
                            echo "Deployment Finished ..."
                    }
            }
    }
}
