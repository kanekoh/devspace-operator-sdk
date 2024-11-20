#!/bin/bash

export USER_NAME=`oc whoami`
export USER_TOKEN=`oc whoami -t`
export PROJECT_NAME=`oc config get-contexts --no-headers | awk '{print $5}'`

export IMAGE_REGISTRY=image-registry.openshift-image-registry.svc.cluster.local:5000/${PROJECT_NAME}
export IMAGE_TAG_BASE=${IMAGE_REGISTRY}/memcached-operator
export IMG=$IMAGE_TAG_BASE:${VERSION}
