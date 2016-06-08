#!/bin/bash

DOCKER_SOCK="/var/run/docker.sock:/var/run/docker.sock"
BUNDLE_COMMAND="bundle install --path=./vendor/cache && bundle update"
ERROR="Please specify a command and ip address for your docker host: ./bash.sh [build_image, cucumber] -i [DOCKER_IP]"
COMMAND=""

OPTIND=2
if [ "$1" != 'build_image' ] && [ "$1" != 'cucumber' ] ; then
  echo $ERROR;
  exit 1;
fi

COMMAND=$1

while getopts ":i:" opt; do
  case $opt in
    i)
      IP_ADDRESS=$OPTARG >&2
      ;;
    \?)
      echo $ERROR;
      exit 1;
      ;;
    :)
      echo $ERROR;
      exit 1;
      ;;
  esac
done

eval "docker run --rm -it -v $DOCKER_SOCK -v $(pwd):$(pwd) -w $(pwd)/_build -e=DOCKER_IP=$IP_ADDRESS rubydev /bin/bash -c '$BUNDLE_COMMAND && rake app:$COMMAND'"
