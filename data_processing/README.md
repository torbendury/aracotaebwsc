# Data Processing

This is an application - written in Golang - that transforms data.

It does not have its own data backend, but has a callable HTTP endpoint to which XML data is sent.

This is transformed to JSON within the application and sent to another HTTP endpoint.

The data is product and stocking data that is processed for a business consisting of stores and an online shop and is used in both the stores and the online shop.

The application artifact is shipped within a Docker container image.

## Functional requirements

The application fulfills a contract for product data transformations.

Every product contains an `EAN`, a `ShortName` and `PriceCents`.

Incoming data which does not fulfill this contract should be discarded and not forwarded.

Data which fulfills this contract should be transformed from `XML` to `JSON`.

This serves multiple purposes:

- Check data integrity
- Catch business errors
- Enable high speed data processing

The application itself is stateless and does not save any kind of data, it only transforms and forwards it.

## Quality requirements

- No data may be lost
- Every data object must be processed in no more than 0.5 seconds (including receiving incoming data, transforming it and forwarding it to another endpoint)

## Overview

![A graphical context overview](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.github.com/torbendury/aracotaebwsc/main/data_processing/_docs/systemcontext.puml)
