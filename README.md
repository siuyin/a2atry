# A2A (Agent to Agent) Protocol experiments.

The code in this repository is based on:
- https://github.com/a2aproject/a2a-go
- https://github.com/a2aproject/a2a-samples/tree/main/samples/go

## Observations
The protocol spec as of 16 Aug 2025 is at v0.3 and is still fluid.

For example: https://github.com/a2aproject/a2a-samples/blob/dbd790fc65be9a7c7de8fd64f22cc45872bf4d44/samples/go/models/requests.go#L4
specifies type `TaskSendParams`.
However https://github.com/a2aproject/a2a-go/blob/143403d47d851604e0b128fbdc341f06fc7c8852/a2a/core.go#L334
specifies type `MessageSendParams`.
And in the official spec: https://github.com/a2aproject/A2A/blob/00cf76e7bbc752842ef254f3d4136ed1b5751f6e/types/src/types.ts#L650
we also have `MessageSendParams`.

This means `samples` is out of date with respect to the official spec.
