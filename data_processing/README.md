# Data Processing

This is an application - written in Golang - that transforms data.

It does not have its own data backend, but has a callable HTTP endpoint to which XML data is sent.

This is transformed to JSON within the application and sent to another HTTP endpoint.

The data is product and stocking data that is processed for a business consisting of stores and an online shop and is used in both the stores and the online shop.
