# SW Komponente tim_test_ms_gennumrange Zweck/Aufgabe

* Webserver(Backend+Frontend) zum Testen des Packages tim_utils_numrange

Restriktionen
-------------------
* Nur für Tests der Entwicklungs- und Testplattformen

Betrieb/Start
-------------

````
go build ./cmd/tim_test_ms_gennumrange
./tim_test_ms_gennumrange confLocation=./config 
````
oder
````
./startApp
````
bzw
````
./startApp.sh <pathSettings> <ServiceDB> <PortDB>"  
````
Beispiel
````
 ./startApp.sh  ~/ZZDevelop/goProjects/src/tim_presse/timFileSys/settings localhost 33306"´
 ````

 Einstellungen 
--------------------------
\<pathSettings\>/configtestnumrange.json  

Beispiel unter tim_test_ms_gennumrange/config/configtestnumrange.json