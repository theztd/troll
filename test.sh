#!/bin/bash

cat deploy.nomad | sed "s/__JOB_NAME__/troll-test/g" | nomad run -