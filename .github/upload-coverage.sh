#!/bin/bash
apt-get update; apt-get -y install curl git
curl -s https://codecov.io/bash | bash
