# **SiJaki-Backend**
Sijaki-Backend is the backend service for the **SiJaki** project, a centralized waste management system for smart cities. It handles smart trash bin data, user management, and automated messaging to workers when bins are full.

---

## **📃Table of Contents**
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Configuration](#configuration)
- [API Documentation](#api-documentation)
- [Closing Remarks](#closing-remarks)

---

## **🚀Features**
- Manage smart trash bin data (locations, status, volume).
- Automatic notifications to workers when bins are full.
- MQTT-based real-time data processing from trash bins.
- WebSocket integration for real-time updates on device data.
- Admin functionalities for worker management.
- RESTful API endpoints for easy integration with the frontend.

---

## **⚒️Technologies Used**
- **Programming Language:** [Go](https://go.dev/)
- **Database:** PostgreSQL
- **Real-Time Communication:** MQTT with [Eclipse Mosquitto](https://mosquitto.org/)
- **Web Framework:** [Fiber](https://gofiber.io/)

---

## **✈️Installation**
### Prerequisites
- Docker
- Docker Compose
### Usage
1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/sijaki-backend.git
    ```
2. Run docker compose using:
    ```bash
    docker-compose up -d
    ```
    This will build and start all the necessary services in detached mode.
3. Access the Websocket at:
    this endpoint to get infromation about data in device sijaki
    ```
    ws://localhost:8080/ws
    ```
    You can use this WebSocket to listen for real-time data updates sent from the connected devices.
4. Access the API at:
    ```
    http://localhost:8080
    ```
    You can use standard HTTP requests for interacting with other API endpoints.
---

## **👨‍💻API Documentation**
### Users
- ```POST /api/users/register```    
- ```POST /api/users/login```
- ```POST /api/users/logout```
### Workers
- ```POST /api/workers/add```
- ```GET /api/workers/register```
- ```PUT /api/workers/edit```
- ```DEL /api/workers/delete```

---

## **🫡Closing Remarks**
Thank you for taking the time to review this project. While this project may not be fully developed to its maximum potential, I have focused on delivering a working solution that addresses urban waste management.

Given the breadth of other projects and commitments at the university, I wasn't able to dedicate as much time as I would have liked, but I believe the concept holds significant potential for further development.

I welcome any feedback or suggestions for improvement and would be grateful for your insights.

Thank you again for your time and consideration.

---
