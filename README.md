# 0-Microservice in 5 minutes :- June 8th 2016

## Build Status:  
[![CircleCI](https://circleci.com/gh/nicholasjackson/containersched.svg?style=svg)](https://circleci.com/gh/nicholasjackson/containersched)

## Gitter:  
[![Join the chat at https://gitter.im/nicholasjackson/minke](https://badges.gitter.im/nicholasjackson/minke.svg)](https://gitter.im/nicholasjackson/minke?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## Video:
[https://skillsmatter.com/skillscasts/8097-0-to-microservice-in-5-minutes](https://skillsmatter.com/skillscasts/8097-0-to-microservice-in-5-minutes)

## How to build - Pure Docker
Whilst having ruby installed on your dev machine can be handy there is no requirement to install this or any other dependency than Docker.  You can use the simple bash script which starts a Ruby docker container to run the build and test script, because this maps a reference to the docker sock on your local machine any commands to start docker container from within the ruby container will be proxied to your local machine, a docker server is not running inside the Ruby container.

### NOTE: Beta version of docker for mac, additional step
If you use Docker for Mac then the ip address for your docker host will appear to be 127.0.0.1, you can not use this ip to resolve resources when inside a container as localhost is a different localhost.  To get around this alias localhost on your mac to a different ip address and then reference this as your docker host in the build script.
```bash
$ sudo ifconfig lo0 alias 10.254.254.254 # then use like $ ./build.sh build_image -i 10.254.254.254
```

### Step 1: Build a Docker image
```bash
$ ./build.sh build_image -i 192.168.99.100 # Where 192.168.99.100 is the ip address of your docker host
```

### Step 2: Run the functional tests
```bash
$ ./build.sh cucumber -i 192.168.99.100 # Where 192.168.99.100 is the ip address of your docker host
```

## How to build - With Ruby installed
### Setup
```bash
$ cd _build
$ bundle
```

### Build a Docker image
```
$ rake app:build_image
```

### Run the functional tests
```
$ rake app:cucumber
```

For more build options please see the minke documentation http://nicholasjackson.github.io/minke/  


## Resources:

### A little blog post I wrote on this subject:  
http://nicholasjackson.github.io/microservices/go/building-and-testing-microservices-part1/

### Minke (updated docs coming real soon):  
Docs: http://nicholasjackson.github.io/minke/  
Source:   
- http://github.com/minke  
- https://github.com/nicholasjackson/minke-generator-go

### Contact:
- twitter: @sherrifjackson @NOTHS_Tech
- email: nicholasjackson@notonthehighstreet.com
- linkedin: https://www.linkedin.com/in/jacksonnic

### We are hiring
Want to come and join us on our awesome journey?   
We have a number of vacancies, not just in technology but across the whole company.  Please feel free to either contact me directly or you can browse our open vacancies and apply using the link below.

http://www.notonthehighstreet.com/work-with-us
