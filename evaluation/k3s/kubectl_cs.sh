kubectl get nodes

kubectl get pods -n=spegel

kubectl logs spegel-njqqb -n=spegel

kubectl describe pod spegel-njqqb -n=spegel

kubectl delete all --all --namespace=spegel



#### SPEGEL

sudo crictl images
sudo crictl pull node:slim
