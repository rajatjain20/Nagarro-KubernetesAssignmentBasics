## ***** Kubernetes Basics Assignment *****

# Overview
 I have created one helloworld app in golang. It is just one static page.  
 I have setup kubernetes and minikube on my windows machine. (refer `Installation_kubernetes.docx`)  

**Below is the directory structure**  

 `./app/`                 - it contains helloworld application with its Dockerfile.  
 `./Deployment.yaml`      - this is deployment yaml file to deploy application on kubernetes cluster.  
 `./Service.yaml`         - this is service yaml file to expose the application to excess it externally (outside kubernetes).  
 `./Deployment-Bonus.yaml`  - this is deployment yaml file for bonus question having liveness and readiness probes configured.  
`./Installation_kubernetes.docx` - Steps to install Kubernetes on windows machine.  

# Create Docker image inside kubernetes cluster (I am using kubernetes on my local machine. we can use image from dockerhub as well)

**Below commands are not necessary to run as I have pushed docker image to docker hub. We can skip this part and jump to next step - "Deployment of application in Kubernetes Cluster"**  
Command to set windows poweshell to use Minikube’s Docker daemon

    > & minikube -p minikube docker-env | Invoke-Expression

    now check available images inside kubernetes  

    > docker images

   goto dir "/Kubernetes_Assignment_Basics/app/" __replace path with complete path on your local__  

    > cd "/Kubernetes_Assignment_Basics/app/"   

    command to build image inside kubernetes  

    > docker build -t k8sassignmentbasics:v0 .  

    > docker images  

    it should show image with name "k8sassignmentbasics:v0"

# Deployment of application in Kubernetes Cluster
I have created a yaml file named "Deployment.yaml", placed inside "./Kubernetes_Assignment_Basics/"

- Command to deploy:  
    goto dir "/Kubernetes_Assignment_Basics/" __replace path with complete path on your local__
    > cd "/Kubernetes_Assignment_Basics/"

    > kubectl create -f ./Deployment.yaml --save-config  

    deployment.apps/k8s-basics-assignment created    

- To check deployment:
    > kubectl get all  

NAME                                         READY   STATUS    RESTARTS   AGE  
pod/k8s-basics-assignment-859b784cb7-w2kk2   1/1     Running   0          33s  

NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE  
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   3d1h  

NAME                                    READY   UP-TO-DATE   AVAILABLE   AGE  
deployment.apps/k8s-basics-assignment   1/1     1            1           33s  

NAME                                               DESIRED   CURRENT   READY   AGE  
replicaset.apps/k8s-basics-assignment-859b784cb7   1         1         1       33s  

# Expose the application using a Kubernetes Service to access it externally.  
I have created a yaml file named "Service.yaml", placed inside directory "./Kubernetes_Assignment_Basics/"  

- Command to create service to expose the application:  
    goto dir `/Kubernetes_Assignment_Basics/` __replace path with complete path on your local__
    > cd "./Kubernetes_Assignment_Basics/"

    >  kubectl create -f ./Service.yaml --save-config  

    service/k8s-basics-assignment-service created

- To Check created service:
    >kubectl get services  

NAME                            TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE  
`k8s-basics-assignment-service   NodePort    10.109.137.31   <none>        3010:30036/TCP   79s`  
kubernetes                      ClusterIP   10.96.0.1       <none>        443/TCP          3d3h  

- To run exposed application locally:
    > minikube service k8s-basics-assignment-service  

|-----------|-------------------------------|-------------|---------------------------|  
| NAMESPACE |             NAME              | TARGET PORT |            URL            |  
|-----------|-------------------------------|-------------|---------------------------|  
| default   | k8s-basics-assignment-service |        3010 | http://192.168.49.2:30036 |  
|-----------|-------------------------------|-------------|---------------------------|  
🏃  Starting tunnel for service k8s-basics-assignment-service.  
|-----------|-------------------------------|-------------|------------------------|  
| NAMESPACE |             NAME              | TARGET PORT |          URL           |  
|-----------|-------------------------------|-------------|------------------------|  
| default   | k8s-basics-assignment-service |             | http://127.0.0.1:52925 |  
|-----------|-------------------------------|-------------|------------------------|  

🎉  Opening service default/k8s-basics-assignment-service in default browser...  
❗  Because you are using a Docker driver on windows, the terminal needs to be open to run it.  

**Note:** This will execute the application on browser.  

# Scale the application by increasing the number of replicas

   goto dir `/Kubernetes_Assignment_Basics/` __replace path with complete path on your local__  
   
    > cd "./Kubernetes_Assignment_Basics/"  

- Scale the application:  

    * Method 1:  
        > kubectl scale deployment k8s-basics-assignment --replicas=3  

        deployment.apps/k8s-basics-assignment scaled

    * Method 2:  
        We can change the count from 1 to any number (max 110) in Deployment.yaml file and can execute below command:  

    > kubectl apply -f ./Deployment.yaml  

    deployment.apps/k8s-basics-assignment configured

- To verify scaled application:  
    > kubectl get deployment k8s-basics-assignment  

NAME                    READY   UP-TO-DATE   AVAILABLE   AGE  
k8s-basics-assignment   5/5     5            5           31m  

- To check pods:
    > kubectl get pods  

NAME                                    READY   STATUS    RESTARTS   AGE  
k8s-basics-assignment-898db5488-7htxt   1/1     Running   0          2m15s  
k8s-basics-assignment-898db5488-kvsng   1/1     Running   0          2m15s  
k8s-basics-assignment-898db5488-n989j   1/1     Running   0          2m15s  
k8s-basics-assignment-898db5488-n9jh7   1/1     Running   0          31m  
k8s-basics-assignment-898db5488-wfpkm   1/1     Running   0          2m15s  


# Bonus Question -

# Enhance the application deployment by adding a health check endpoint and configuring liveness and readiness probes in the Kubernetes deployment.

I have added one health check endpoint (**"/healthz"**) in my source code (image tag is updated to v1) and added liveness and readiness probes in Deployment-Bonus.yaml.  
This endpoint returns error code (500) after 30secs of the server startup and keeps returning it till 50secs. This error code (500) is considered as its failure state by kubernetes liveness probe and readiness probe.  

- Let's deploy this deployment:
    > kubectl create -f ./Deployment-Bonus.yaml --save-config  

    deployment.apps/k8s-basics-assignment-bonus created  

    > kubectl get deployment k8s-basics-assignment-bonus  

NAME                          READY   UP-TO-DATE   AVAILABLE   AGE  
k8s-basics-assignment-bonus   1/1     1            1           16s  

- Expose this to a service:
    > kubectl expose deployment k8s-basics-assignment-bonus --name=k8s-basics-assignment-bonus-service --port=3020 --target-port=3000 --type=NodePort  

    service/k8s-basics-assignment-bonus-service exposed  

    > kubectl get service k8s-basics-assignment-bonus-service  

NAME                                  TYPE       CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE  
k8s-basics-assignment-bonus-service   NodePort   10.105.184.6   <none>        3020:30278/TCP   13s  

    > kubectl get endpoints k8s-basics-assignment-bonus-service -o wide  

NAME                                  `ENDPOINTS`          AGE  
k8s-basics-assignment-bonus-service   `10.244.0.64:3000`   31s  

- Wait for ~50secs or more and check the resources:  
    > kubectl get all  

NAME                                               READY   STATUS    RESTARTS      AGE  
pod/k8s-basics-assignment-898db5488-wvcw9           1/1     Running      0           38h  
`pod/k8s-basics-assignment-bonus-7f96c49df4-kcwz9   1/1     Running   1 (38s ago)   93s`  

NAME                                          TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE  
`service/k8s-basics-assignment-bonus-service  NodePort    10.105.184.6   <none>        3020:30278/TCP   53s`  
service/k8s-basics-assignment-service         NodePort    10.110.49.91   <none>        3010:30036/TCP   38h  
service/kubernetes                            ClusterIP   10.96.0.1      <none>        443/TCP          9d  

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE  
deployment.apps/k8s-basics-assignment         1/1     1            1           38h  
`deployment.apps/k8s-basics-assignment-bonus   1/1     1            1           93s`  

NAME                                                     DESIRED   CURRENT   READY   AGE  
replicaset.apps/k8s-basics-assignment-898db5488          1         1         1       38h  
`replicaset.apps/k8s-basics-assignment-bonus-7f96c49df4   1         1         1       93s`  

- Let's check the events in pod (pick the pod name from above command)**  
    > kubectl describe pod k8s-basics-assignment-bonus-7f96c49df4-kcwz9  

Events:  
  Type     Reason     Age                 From               Message  
  ----     ------     ----                ----               -------  
  Normal   Scheduled  2m19s               default-scheduler  Successfully assigned default/k8s-basics-assignment-bonus-7f96c49df4-kcwz9 to minikube  
  Normal   Pulling    2m19s               kubelet            Pulling image "rajatjain20/k8sassignmentbasics:v1"  
  Normal   Pulled     2m8s                kubelet            Successfully pulled image "rajatjain20/k8sassignmentbasics:v1" in 10.778s (10.778s including waiting). Image size: 918814242 bytes.  
  Normal   Created    39s (x3 over 2m7s)  kubelet            Created container: kubernetes-assignment-bonus  
  Normal   Started    39s (x3 over 2m7s)  kubelet            Started container kubernetes-assignment-bonus  
  `Normal   Killing    39s (x2 over 84s)   kubelet            Container kubernetes-assignment-bonus failed liveness probe, will be restarted`  
  Normal   Pulled     39s (x2 over 84s)   kubelet            Container image "rajatjain20/k8sassignmentbasics:v1" already present on machine  
  `Warning  Unhealthy  4s (x7 over 94s)    kubelet            Liveness probe failed: HTTP probe failed with statuscode: 500`  
  `Warning  Unhealthy  3s (x12 over 95s)   kubelet            Readiness probe failed: HTTP probe failed with statuscode: 500`  


- Check logs on the pod:
    > kubectl logs pod/k8s-basics-assignment-bonus-7f96c49df4-kcwz9

    __pod name can be fetched from `kubectl get pods` command__  

Hello World application  
Listning at port :3000  
2025/04/21 12:03:57 getHealth() -> Success: 4.694275966 secs elapsed since server started.  
2025/04/21 12:03:59 getHealth() -> Success: 6.542204759 secs elapsed since server started.  
2025/04/21 12:04:02 getHealth() -> Success: 9.692739145000001 secs elapsed since server started.  
2025/04/21 12:04:04 getHealth() -> Success: 11.542168851 secs elapsed since server started.  
2025/04/21 12:04:07 getHealth() -> Success: 14.692214375 secs elapsed since server started.  
2025/04/21 12:04:09 getHealth() -> Success: 16.541696179 secs elapsed since server started.  
2025/04/21 12:04:12 getHealth() -> Success: 19.692441263 secs elapsed since server started.  
2025/04/21 12:04:14 getHealth() -> Success: 21.541252637 secs elapsed since server started.  
2025/04/21 12:04:17 getHealth() -> Success: 24.692860701 secs elapsed since server started.  
2025/04/21 12:04:19 getHealth() -> Success: 26.54174077 secs elapsed since server started.  
2025/04/21 12:04:22 getHealth() -> Success: 29.692908308 secs elapsed since server started.  
2025/04/21 12:04:24 getHealth() -> Failed: 31.541647555 secs elapsed since server started.  
2025/04/21 12:04:27 getHealth() -> Failed: 34.693029777 secs elapsed since server started.  


- Check if pod was removed from service endpoint (`if pod IP is not listed that means Readines Probe failed`)          
    > kubectl get endpoints k8s-basics-assignment-bonus-service -o wide  

NAME                                  `ENDPOINTS`   AGE  
k8s-basics-assignment-bonus-service               3m20s  

**Note:** Endpoints has been removed that is because of Readiness Probe failure.  


**Below commands can be used to delete deployments and services**  

- To delete Deployment:
    > kubectl delete deployment <deployment-name>  

- To delete Service:
    > kubectl delete service <service-name>  
