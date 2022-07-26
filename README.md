# Honey

*Get that sweet sweet honey!*

A scalable HTTP Honey Pot for Fun and Mischief.

Broken into 2 primary parts (the collector and the consumer), Honey is able to de deployed across a variety of infrastructure types. Deploying sensors in a large number is advantageous to collect the most interesting data. Honey allows you to deploy a single Go binary to machines of any shape and any size to allow for the maximum flexibility possible.

## The setup is simple

![HTTP Honey Pot Image](docs/img/http_honey_pot.png)

* Collectors capture requests, serialize them, and toss them on to a pubsub queue.
* Consumers grab tasks, parse them, format them and insert them into a MySQL table for analysis
