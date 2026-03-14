# 🔐 Secure Access & Audit Microservice Platform

A **DevOps + Cybersecurity focused microservice platform** built using **Golang, Docker, Terraform and Jenkins**.
This project demonstrates how to secure APIs using **JWT authentication** and automate build processes with a **CI pipeline**.

---

## 🚀 Project Overview

This platform provides **secure access to APIs** and keeps track of system activity for auditing purposes.

Main goals of this project:

* Implement **secure authentication using JWT**
* Build **microservices using Golang**
* Automate builds using **Jenkins CI**
* Containerize services with **Docker**
* Manage infrastructure using **Terraform**

---

## 🏗️ Architecture

```
Developer
   │
   ▼
GitHub Repository
   │
   ▼
Jenkins CI Pipeline
   │
   ▼
Docker Container Build
   │
   ▼
Golang Microservices
   │
   ▼
JWT Authentication & Secure APIs
```

---

## 🧩 Project Features

### 🔑 JWT Authentication

* Secure login endpoint
* Token-based authentication
* Protected APIs
* Token validation

### ⚙️ DevOps Automation

* Jenkins CI pipeline
* Automated Go build process
* Docker container integration

### 🛡️ Security

* Token based access control
* Secure API endpoints
* Authentication layer for microservices

---

## 📁 Project Structure

```
devops-security-platform
│
├── terraform/
│   └── main.tf
│
├── token_service.go
├── Jenkinsfile
└── README.md
```

---

## 🛠️ Technologies Used

* Golang
* Docker
* Terraform
* Jenkins
* JWT Authentication
* GitHub

---

## ▶️ Running the Project

### 1️⃣ Clone the Repository

```
git clone https://github.com/NamanPachauli/devops-security-platform.git
```

### 2️⃣ Run the Authentication Service

```
go run token_service.go
```

Server will start on:

```
http://localhost:7072
```

---

### 3️⃣ Generate JWT Token

```
http://localhost:7072/login?user=test
```

Response:

```
Token: <JWT_TOKEN>
```

---

### 4️⃣ Access Secure Endpoint

Send request with header:

```
Authorization: <JWT_TOKEN>
```

Endpoint:

```
GET /secure
```

---

## 📦 CI Pipeline

The project includes a **Jenkins pipeline** that:

* Pulls code from GitHub
* Builds the Golang service
* Runs inside Docker container

---

## 🎯 Learning Objectives

This project was built to practice:

* DevOps CI/CD workflows
* Secure API design
* Containerized microservices
* Infrastructure automation

---

## 👨‍💻 Author

**Naman Pachauli**

DevOps & Cloud Enthusiast

---

## ⭐ Future Improvements

* Kubernetes deployment
* Audit logging microservice
* Database integration
* Role-based access control
