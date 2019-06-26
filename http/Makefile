apiVersion: apps/v1
kind: Deployment
metadata:
  name: jessewudemo
  labels:
    app: k8s-jessewudemo
spec:
  replicas: 1
  selector:
    matchLabels:
      app: k8s-jessewudemo
  template:
    metadata:
      labels:
        app: k8s-jessewudemo
    spec:
      containers:
        - name: jessewudemo
          image: jessewu/demo:latest
          ports:
            - containerPort: 80
              protocol: TCP

---
apiVersion: v1
kind: Service
metadata:
  name: jessewudemo
  labels:
    app: k8s-jessewudemo
spec:
  type: NodePort
  selector:
    app: k8s-jessewudemo
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80