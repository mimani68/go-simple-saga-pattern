#SAGA Pattern

ACID tnx in distributed system

The Saga Pattern is as microservices architectural pattern to implement a transaction that spans multiple services. A saga is a sequence of local transactions. Each service in a saga performs its own transaction and publishes an event. The other services listen to that event and perform the next local transaction

Credit: dzone
