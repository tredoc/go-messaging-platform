[![My Skills](https://skillicons.dev/icons?i=golang,docker,mongo,kafka)](https://skillicons.dev)

# Messaging platform

This repository contains the training material for the Messaging Platform.
The platform handles the messages sending to the users. It is overcomplicated for the sake of the training.

## Overview

The platform is composed of the following components:
1. **API Gateway**: The entry point of the platform. It is responsible for routing the requests to the appropriate services.
2. **Message Orchestrator**: The service responsible for orchestrating the messages sending. It is responsible for the following:
    - Enriching the message requests with the template.
    - Sending the message requests to the appropriate message sender.
    - Sending the message requests to service responsible for persisting the messages.
3. **Message Sender**: The service responsible for sending the messages to the users. It is responsible for the following:
    - Sending the message to the user.
    - Handling the message sending errors.
    - Sending the message sending status.
4. **Message Service**: The service responsible for persisting the messages. It is responsible for the following:
    - Persisting the message requests.
    - Handling the message persistence errors.
    - Persisting message status updates
5. **Message Template Service**: The service responsible for managing the message templates. It is responsible for the following:
    - Managing the message templates.
