# Honey

*Get that sweet sweet honey!*

A scalable HTTP Honey Pot for Fun and Mischief.

Broken into 2 primary parts (the collector and the consumer), you can deploy Honey across VMs of any shape and any size.

![HTTP Honey Pot Image](docs/img/http_honey_pot.png)

* Collectors capture requests, serialize them, and toss them on to a pubsub queue.
* Consumers grab tasks, parse them, format them and insert them into a MySQL table for analysis


## The setup is simple

1. Clone this repo
2. `go build .`
3. Config a GCP service account with pub/sub publisher permissions and download the json key.

```
GCP_CREDS=$(cat service-account-cred.json) ./honey-collector --ports "8080"
```
