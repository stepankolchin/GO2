# Практическое занятие №16 Публикация приложения в Kubernetes (минимальный манифест). Колчин Степан Сергеевич ЭФМО-02-25.

## Используемый сервис

Для деплоя выбран сервис `tasks` из практической работы №9 [prac9-redis-cache](../prac9-redis-cache/), который уже упакован в Docker-образ `techip-tasks:0.1`. Сервис имеет endpoint `/health` для проверки состояния.

## Окружение

- **Kubernetes-стенд:** Minikube v1.38.1 (локальный кластер)
- **Драйвер:** Docker
- **Инструменты:** kubectl v1.30.0

## Отчетные материалы

`Проверка доступа к кластеру`

<img width="974" height="214" alt="image" src="https://github.com/user-attachments/assets/111cab57-cb41-44fb-98fb-550b7aca3433" />

`Проверка Docker-образа`

<img width="974" height="138" alt="image" src="https://github.com/user-attachments/assets/8bb91c06-07d7-4d77-8c29-e8b80e0e1b56" />

`Создание манифестов`

## Манифесты

`configmap.yaml`

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: tasks-config
data:
  TASKS_PORT: "8082"
  REDIS_ADDR: "localhost:6379"
  LOG_LEVEL: "info"
```

`deployment.yaml`

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: tasks
spec:
  replicas: 1
  selector:
    matchLabels:
      app: tasks
  template:
    metadata:
      labels:
        app: tasks
    spec:
      containers:
      - name: tasks
        image: techip-tasks:0.1
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 8082
        envFrom:
        - configMapRef:
            name: tasks-config
        readinessProbe:
            httpGet:
                path: /health
                port: 8082
            initialDelaySeconds: 3
            periodSeconds: 5
            livenessProbe:
            httpGet:
                path: /health
                port: 8082
            initialDelaySeconds: 10
            periodSeconds: 10
```

`service.yaml`

```yaml
apiVersion: v1
kind: Service
metadata:
  name: tasks
spec:
  type: ClusterIP
  selector:
    app: tasks
  ports:
  - protocol: TCP
    port: 8082
    targetPort: 8082
```

`Применение манифестов`

<img width="944" height="267" alt="image" src="https://github.com/user-attachments/assets/9a726632-d6f8-4092-9c71-d2f9aa8336b5" />

`Проверка Pod и Deployment`

<img width="909" height="114" alt="image" src="https://github.com/user-attachments/assets/a7a0aa26-bcc2-4a17-8c6a-48bbc843f508" />


<img width="974" height="901" alt="image" src="https://github.com/user-attachments/assets/5af9dc0a-6a37-4f5e-9c8f-73795ae187ce" />


`Проверка Service`

<img width="911" height="573" alt="image" src="https://github.com/user-attachments/assets/9e8cd34f-6dae-40db-8c6c-f216a3c7f1de" />

`Проверка port-forward:`


<img width="905" height="86" alt="image" src="https://github.com/user-attachments/assets/3c5d854e-8173-43e0-9b4f-af7d097d02fa" />

<img width="425" height="195" alt="image" src="https://github.com/user-attachments/assets/3139dd69-d0e7-436d-832c-23ecc04a6301" />


