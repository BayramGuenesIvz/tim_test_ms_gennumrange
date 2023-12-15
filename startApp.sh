#!/bin/bash

if  [[ "$1" == "" || "$2" == "" || "$3" == "" ]]
then
      echo "Usage startApp.sh <pathSettings> <ServiceDB> <PortDB>" ;
      echo " "
      echo "example call: ./startApp.sh  ~/ZZDevelop/goProjects/src/tim/timFileSys/settings/confignumrange.json localhost 33306"
      exit 0
fi;
export useSettingsPath=$1
export  SERVICEDB=$2
export  PORTDB=$3


declare killproc=$(pidof "tim_test_ms_gennumrange")
echo killproc=$killproc
kill -9 $killproc


rm ./tim_test_ms_gennumrange
CGO_ENABLED=0 GOOS=linux go build ./cmd/tim_test_ms_gennumrange
#go build
sleep 2&&(echo "************************************************************************"\
&&echo "Access to App via http://localhost:7000/NumRangeServices"\
&&echo "************************************************************************")&
  
./tim_test_ms_gennumrange confLocation=$useSettingsPath 

