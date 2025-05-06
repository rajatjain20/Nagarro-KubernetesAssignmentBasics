# Kubernetes Basics Assignment

**`Note:`** I have installed Kubernetes and minikube on my windows machine. (Refer Installation_kubernetes.docx`) 

Start Minikube using the following command:
<pre>minikube start --driver=docker </pre>

## Overview
 I have created one helloworld app in golang. It is just one static page.  
   
📂 Directory Descriptions:

| Path                       | Description                                                   |
|---------------------------|---------------------------------------------------------------|
| `./Kubernetes_Assignment_Basics/app/`                  | It contains helloworld application with its Dockerfile.       |
| `./Kubernetes_Assignment_Basics/Deployment.yaml`       | Deployment YAML file for Kubernetes.                          |
| `./Kubernetes_Assignment_Basics/Service.yaml`          | Service YAML to expose the app externally.                    |
| `./Kubernetes_Assignment_Basics/Deployment-Bonus.yaml` | Bonus deployment with liveness & readiness probes.            |
| `./Kubernetes_Assignment_Basics/Helloworld_App_SS.jpg` | Screenshot of the helloworld web app.                         |
| `./Installation_kubernetes.docx` | Steps to install Kubernetes on Windows.              |

## Create Docker image inside Kubernetes cluster

This part can be skipped as it is not necessary to build an image inside Kubernetes cluster. I am using pushed image from docker hub in deployments.

- Execute Windows PowerShell.

- Command to set Windows Poweshell to use Minikube’s Docker daemon:
    <pre> > & minikube -p minikube docker-env | Invoke-Expression </pre>

- Now check available images inside Kubernetes:  

    <pre> > docker images </pre>

- Go to directory `/Kubernetes_Assignment_Basics/app/` [replace path with complete path on your local]  

    <pre>> cd "/Kubernetes_Assignment_Basics/app/"</pre>  

- Command to build image inside kubernetes  

    <pre> > docker build -t k8sassignmentbasics:v0 . </pre> 
    <pre> > docker images
    REPOSITORY                      TAG        IMAGE ID       CREATED          SIZE
    k8sassignmentbasics             v0         6b2d4a666de9   10 seconds ago   919MB</pre>

    Image has been created inside Kubernetes cluster as shown above.

## Deployment of application in Kubernetes Cluster

I have created a yaml file named `Deployment.yaml`, placed inside `./Kubernetes_Assignment_Basics/`
  
- Execute Windows Powershell.

- Go to directory `/Kubernetes_Assignment_Basics/` [replace path with complete path on your local]
    <pre> > cd "/Kubernetes_Assignment_Basics/" </pre>

- Create Deployment:
    <pre> > kubectl create -f ./Deployment.yaml --save-config 
    deployment.apps/k8s-basics-assignment created</pre>      

- To check deployment:
    <pre> > kubectl get all

    NAME                                        READY   STATUS    RESTARTS   AGE
    pod/k8s-basics-assignment-898db5488-twx22   1/1     Running   0          11s

    NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
    service/kubernetes   ClusterIP   10.96.0.1    <<none>none>        443/TCP   65s

    NAME                                    READY   UP-TO-DATE   AVAILABLE   AGE
    deployment.apps/k8s-basics-assignment   1/1     1            1           11s

    NAME                                              DESIRED   CURRENT   READY   AGE
    replicaset.apps/k8s-basics-assignment-898db5488   1         1         1       11s</pre> 


## Expose the application using a Kubernetes Service to access it externally. 

I have created a yaml file named `Service.yaml`, placed inside directory `./Kubernetes_Assignment_Basics/`.  

- Command to create service to expose the application:  
    <pre> > kubectl create -f ./Service.yaml --save-config  
    service/k8s-basics-assignment-service created </pre>

- To check created service:
    <pre> > kubectl get services  
    NAME                            TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
    k8s-basics-assignment-service   NodePort    10.109.216.141   <<none>none>        3010:30036/TCP   10s
    kubernetes                      ClusterIP   10.96.0.1        <<none>none>        443/TCP          6m30s</pre>  

- To run exposed application locally:
    <pre> > minikube service k8s-basics-assignment-service 
    |-----------|-------------------------------|-------------|---------------------------|
    | NAMESPACE |             NAME              | TARGET PORT |            URL            |
    |-----------|-------------------------------|-------------|---------------------------|
    | default   | k8s-basics-assignment-service |        3010 | http://192.168.49.2:30036 |
    |-----------|-------------------------------|-------------|---------------------------|
    * Starting tunnel for service k8s-basics-assignment-service.
    |-----------|-------------------------------|-------------|------------------------|
    | NAMESPACE |             NAME              | TARGET PORT |          URL           |
    |-----------|-------------------------------|-------------|------------------------|
    | default   | k8s-basics-assignment-service |             | http://127.0.0.1:55435 |
    |-----------|-------------------------------|-------------|------------------------|
    * Opening service default/k8s-basics-assignment-service in default browser...
    ! Because you are using a Docker driver on windows, the terminal needs to be open to run it.</pre>  

**`Note:`** This will execute the application on default browser.  

## Scale the application by increasing the number of replicas

- Scale the application:  

    * Method 1:  
        <pre>> kubectl scale deployment k8s-basics-assignment --replicas=3  
        deployment.apps/k8s-basics-assignment scaled</pre>

    * Method 2:  
        We can change the count from 1 to any number (max 110) in Deployment.yaml file and can execute below command:  
        <pre>> kubectl apply -f ./Deployment.yaml  
        deployment.apps/k8s-basics-assignment configured </pre>

- To verify scaled application:  
    <pre> > kubectl get deployment k8s-basics-assignment  
    NAME                    READY   UP-TO-DATE   AVAILABLE   AGE
    k8s-basics-assignment   5/5     5            5           13m</pre>

- To check pods:
    <pre> > kubectl get pods
    NAME                                    READY   STATUS    RESTARTS   AGE
    k8s-basics-assignment-898db5488-72wx8   1/1     Running   0          103s
    k8s-basics-assignment-898db5488-8ksgb   1/1     Running   0          3m8s
    k8s-basics-assignment-898db5488-p2pjd   1/1     Running   0          103s
    k8s-basics-assignment-898db5488-rlwpl   1/1     Running   0          3m8s
    k8s-basics-assignment-898db5488-twx22   1/1     Running   0          15m</pre>  

## `Bonus Question`: Enhance the application deployment by adding a health check endpoint and configuring liveness and readiness probes in the Kubernetes deployment.

I have added one health check endpoint (**`"/healthz"`**) in my source code (image tag is updated to v1) and added liveness and readiness probes in `Deployment-Bonus.yaml`.  
This endpoint returns error code (500) after 30secs of the server startup and keeps returning it till 50secs. This error code (500) is considered as its failure state by kubernetes liveness probe and readiness probe.  

- Let's deploy this deployment:
    <pre> > kubectl create -f ./Deployment-Bonus.yaml --save-config  
    deployment.apps/k8s-basics-assignment-bonus created  </pre>

- Check deployment:
    <pre> > kubectl get deployment k8s-basics-assignment-bonus
    NAME                          READY   UP-TO-DATE   AVAILABLE   AGE
    k8s-basics-assignment-bonus   1/1     1            1           29s</pre>  

- Expose this to a service:
    <pre> > kubectl expose deployment k8s-basics-assignment-bonus --name=k8s-basics-assignment-bonus-service --port=3020 --target-port=3000 --type=NodePort  
    service/k8s-basics-assignment-bonus-service exposed </pre>  

- Check the service created:
    <pre> > kubectl get svc k8s-basics-assignment-bonus-service  
    NAME                                  TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
    k8s-basics-assignment-bonus-service   NodePort   10.97.149.100   <<none>none>        3020:31508/TCP   62s</pre>

- Check Endpoints associated with service created above (`k8s-basics-assignment-bonus-service`):
    <pre> > kubectl get endpoints k8s-basics-assignment-bonus-service -o wide
    NAME                                  ENDPOINTS           AGE
    k8s-basics-assignment-bonus-service   10.244.0.151:3000   5m</pre>  
  
- Wait for ~50secs or more and check the resources:  
    <pre> > kubectl get all 
    NAME                                               READY   STATUS    RESTARTS        AGE
    pod/k8s-basics-assignment-898db5488-72wx8          1/1     Running   0               16m
    pod/k8s-basics-assignment-898db5488-8ksgb          1/1     Running   0               17m
    pod/k8s-basics-assignment-898db5488-p2pjd          1/1     Running   0               16m
    pod/k8s-basics-assignment-898db5488-rlwpl          1/1     Running   0               17m
    pod/k8s-basics-assignment-898db5488-twx22          1/1     Running   0               29m
    pod/k8s-basics-assignment-bonus-7f96c49df4-nppmz   1/1     Running   7 (3m21s ago)   10m 👈

    NAME                                          TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
    service/k8s-basics-assignment-bonus-service   NodePort    10.97.149.100    <none>        3020:31508/TCP   8m31s 👈
    service/k8s-basics-assignment-service         NodePort    10.109.216.141   <none>        3010:30036/TCP   24m
    service/kubernetes                            ClusterIP   10.96.0.1        <none>        443/TCP          30m

    NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
    deployment.apps/k8s-basics-assignment         5/5     5            5           29m
    deployment.apps/k8s-basics-assignment-bonus   1/1     1            1           10m 👈

    NAME                                                     DESIRED   CURRENT   READY   AGE
    replicaset.apps/k8s-basics-assignment-898db5488          5         5         5       29m
    replicaset.apps/k8s-basics-assignment-bonus-7f96c49df4   1         1         1       10m 👈</pre> 

    We can see `pod/k8s-basics-assignment-bonus-7f96c49df4-nppmz` was restarted 7 times. That is because of Liveness failure. We can verify this in pod description's Events section in next step.

- Let's check the events in pod (pick the pod name from above command):  
    <pre> > kubectl describe pod k8s-basics-assignment-bonus-7f96c49df4-nppmz  
    Name:             k8s-basics-assignment-bonus-7f96c49df4-nppmz
    Namespace:        default
    ------
    ------

    Events:
    Type     Reason     Age                  From               Message
    ----     ------     ----                 ----               -------
    Normal   Scheduled  16m                  default-scheduler  Successfully assigned default/k8s-basics-assignment-bonus-7f96c49df4-nppmz to minikube
    Normal   Pulling    16m                  kubelet            Pulling image "rajatjain20/k8sassignmentbasics:v1"
    Normal   Pulled     15m                  kubelet            Successfully pulled image "rajatjain20/k8sassignmentbasics:v1" in 16.653s (16.653s including waiting). Image size: 918814242 bytes.
    Warning  Unhealthy  13m (x9 over 15m)    kubelet            Liveness probe failed: HTTP probe failed with statuscode: 500
    Normal   Created    12m (x6 over 15m)    kubelet            Created container: kubernetes-assignment-bonus
    Normal   Started    12m (x6 over 15m)    kubelet            Started container kubernetes-assignment-bonus
    Normal   Killing    5m46s (x8 over 15m)  kubelet            Container kubernetes-assignment-bonus failed liveness probe, will be restarted
    Warning  BackOff    5m9s (x27 over 11m)  kubelet            Back-off restarting failed container kubernetes-assignment-bonus in pod k8s-basics-assignment-bonus-7f96c49df4-nppmz_default(e6ea025d-da65-41ec-b96e-5ceee541fe5c)
    Normal   Pulled     39s (x8 over 15m)    kubelet            Container image "rajatjain20/k8sassignmentbasics:v1" already present on machine
    Warning  Unhealthy  7s (x38 over 15m)    kubelet            Readiness probe failed: HTTP probe failed with statuscode: 500 👈</pre>

    If we check the Messages against `Unhealthy` warnings, the first time pod was Unhealthy due to Liveness probe failed and the next Unhealthy was due to Readiness probe failed.

    Also, if we check the Message for the reason of `Killing` the pod, it says `Container kubernetes-assignment-bonus failed liveness probe, will be restarted`. 

- Check logs on the pod:
    <pre> > kubectl logs pod/k8s-basics-assignment-bonus-7f96c49df4-nppmz
    Hello World application
    Listning at port :3000
    2025/04/21 21:25:39 getHealth() -> Success: 4.739826326 secs elapsed since server started.
    2025/04/21 21:25:41 getHealth() -> Success: 6.4278104 secs elapsed since server started.
    2025/04/21 21:25:44 getHealth() -> Success: 9.739792995 secs elapsed since server started.
    2025/04/21 21:25:46 getHealth() -> Success: 11.427414692 secs elapsed since server started.
    2025/04/21 21:25:49 getHealth() -> Success: 14.740872 secs elapsed since server started.
    2025/04/21 21:25:51 getHealth() -> Success: 16.427930092 secs elapsed since server started.
    2025/04/21 21:25:54 getHealth() -> Success: 19.740019475 secs elapsed since server started.
    2025/04/21 21:25:56 getHealth() -> Success: 21.428176116 secs elapsed since server started.
    2025/04/21 21:25:59 getHealth() -> Success: 24.740330353 secs elapsed since server started.
    2025/04/21 21:26:01 getHealth() -> Success: 26.428349498 secs elapsed since server started.
    2025/04/21 21:26:04 getHealth() -> Success: 29.740868935 secs elapsed since server started.
    2025/04/21 21:26:06 getHealth() -> Failed: 31.427947588 secs elapsed since server started.
    2025/04/21 21:26:09 getHealth() -> Failed: 34.739827316 secs elapsed since server started.
    2025/04/21 21:26:11 getHealth() -> Failed: 36.428470035 secs elapsed since server started.
    2025/04/21 21:26:14 getHealth() -> Failed: 39.739442253 secs elapsed since server started.
    2025/04/21 21:26:16 getHealth() -> Failed: 41.428446596 secs elapsed since server started.
    2025/04/21 21:26:16 getHealth() -> Failed: 41.431024515 secs elapsed since server started.
    2025/04/21 21:26:19 getHealth() -> Failed: 44.740278361 secs elapsed since server started.
    2025/04/21 21:26:19 getHealth() -> Failed: 44.742119879 secs elapsed since server started.</pre>


- Check if pod was removed from service endpoint (`if pod IP is not listed that means Readines Probe failed`)          
    <pre> > kubectl get endpoints k8s-basics-assignment-bonus-service -o wide  
    NAME                                  ENDPOINTS   AGE
    k8s-basics-assignment-bonus-service               19m</pre>

**`Note:`** Endpoint has been removed that is because of Readiness Probe failure.  

### Below commands can be used to delete deployments and services  

- To delete Deployment:
    > kubectl delete deployment `<deployment-name>`  

- To delete Service:
    > kubectl delete service `<service-name>`  
