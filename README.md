# OpenUEM NATS Service

This repository contains the definition of the OpenUEM NATS service.

NATS is free to use under the Apache-2.0 Open Source License. OpenUEM uses a NATS server that uses TLS and mutual TLS authentication. Also it uses Jetstream, which enables messages to be stored and replayed at a later time (for example mail messages that haven't been delivered to the server or requests to generate user certificates)

## NATS Configuration summary

Jetstream. Here are the current use of Jetstreams

- Notification Worker. Creates the NOTIFICATION_STREAM to consume notification requests
- Agent Worker. Creates the AGENT_CERTIFICATE_STREAM to consume requests to send agent server certificates to the agents
