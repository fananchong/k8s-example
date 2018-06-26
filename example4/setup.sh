#!/bin/bash


helm del --purge my-redis
kubectl delete pvc --all -n my-redis
kubectl delete -f pv.yaml
kubectl create -f pv.yaml
mkdir -p /tmp/data/local-redis-pv-1
mkdir -p /tmp/data/local-redis-pv-2
mkdir -p /tmp/data/local-redis-pv-3
mkdir -p /tmp/data/local-redis-pv-4
chmod a+rw /tmp/data/*
helm install --name=my-redis --namespace=my-redis -f values-production.yaml stable/redis
