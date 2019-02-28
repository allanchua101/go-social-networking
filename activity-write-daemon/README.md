# Activity Write Daemon

This daemon is in-charge of continously monitoring the activity write queue. A slave replica of it is also in-charge of monitoring the fail-over / slave queue.

Each message received from the write queue will undergo the following processing:

- Activity event is then appended to the event store.
- A projection update is then offloaded to worker queues dedicated for presenting different views of application state.
