# golang-kubernetes-api
A simple hello world API written in Go, containerized and deployed using Kubernetes.

To setup on Windows:

1. Ensure Docker Desktop is installed
2. Run Docker Desktop
3. Navigate to settings and enable Kubernetes
4. Open a terminal in this working directory
5. run:
`docker build --tag=sunrisebanana/test` (note: this will only work if you are signed in to docker as sunrisebanana. Otherwise, change the command to `docker build --tag=<YOUR_DOCKER_USERNAME_HERE>/test` and edit index.yml line 18 to correspond to the `<YOUR_DOCKER_USERNAME_HERE>/test` tag.)
6. now run `kubectl apply -f index.yml` and navigate to [localhost:30001](localhost:30001). You're Done! 

Once finished, make sure to run `kubectl delete -f index.yml` to destruct the created kubernetes objects.