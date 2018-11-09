# DNS to dns to tls proxy

- What are the security concerns for this kind of service?

You can't verify in this case that's dns answers received from authorithive DNS server. 
All requests to this service comes without any encryption so possible MiTM attacks or all request can be tracked.

- Considering a microservice architecture; how would you see this the dns to dns-over-tls proxy used?

It can be used for some legacy software which doesn't support dns-over-tls or as workaround solution for some of this environments. 
Some docker container which can use this service can be set up as with  running with `--dns` running flag which overwrite default `resolv.conf` 

- What other improvements do you think would be interesting to add to the project?
  
Add some unctional/integration tests to services
Add some configuration options for setting listener ports before start of services and for setting backend dns-tls server, configuration should come from ENV variables
Add more logs for requests, answers and timings


`./build.sh` - for building docker image
`./run.sh` - run docker image on port tcp/udp:8053
`./check.sh` - run checking via `dig`  
