package readConfig

import (
	"encoding/json"
	"errors"
	"log" //"strconv"
	"os"

	"github.com/BayramGuenesIvz/tim_test_ms_gennumrange/internal/data"
)

func GetApplConf() (eConf data.ApplConfStruct, err error) {

	data.ApplConf = data.ApplConfStruct{}
	if len(data.ThisConfLocation) == 0 {
		errString :=
			"Achtung !!!  Service konnte nicht gestartet werden." +
				" Bitte Pfad der Konfigurationsdatei confignumrange.json  über Umgebungsparameter" +
				" 'confLocation' bzw als Startargument(osparam) 'confLocation=<path>' angeben."
		println(errString)
		return data.ApplConf, errors.New(errString)
	}

	filePathAndName := data.ThisConfLocation
	file, _ := os.Open(filePathAndName)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data.ApplConf)

	if err != nil {
		println("Decoding confignumrange.json"+" FAILED", err.Error())
		log.Fatal(err.Error() + ". Möglicherweise fehlt die confignumrange.json Datei im Pfad " + data.ThisConfLocation)
		return data.ApplConf, err
	}
	eConf = data.ApplConf
	if !data.IsDBInfoFromEnv {
		if len(data.ApplConf.SERVICEDB) > 0 {
			data.ServiceDB = data.ApplConf.SERVICEDB
		}
		if len(data.ApplConf.PORTDB) > 0 {
			data.PortDB = data.ApplConf.PORTDB
		}
	}
	return eConf, nil

}
