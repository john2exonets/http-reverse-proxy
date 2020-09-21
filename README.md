# http-reverse-proxy
A dirt simple HTTP Reverse Proxy

This is a very simple HTTP Reverse Proxy that is written in GO that I use to get around 
a limitation of Terraform's Docker Provider that only allows for ports to be bound
to the first network on a Container.

Simple set the environment variables for where to proxy the connection and away it goes!

The 'revproxy.tf' is a Terraform file that shows how to run the proxy as a Docker Container.
