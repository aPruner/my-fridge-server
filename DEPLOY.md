# Deploying the app (to GCP K8s Engine Cluster)

## Prerequisites:
- Kubernetes cli (`kubectl`): [Installation guide](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- Docker cli (`docker` and `docker-compose`): [Installation guide for docker](https://docs.docker.com/get-docker/), [installation guide for docker-compose](https://docs.docker.com/compose/install/)
- Kompose cli (`kompose`) - only required if there are changes made to the `docker-compose.yml` and so `deploy/*.yaml` files need to be regenerated: [Installation guide](https://kompose.io/setup/)

## Follow these steps: (TODO: AUTOMATE THIS with a shell script or CI/CD deployment pipeline)
1. Checkout master branch that is ready to deploy, pull in new changes
2. Ensure the go app compiles and runs locally with `bash start.sh` and there are no breaking changes
3. Build, tag, and push docker images to google cloud container repository
  For now (until this is automated), this step will look like the following:
    1. Run `docker-compose up --build` to build the new images (this will build and start running locally all of adminer, postgres, adampruner/my-fridge-server, adampruner/my-fridge-server-migrations)
    2. Run `docker tag <base-image-name> <gcp-container-registry-hostname>/<gcp-project-id>/<new-image-name>:<optional-tag>`, for both the adampruner/my-fridge-server and adampruner/my-fridge-server-migrations base image names
    3. Run `docker push <new-image-name> <gcp-container-registry-hostname>/<gcp-project-id>/<new-image-name>:<optional-tag>`
    (Note: ideally I will bump versions with each deploy/release, but I'm not sure at which step this should happen. Maybe before step 1, and then when I tag the image)
4. If there are changes to the `docker-compose.yml`, mirror them in the appropriate `deploy/*.yaml` file. Running `kompose convert` to regenerate the `.yaml` files might be necessary in some rare cases
5. Ensure `kubectl` current-context (can check with `kubectl get current-context`) is set to the GCP cluster [Guide on this here](https://cloud.google.com/kubernetes-engine/docs/how-to/cluster-access-for-kubectl) (TODO: AUTOMATE THIS)
6. Bring up deployments, services, and pods in the correct order (TODO: AUTOMATE THIS)
  For now, this will look like:
    1. Run `kubectl apply -f adminer-service.yaml,adminer-deployment.yaml,db-service.yaml,db-deployment.yaml,dbvol-persistentvolumeclaim.yaml` to deploy the db and adminer services
    2. Wait a bit, confirm in GCP that the db and adminer services deployed successfully and are OK, maybe grab a drink
    3. If things went wrong, troubleshoot. If not, Run `kubectl apply -f app-service.yaml,app-deployment.yaml,migrations-pod.yaml` to deploy the app service and migrations pod
    4. If needed, troubleshoot. If not, expose the app and adminer services to the world by running `kubectl expose deployment app --type=LoadBalancer --name=my-fridge-server-service && kubectl expose deployment adminer --type=LoadBalancer --name=my-fridge-server-adminer`

Voila! After all of these steps, the app should be deployed to production! There is still lots of work to be done in automating this process, and enforcing security, no db bumps in the road, etc

## Future issues to consider/work on (create github issues):
1. HTTPS on K8s cluster, right now there is no SSL layer
2. Deployment automation or CI/CD pipeline
3. Potential postgres issues (right now using PGDATA=/tmp env var, might reset the database on deploy)
4. Mapping LoadBalancer external IPs to actual domain names, so that mobile and web apps can reliably hit the server API
    