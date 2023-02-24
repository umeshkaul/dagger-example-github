### Pipelines as Code with Go, Dagger 
 

#### Build and push image to docker
Setup docker hub login  

```
% docker login  
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: ukaul
Password: 
Login Succeeded
```

Build and Publish  

```
l_proxy go run dagger/dagger.go

#1 resolve image config for docker.io/library/golang:latest
#1 DONE 2.4s

#2 
#2 DONE 0.0s

#3 host.directory /Users/ukaul/Projects/dagger-go-example/dagger-example
#3 transferring /Users/ukaul/Projects/dagger-go-example/dagger-example: 13.67kB 0.0s done
#3 DONE 0.0s

#4 from golang:latest
...


#8 pushing layers
#8 pushing layers 3.0s done
#8 pushing manifest for docker.io/ukaul/dagger-example:latest@sha256:40b2ae7727ce25191a705dda8beaf2168ba54df427d2d3c0ee6e88db2c4caf93
Succesfully created new container: docker.io/ukaul/dagger-example:latest@sha256:40b2ae7727ce25191a705dda8beaf2168ba54df427d2d3c0ee6e88db2c4caf93%   
```


#### Test the above container

launch container from one window   

`docker run --name test --rm -p 9090:9090 ukaul/dagger-example`

Test from other window   

```
% docker ps  
CONTAINER ID   IMAGE                           COMMAND                  CREATED         STATUS         PORTS                    NAMES
72be497f69b4   ukaul/dagger-example            "/app/dagger-example"    3 seconds ago   Up 2 seconds   0.0.0.0:9090->9090/tcp   test
f418f1ac69d3   ghcr.io/dagger/engine:v0.3.12   "dagger-entrypoint.s…"   9 minutes ago   Up 9 minutes                            dagger-engine-ee2f18ae9124c552

% curl localhost:9090  
Hello Dagger% 

```

#### Local build/test

```
% l_proxy docker build .    
[+] Building 2.7s (10/10) FINISHED                                                                                                    
 => [internal] load build definition from Dockerfile                                                                             0.0s
 => => transferring dockerfile: 37B                                                                                              0.0s
 => [internal] load .dockerignore                                                                                                0.0s
 => => transferring context: 2B                                                                                                  0.0s
 => [internal] load metadata for docker.io/library/alpine:latest                                                                 2.1s
 => [auth] library/alpine:pull token for registry-1.docker.io                                                                    0.0s
 => [1/4] FROM docker.io/library/alpine:latest@sha256:69665d02cb32192e52e07644d76bc6f25abeb5410edc1c7a81a10ba3f0efb90a           0.0s
 => => resolve docker.io/library/alpine:latest@sha256:69665d02cb32192e52e07644d76bc6f25abeb5410edc1c7a81a10ba3f0efb90a           0.0s
 => => sha256:69665d02cb32192e52e07644d76bc6f25abeb5410edc1c7a81a10ba3f0efb90a 1.64kB / 1.64kB                                   0.0s
 => => sha256:c41ab5c992deb4fe7e5da09f67a8804a46bd0592bfdf0b1847dde0e0889d2bff 528B / 528B                                       0.0s
 => => sha256:d74e625d91152966d38fe8a62c60daadb96d4b94c1a366de01fab5f334806239 1.49kB / 1.49kB                                   0.0s
 => [internal] load build context                                                                                                0.1s
 => => transferring context: 6.21MB                                                                                              0.1s
 => [2/4] RUN mkdir /app                                                                                                         0.2s
 => [3/4] COPY ./build/dagger-example /app/dagger-example                                                                        0.0s
 => [4/4] RUN chmod +x /app/dagger-example                                                                                       0.2s
 => exporting to image                                                                                                           0.0s
 => => exporting layers                                                                                                          0.0s
 => => writing image sha256:46f5f46a59eca1a65814da4efcc357d57a9a2099a31b7b73b02b1ea023a56357                                     0.0s
Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them


% docker ps                                                                                                                     14:38
CONTAINER ID   IMAGE                           COMMAND                  CREATED          STATUS          PORTS     NAMES
f418f1ac69d3   ghcr.io/dagger/engine:v0.3.12   "dagger-entrypoint.s…"   18 minutes ago   Up 18 minutes             dagger-engine-ee2f18ae9124c552
~/Projects/dagger-go-example/dagger-example                                                                                          

% docker image ls                                                                                                               14:38
REPOSITORY                                                                     TAG          IMAGE ID       CREATED          SIZE
<none>                                                                         <none>       46f5f46a59ec   11 seconds ago   19.9MB
ukaul/dagger-example                                                           latest       f164f7bf3e37   13 minutes ago   13.7MB
catthehacker/ubuntu                                                            act-latest   b752731e9ad5   8 days ago       1.08GB
ghcr.io/dagger/engine                                                          v0.3.12      2a8b21db3b83   2 weeks ago      183MB
catthehacker/ubuntu                                                            <none>       069c0c5b7bb5   3 weeks ago      1.08GB
ghcr.io/dagger/engine                                                          v0.3.10      3bf79859b97d   4 weeks ago      183MB
management-api                                                                 latest       772ee438406b   4 months ago     868MB
artprod.dev.bloomberg.com/babka/memcached                                      1.5.22.1     bd39b064e219   2 years ago      224MB
artifactory.inf.bloomberg.com/ds/ext/registry-1.docker.io/mysql/mysql-server   5.7.31       9c31a29b3f30   2 years ago      322MB
artifactory.inf.bloomberg.com/ds/ext/registry-1.docker.io/library/redis        6.0.5        235592615444   2 years ago      104MB
artifactory.inf.bloomberg.com/dspuser/s-babka-test/babka-argo-deploy-job       latest       96d7ea23f84e   43 years ago     887MB
                                                                               
# launch local container
% docker run --rm --name test -p9090:9090 46f5f46a59ec   
...

#from another window

% curl localhost:9090                                                                                                                            14:40
Hello Dagger%

```

