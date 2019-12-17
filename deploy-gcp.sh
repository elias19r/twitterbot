#!/bin/bash

set -v

localImage=twitterbot
projectName=twitterbot
projectID=twitterbot-193613
imageURL=gcr.io/$projectID/$projectName
imageTag=latest
vmName=twitterbot-vm
vmZone=southamerica-east1-b

# Set project.
gcloud config set project $projectID

# Delete docker image.
gcloud container images delete $imageURL

# Push new docker image.
docker tag $localImage $imageURL 
gcloud docker -- push $imageURL

# Create VM instance with container.
# gcloud beta compute instances create-with-container $vmName --container-image $imageURL

# Update container in VM instance.
gcloud beta compute instances update-container $vmName --container-image $imageURL

# SSH to VM instance.
# gcloud compute --project $projectID ssh --zone $vmZone $vmName