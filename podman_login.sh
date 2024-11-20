#!/bin/bash

if [ -z $USER_NAME ]; then
    echo "Please execute the following command first."
    echo "source env.sh"
    exit 1
fi

podman login -u $USER_NAME -p $USER_TOKEN $IMAGE_REGISTRY
