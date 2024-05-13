# GogRPC
<<<<<<< HEAD
=======
Projet Go et gRPC
>>>>>>> 8968994f119f5324e7717ef1da3807a649fd0f23

<div id="top"></div>
<!-- PROJECT LOGO -->

<p align="left">
<<<<<<< HEAD
  An integrated project for managing IoT devices, APIs, and statistical dashboards. This project focuses on efficiently handling QR code-based ticketing systems and other IoT functionalities. Explore advanced data analytics and visualizations through the provided dashboard for enhanced insights and management capabilities.
=======
  The project is to develop two programs in Go language, initiated to meet the needs of our company. The latter manages a fleet of devices deployed at various clients, these devices generating daily files in JSON format containing a summary of their operations. The objective of the project is to design a system composed of a client and a server, using gRPC as a communication protocol, in order to collect this data, transmit it to the server, store it in the MongoDB database, and perform necessary updates.
>>>>>>> 8968994f119f5324e7717ef1da3807a649fd0f23
  <br />
  <br />
</p>

---

## Table of Contents

<details open>
  <summary>Click here to expand the Table of Contents</summary>
  <ol>
    <li>
      <a href="#overview">Overview</a>
      <ul>
        <li><a href="#Client Module">Client Module</a></li>
        <li><a href="#Server Module">Server Module</a></li>
      </ul>
    </li>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li><a href="#getting-started">Getting Started</a></li>
    <ul>
      <li><a href="#installation">Installation</a></li>
      <li><a href="#docker-usage">Docker Usage</a></li>
      <li><a href="#prerequisites">Prerequisites</a></li>
    </ul>
    <li><a href="#contributors">Contributors</a></li>
  </ol>
</details>

---

## Overview

In this project, we have 2 main parts :

### Client Module
The client is responsible for reading JSON files from remote devices.
It uses gRPC to send this data to the server.
Its main role is to collect information from JSON files and initiate communication with the server for sending this data.

</br>

### Server Module

The server receives data from the client via gRPC.
It stores this data in a MongoDB database.
He is responsible for managing the storage and updating operations of data records.

</br>

### System Architecture

![Project Diagramme][project_diagram]

---

## About The Project

The project is to develop two programs in Go language, initiated to meet the needs of our company. The latter manages a fleet of devices deployed at various clients, these devices generating daily files in JSON format containing a summary of their operations. The objective of the project is to design a system composed of a client and a server, using gRPC as a communication protocol, in order to collect this data, transmit it to the server, store it in the MongoDB database, and perform necessary updates.

### Key Features:

- **Ticket Purchase and Wallet Integration:** The project enables clients to purchase tickets directly from the website and seamlessly integrates a wallet feature to ensure they never lose their tickets. This allows for easy access and management of tickets.

- **QR Code Simulation:** Clients can simulate their QR codes to verify their validity before use. This feature ensures that tickets are valid and can be easily authenticated.

- **Integration with IoT Devices:** The website communicates with IoT devices to send QR codes, enabling seamless interaction between the user interface and physical devices.

- **API Integration:** The project utilizes APIs to manage and retrieve data related to tickets and other system functionalities. This ensures smooth communication and data exchange between different components of the system.

![API Diagram][api_diagram]

The project provides advanced data analytics and visualizations through the dashboard, offering enhanced insights and management capabilities for users.

![Project Ticket List][liste_ticket]

This dashboard provides a comprehensive view of key statistics for the transportation service, enabling data-driven decision making and strategic planning."

![Project Stats][stats]

<p align="right">(<a href="#GogRPC">back to top</a>)</p>

## `Getting Started`

To set up the project locally, follow these steps:

### `Installation`

1. **Clone the Repository**: Clone the VideoLibManager repository to your local machine:

   ```bash
   git clone https://github.com/Suhail1929/GogRPC.git
   ```

2. **Navigate to Project Directory**: Move into the project directory:

   ```bash
   cd GogRPC
   ```

   then escape and save the file by typing `:wq` and pressing `Enter`.

## `Docker Usage`

For Docker usage, follow these steps:

1. **Build Docker Containers**: To build the Docker containers, execute:

   ```bash
   sudo docker-compose up --build -d
   ```

2. **Start Containers**: Start the containers with:

   ```bash
   sudo docker-compose up -d
   ```

3. **Stop Containers**: To stop the running containers, use:
   ```bash
   sudo docker-compose down
   ```

## `Prerequisites`

List any prerequisites or dependencies users need to have before using the project.

- **Go**: Make sure GO is installed on your system. If not, you can download and install it from the [official GO website](https://go.dev/doc/install).

- **Install Docker**: Ensure Docker is installed on your system. If not, follow the official [Docker installation guide](https://docs.docker.com/get-docker/).

- **gRPC**: Make sure GO is installed on your system. If not, you can download and install it from the [official gRPC website](https://grpc.io/docs/languages/go/quickstart/).

<p align="right">(<a href="#GogRPC">back to top</a>)</p>

---

### `Contributors`

Thank you to the following contributors for their valuable contributions to this project:

<div style="display: flex; flex-wrap: wrap;">

  <div style="display: flex; align-items: center; margin-right: 20px;">
    <img src="https://github.com/Suhail1929.png" alt="Souhail's Profile" style="width: 50px; height: 50px; border-radius: 50%; margin-right: 10px;">
    <a href="https://github.com/Suhail1929" style="text-decoration: none;">Souhail</a>
  </div>

  <div style="display: flex; align-items: center; margin-right: 20px;">
    <img src="https://github.com/MedBakri.png" alt="Hugo's Profile" style="width: 50px; height: 50px; border-radius: 50%; margin-right: 10px;">
    <a href="#" style="text-decoration: none;">Med Bakri</a>
  </div>
</div>
