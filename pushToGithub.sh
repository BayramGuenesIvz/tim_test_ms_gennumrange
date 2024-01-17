#!/bin/bash

if  [[ "$1" == "" ]]
then
  echo "Bitte git-commit text angeben."
  exit 0;
fi

declare FILE=./assets/DockerFileVariants/DockerfileGithub
declare existsProjectFilesys="false";
[ -f $FILE ] && existsProjectFilesys="true"

if  [[ "${existsProjectFilesys}" == "false" ]]; then 
    echo "Bitte DockerfileGithub unter ./assets/DockerFileVariants/ anlegen"
    exit 0
fi

rm Dockerfile
cp ./assets/DockerFileVariants/DockerfileGithub Dockerfile
git add .
git commit -a -m "$1"
git push
echo "Erfolgreich nach Github gepushed."

rm Dockerfile
cp ./assets/DockerFileVariants/DockerLoclDev Dockerfile
