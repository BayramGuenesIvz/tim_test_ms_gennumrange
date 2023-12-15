package readSettings

import (
	"flag"
	"os"
	"strings"

	"github.com/BayramGuenesIvz/tim_test_ms_gennumrange/internal/data"
)

func LoadExternalSettings() {

	loadFromOSFlags()
	// Überschreiben OS Flags durch ENV Param
	loadFromOSEnv()
	// Überschreiben Parameter, falls durch args mitgegeben
	loadFromOSArgs()
	//println("ThisPort:" + ThisPort)
	//	println("ServiceDB, PortDB:" + ServiceDB + ":" + PortDB)
	return
}

func loadFromOSFlags() {
	protokollPointer := flag.String("protokoll", "http", "application-communication protokoll")
	servLocaIDPointer := flag.String("serviceLocationID", "127.0.0.1", "service location id")
	portPointer := flag.String("port", "7000", "application port")
	servDBPointer := flag.String("serviceDB", "127.0.0.1", "service location id DBServer")
	portDBPointer := flag.String("DBport", "3306", "DB port")

	//portPresServerPointer := flag.String("PortPresentation", "7500", "presentation port")
	confLocationPointer := flag.String("confLocation", "", "confPSSysCompReg.json path")

	flag.Parse()
	protokoll := *protokollPointer
	port := *portPointer
	confLocation := *confLocationPointer
	//portPresServer := *portPresServerPointer
	servLocaID := *servLocaIDPointer
	data.ServiceDB = *servDBPointer
	data.PortDB = *portDBPointer
	data.ThisCommunnicationProtokoll = protokoll
	data.ThisPort = port
	data.ThisConfLocation = confLocation
	//data.PortPresentation = portPresServer
	data.ServiceLocationID = servLocaID
	if data.ThisCommunnicationProtokoll == "http" {
		data.ThisServerAddr = data.ThisCommunnicationProtokoll + "://" + data.ServiceLocationID + ":" + data.ThisPort
	}
	return
}

func loadFromOSEnv() {

	servLocaID, protokoll, port, servDB, portDB, confLocation := //, portPresServer :=
		data.ThisCommunnicationProtokoll,
		data.ServiceLocationID, data.ThisPort, data.ServiceDB, data.PortDB, data.ThisConfLocation //, data.PortPresentation
	environList := os.Environ()
	leng := len(environList)
	for i := 0; i < leng; i++ {
		if i > 0 {
			osenvparamval := environList[i]
			splittedString := strings.Split(osenvparamval, "=")
			paramname := splittedString[0]
			paramval := splittedString[1]

			if paramname == "SVCTimGenNumRange" || paramname == "serviceLocationIDSysReg" || paramname == "SERVICELOCATIONIDSYSREG" {
				servLocaID = paramval
			}

			if paramname == "PortTimGenNumRange" ||
				paramname == "port" || paramname == "PORT" || paramname == "Port" ||
				paramname == "port_sysreg" || paramname == "PORT_SYSREG" || paramname == "Port_Sysreg" {
				port = paramval
			}

			if paramname == "serviceDB" || paramname == "SERVICEDB" {
				servDB = paramval
				data.IsDBInfoFromEnv = true
			}

			if paramname == "portDB" || paramname == "PORTDB" || paramname == "PortDB" {
				portDB = paramval
				data.IsDBInfoFromEnv = true
			}

			if paramname == "protokoll" || paramname == "Protokoll" || paramname == "PROTOKOLL" {
				protokoll = paramval
			}
			if paramname == "configfile" || paramname == "confLocation" || paramname == "ConfLocation" ||
				paramname == "conflocation" || paramname == "Conflocation" || paramname == "CONFLOCATION" {
				confLocation = paramval
			}
			if paramname == "portpresentation" || paramname == "portPresentation" || paramname == "PortPresentation" || paramname == "PORTPRESENTATION" {
				//	portPresServer = paramval
			}

		}
	}
	data.ServiceLocationID,
		data.ThisCommunnicationProtokoll, data.ThisPort, data.ServiceDB, data.PortDB, data.ThisConfLocation = //, data.PortPresentation
		servLocaID, protokoll, port, servDB, portDB, confLocation //, portPresServer
	if data.ThisCommunnicationProtokoll == "http" {
		data.ThisServerAddr = data.ThisCommunnicationProtokoll + "://" + data.ServiceLocationID + ":" + data.ThisPort
	}
}

func loadFromOSArgs() {
	servLocaID, protokoll, port, servDB, portDB, confLocation := //, portPres :=
		data.ServiceLocationID, data.ThisCommunnicationProtokoll,
		data.ThisPort, data.ServiceDB, data.PortDB, data.ThisConfLocation //, data.PortPresentation

	leng := len(os.Args)
	for i := 0; i < leng; i++ {
		if i > 0 {
			osparam := os.Args[i]
			splittedString := strings.Split(osparam, "=")
			var namevalues []string
			namevalues = append(namevalues, splittedString...)
			leng := len(namevalues)
			if leng > 1 {
				paramname := splittedString[0]
				paramval := splittedString[1]

				if paramname == "SVCTimGenNumRange" || paramname == "serviceLocationIDSysReg" || paramname == "SERVICELOCATIONIDSYSREG" {
					servLocaID = paramval
				}
				if paramname == "PortTimGenNumRange" || paramname == "port" || paramname == "PORT" || paramname == "Port" ||
					paramname == "port_sysreg" || paramname == "PORT_SYSREG" || paramname == "Port_Sysreg" {
					port = paramval
				}
				if paramname == "protokoll" || paramname == "Protokoll" || paramname == "PROTOKOLL" {
					protokoll = paramval
				}
				if paramname == "serviceDB" || paramname == "SERVICEDB" {
					servDB = paramval
					data.IsDBInfoFromEnv = true
				}

				if paramname == "portDB" || paramname == "PORTDB" || paramname == "PortDB" {
					portDB = paramval
					data.IsDBInfoFromEnv = true
				}
				if paramname == "configfile" || paramname == "confLocation" || paramname == "ConfLocation" ||
					paramname == "conflocation" || paramname == "Conflocation" || paramname == "CONFLOCATION" {
					confLocation = paramval
				}
				if paramname == "portpresentation" || paramname == "portPresentation" || paramname == "PortPresentation" || paramname == "PORTPRESENTATION" {
					//portPres = paramval
				}

			}

		}
	}
	data.ServiceLocationID, data.ThisCommunnicationProtokoll, data.ThisPort, data.ServiceDB, data.PortDB, data.ThisConfLocation = //, data.PortPresentation =
		servLocaID, protokoll, port, servDB, portDB, confLocation //, portPres
	if data.ThisCommunnicationProtokoll == "http" {
		data.ThisServerAddr = data.ThisCommunnicationProtokoll + "://" + data.ServiceLocationID + ":" + data.ThisPort
	}
}
