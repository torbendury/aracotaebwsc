# Architectural Requirements and Code Optimizations to Achieve Environmental Benefits With Serverless Computing

## RESTful API

The first scenario is a classic RESTful application that offers a CRUD interface. The application's data store is a self-managed PostgreSQL database cluster that runs on 2 VMs (1 main instance, 1 failover instance). The database stores order data for an e-commerce store.

The application was written in Golang and runs within a Debian Docker container on an on-premises Kubernetes cluster with three pods in the production environment (the development and integration environments are ignored for this scenario, but they are structured the same). The app has a predictable usage pattern; users use it on working days (Monday-Saturday) mainly between 7:00 a.m. and 10:00 p.m. and send orders.

## Data Processing

The second scenario is an application - also written in Golang - that transforms data. It does not have its own data backend, but has a callable HTTP endpoint to which XML data is sent. This is transformed to JSON within the application and sent to another HTTP endpoint. The data is product and stocking data that is processed for a business consisting of stores and an online shop and is used in both the stores and the online shop.
