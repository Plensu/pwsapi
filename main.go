package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Hosts struct {
	XMLName string `xml:"nmaprun"`
	Hosts   []Host `xml:"host"`
}

type Host struct {
	Addresses []Address `xml:"address"`
	Ports     []Port    `xml:"ports>port"`
}

type Address struct {
	Addr     string `xml:"addr,attr"`
	AddrType string `xml:"addrtype,attr"`
}

type Port struct {
	Portid  string   `xml:"portid,attr"`
	Proto   string   `xml:"protocol,attr"`
	State   *State   `xml:"state"`
	Service *Service `xml:"service"`
}

type State struct {
	StateAttr  string `xml:"state,attr"`
	ReasonAttr string `xml:"reason,attr"`
}

type Service struct {
	ServName string `xml:"name,attr"`
}

// API URL info for ezpz usage.
var apiBase string = "https://pentest.ws/api/v1"
var apiEng string = "/e"
var apiHost string = "/hosts"
var apiPort string = "/ports"

// Change this to be a string with your api key have it as an environment variable
var apiKey string = os.Getenv("PWSAPIKEY")

// Colored text is important for importance...
var colorReset string = "\033[0m"

var colorRed string = "\033[31m"

// colorGreen := "\033[32m"
// colorYellow := "\033[33m"
// colorBlue := "\033[34m"
// colorPurple := "\033[35m"
// colorCyan := "\033[36m"
// colorWhite := "\033[37m"

func main() {
	// Help flag
	var help bool
	flag.BoolVar(&help, "h", false, "Print the help.")

	//engagement info and args
	var engs bool
	flag.BoolVar(&engs, "engs", false, "List current engagements.")
	var neweng bool
	flag.BoolVar(&neweng, "neweng", false, "Create a new engagement.")
	var eng bool
	flag.BoolVar(&eng, "eng", false, "Get a current engagement using the eid.")
	var ueng bool
	flag.BoolVar(&ueng, "ueng", false, "Update an engagement using the eid.")
	var deleng bool
	flag.BoolVar(&deleng, "deleng", false, "Delete an engagement using the eid.")
	var eid string
	flag.StringVar(&eid, "eid", "", "Engagement ID for commands that require it. ")
	var engName string
	flag.StringVar(&engName, "engName", "", "Name for creating or updating an engagement.")
	var engNotes string
	flag.StringVar(&engNotes, "engNotes", "", "Notes to add to an engagmenet during creation and updating of an engagement.")

	//importing xml args
	var fUpload bool
	flag.BoolVar(&fUpload, "u", false, "Upload nmap or masscan xml file to an engagemnt.")
	var bigfUpload bool
	flag.BoolVar(&bigfUpload, "bu", false, "File to big to upload all at once? Use this to loop through your xml doc and create a host one by one. This will take time!!")
	var fName string
	flag.StringVar(&fName, "file", "", "File for upload")

	// Host info and arguments
	var hosts bool
	flag.BoolVar(&hosts, "hosts", false, "List all hosts for an engagment.")
	var host bool
	flag.BoolVar(&host, "host", false, "Get info for a host for an engagment.")
	var nhost bool
	flag.BoolVar(&nhost, "nhost", false, "Create a new host for an engagment.")
	var uhost bool
	flag.BoolVar(&uhost, "uhost", false, "Update a host for an engagment.")
	var dhost bool
	flag.BoolVar(&dhost, "dhost", false, "Delete a host for an engagment.")
	var hid string
	flag.StringVar(&hid, "hid", "", "Host ID for commands that require it.")
	var hostIP string
	flag.StringVar(&hostIP, "hip", "", "IP of the host.")
	var hostOS string
	flag.StringVar(&hostOS, "hos", "", "OS of the host.")
	var hostType string
	flag.StringVar(&hostType, "htype", "", "Type of host, Ex: server, desktop, phone, printer, etc")
	var hostLabel string
	flag.StringVar(&hostLabel, "hlabel", "", "Label for a host.")
	var hostname string
	flag.StringVar(&hostname, "hname", "", "Host's hostname.")
	var hostShell bool
	flag.BoolVar(&hostShell, "hshell", false, "Set to mark host with a shell.")
	var hostOwned bool
	flag.BoolVar(&hostOwned, "howned", false, "Set to mark a host as owned.")

	//Port arguments
	var ports bool
	flag.BoolVar(&ports, "ports", false, "List all ports for a host.")
	var port bool
	flag.BoolVar(&port, "port", false, "Get info for a port on a host.")
	var nport bool
	flag.BoolVar(&nport, "nport", false, "Create a new port for a host.")
	var uport bool
	flag.BoolVar(&uport, "uport", false, "Update a port for a host.")
	var dport bool
	flag.BoolVar(&dport, "dport", false, "Delete a port for a host.")
	var pid string
	flag.StringVar(&pid, "pid", "", "Port ID for commands that require it.")
	var portNum string
	flag.StringVar(&portNum, "pNum", "", "The number of the port.")
	var portProto string
	flag.StringVar(&portProto, "pProto", "", "Set the protocol, tcp or udp.")
	var portService string
	flag.StringVar(&portService, "pService", "", "Set the port service. Ex: http, ldap, smb, etc.")
	var portVersion string
	flag.StringVar(&portVersion, "pVersion", "", "Set the protocol version if available.")
	var portState string
	flag.StringVar(&portState, "pState", "", "Set the port as open, closed, filtered.")

	//Scratchpad args

	//Note args

	//Credential args

	//Findings args

	//Client args
	var clientID string
	flag.StringVar(&clientID, "clientID", "", "Client ID to tie to an engagement.")

	flag.Parse()

	if apiKey == "" {
		fmt.Println(string(colorRed) + "You need to set your API key in an environment variable called PWSAPIKEY." + string(colorReset))
		fmt.Println("Unix: 'export PWSAPIKEY=youapikey'")
		fmt.Println("Windows: 'setx PWSAPIKEY \"yourapikey\"'")
		fmt.Println("Other: Hard code the api key into the code and rebuild.")
		os.Exit(1)
	}

	if !engs && !neweng && !eng && !ueng && !deleng && !fUpload && !bigfUpload && !hosts && !host && !nhost && !uhost && !dhost && !ports && !port && !nport && !uport && !dport {
		fmt.Println("You forgot to specify the task you wish to complete.")
		printHelp()
		os.Exit(2)
	}

	if help {
		printHelp()
		os.Exit(0)
	}

	if engs {
		listEngagements()
		os.Exit(0)
	}

	if neweng {
		if engName != "" {
			newEngagement(engName, clientID, engNotes)
			os.Exit(0)
		} else {
			fmt.Println("You forgot to add the engagement name. 'pwsapi -neweng -engName NAMESTRING'")
		}
	}

	if eng {
		if eid != "" {
			getEngagement(eid)
			os.Exit(0)
		} else {
			fmt.Println("You forgot to include the EID. 'pwsapi -eng -eid EIDString'")
		}
	}

	if ueng {
		if eid != "" {
			updateEngagement(eid, engName, clientID, engNotes)
			os.Exit(0)
		} else {
			fmt.Println("You forgot to include the EID. 'pwsapi -ueng -eid EIDString'")
		}
	}

	if deleng {
		if eid != "" {
			deleteEngagement(eid)
			os.Exit(0)
		} else {
			fmt.Println("You forgot to include the EID. 'pwsapi -deleng -eid EIDString'")
		}
	}

	if fUpload {
		if eid != "" || fName != "" {
			fileUpload(eid, fName)
			os.Exit(0)
		} else {
			if eid == "" {
				fmt.Println("You forgot to inclide the EID. 'pwsapi -u -eid EIDString -file " + fName + "'")
			}
			if fName == "" {
				fmt.Println("You forgot to inclide the file for upload. 'pwsapi -u -eid " + eid + " -file fileNameOrPath'")
			}
		}
	}

	if bigfUpload {
		if eid != "" || fName != "" {
			bigUpload(eid, fName)
			os.Exit(0)
		} else {
			if eid == "" {
				fmt.Println("You forgot to inclide the EID. 'pwsapi -bu -eid EIDString -file " + fName + "'")
			}
			if fName == "" {
				fmt.Println("You forgot to inclide the file for upload. 'pwsapi -bu -eid " + eid + " -file fileNameOrPath'")
			}
		}
	}

	if hosts {
		if eid != "" {
			listHosts(eid)
			os.Exit(0)
		} else {
			fmt.Println("You forgot to include the EID. 'pwsapi -hosts -eid EIDString'")
		}
	}

	if host {
		if hid != "" {
			getHost(hid)
			os.Exit(0)
		} else {
			fmt.Println("You forgot to include the HID. 'pwsapi -host -hid HIDString'")
		}
	}

	//
	// New host requires EID and a target string (ip).
	// New host can take host type (ex: server), os_type, os, label, hostname(s), notes, board_id, flagged (t/f), reviewed, thumbs_up, thumbs_down, out_of_scope, shell, and owned
	//
	if nhost {
		if eid != "" && hostIP != "" {
			newHost(eid, hostIP, hostOS, hostType, hostLabel, hostname, hostShell, hostOwned)
			os.Exit(0)
		} else {
			if eid == "" {
				fmt.Println("You forgot to include the EID.")
			}

			if hostIP == "" {
				fmt.Println("You forgot to include the Host IP.")
			}

		}
	}

	//
	// Update host requires HID and can take the same inputs as nhost.
	//
	if uhost {
		if hid != "" {
			updateHost(hid, hostIP, hostOS, hostType, hostLabel, hostname, hostShell, hostOwned)
			os.Exit(0)
		} else {
			fmt.Println("You forgot to include the HID. 'pwsapi -uhost -hid HIDString'")
		}
	}

	if dhost {
		if hid != "" {
			deleteHost(hid)
			os.Exit(0)
		} else {
			fmt.Println("You forgot to include the HID. 'pwsapi -dhost -hid HIDString'")
		}
	}

	if ports {
		if hid != "" {
			listPorts(hid)
			os.Exit(0)
		} else {
			fmt.Println("You forgot to include the HID. 'pwsapi -ports -hid HIDString'")
		}
	}

	if port {
		if pid != "" {
			getPort(pid)
			os.Exit(0)
		} else {
			fmt.Println("You forgot to include the PID. 'pwsapi -port -pid PIDString'")
		}
	}

	//
	// New port requires HID and Port number (ex:443)
	// New port can also take proto, service, version, status, state, notes, and checklist
	//
	if nport {
		if hid != "" && portNum != "" {
			newPort(hid, portNum, portProto, portService, portVersion, portState)
			os.Exit(0)
		} else {
			if hid == "" {
				fmt.Println("You forgot to add the Host ID.")
			}

			if portNum == "" {
				fmt.Println("You forgot to add the Port Number")
			}

		}
	}

	if uport {
		if pid != "" {
			updatePort(pid, portNum, portProto, portService, portVersion, portState)
			os.Exit(0)
		} else {
			fmt.Println("blah blah blah")
		}
	}

	if dport {
		if pid != "" {
			deletePort(pid)
			os.Exit(0)
		} else {
			fmt.Println("You forgot to include the PID. 'pwsapi -dport -pid PIDString'")
		}
	}

}

func PrintDefaults() {
	printHelp()
}

func printHelp() {
	fmt.Println("## PWSAPI Help ##")
	fmt.Println("~~ Engagement Arguments ~~")
	fmt.Println(" -engs		List current engagements.")
	fmt.Println(" -eng 		Get engagement info. Requires -eid.")
	fmt.Println(" -neweng 	Create a new engagement. Requires -engName.")
	fmt.Println(" -ueng		Update an engagement. Requires -eid")
	fmt.Println(" -deleng 	Delete an engagement. Requires -eid")
	fmt.Println(" -eid 		Used to set engagement ID.")
	fmt.Println(" -engName 	Used to set engagement name.")
	fmt.Println(" -engNotes	Uses to add notes to an engagement.")
	fmt.Println("")
	fmt.Println("~~ File Arguments ~~")
	fmt.Println(" -u 		Upload nmap of masscan xml file. Requires -eid.")
	fmt.Println(" -bu 		Big file upload. May take some time! Requires -eid.")
	fmt.Println(" -file		File name or path to file you want to upload.")
	fmt.Println("")
	fmt.Println("~~ Host Arguments ~~")
	fmt.Println(" -hosts 	List all hosts for an engagement. Requires -eid.")
	fmt.Println(" -host		List information for a single host. Requires -hid.")
	fmt.Println(" -nhost 	Create a new host. Requires -eid and -hip")
	fmt.Println(" -uhost 	Update a host. Requires -hid.")
	fmt.Println(" -dhost 	Delete a host. Requires -hid.")
	fmt.Println(" -hid 		Used to set host ID.")
	fmt.Println(" -hip 		Used to set host IP.")
	fmt.Println(" -hos 		Used to set host OS.")
	fmt.Println(" -htype 	Used to set host type.")
	fmt.Println(" -hlabel 	Used to set host label.")
	fmt.Println(" -hname 	Used to set host hostname.")
	fmt.Println(" -hshell 	Used to mark host with shell.")
	fmt.Println(" -howned 	Used to mark host as owned.")
	fmt.Println("")
	fmt.Println("~~ Port Arguments ~~")
	fmt.Println(" -ports 	List all ports for a host. Requires -hid.")
	fmt.Println(" -port		List information on a single port. Requires -pid")
	fmt.Println(" -nport 	Create a new port. Requires -hid and -pNum")
	fmt.Println(" -uport 	Update a port for a host. Requires -pid.")
	fmt.Println(" -dport 	Delete a port. Requires -pid.")
	fmt.Println(" -pid 		Used to set port ID.")
	fmt.Println(" -pNum		Used to set the port number.")
	fmt.Println(" -pProto 	Used to set port protocol. tcp or udp")
	fmt.Println(" -pService	Used to set the port service. Ex: http, ldap, smb, etc.")
	fmt.Println(" -pVersion	Used to set the service version if available.")
	fmt.Println(" -pState 	Set the port state as open, closed, or filtered.")
	fmt.Println("")
}

// GET https://pentest.ws/api/v1/e
func listEngagements() {
	url := apiBase + apiEng
	method := "GET"
	var payload []byte = nil
	makeRequest(url, method, payload)
}

// GET https://pentest.ws/api/v1/e/{eid}
func getEngagement(eid string) {
	url := apiBase + apiEng + "/" + eid
	method := "GET"
	var payload []byte = nil
	makeRequest(url, method, payload)
}

// POST https://pentest.ws/api/v1/e
func newEngagement(engName string, clientID string, engNotes string) {
	url := apiBase + apiEng
	method := "POST"
	payload, err := json.Marshal(map[string]string{
		"name":      engName,
		"notes":     engNotes,
		"client_id": clientID,
	})

	if err != nil {
		fmt.Println(err)
	}

	makeRequest(url, method, payload)
}

// PUT https://pentest.ws/api/v1/e/{eid}
func updateEngagement(eid string, engName string, clientID string, engNotes string) {
	url := apiBase + apiEng + "/" + eid
	method := "PUT"
	payload, err := json.Marshal(map[string]string{
		"name":      engName,
		"notes":     engNotes,
		"client_id": clientID,
	})

	if err != nil {
		fmt.Println(err)
	}

	makeRequest(url, method, payload)
}

// DELETE https://pentest.ws/api/v1/e/{eid}
func deleteEngagement(eid string) {
	url := apiBase + apiEng + "/" + eid
	method := "DELETE"
	var payload []byte = nil
	makeRequest(url, method, payload)
}

// POST https://pentest.ws/api/v1/e/{eid}/import/{type}
func fileUpload(eid string, fName string) {
	url := apiBase + apiEng + "/" + eid + "import/nmap?openOnly=off&skipPortless=off&setTargetTo=ip"
	client := &http.Client{}
	file, err := os.ReadFile(fName)

	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(file))

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("accept", "text/plain")
	req.Header.Set("X-API-KEY", apiKey)
	req.Header.Set("Content-Type", "multipart/form-data")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	sb := string(body)
	fmt.Println(sb)
}

// This is actually going to loop through stuff and use multiple API requests.
func bigUpload(eid string, fName string) {

}

// GET https://pentest.ws/api/v1/e/{eid}/hosts
func listHosts(eid string) {
	url := apiBase + apiEng + "/" + eid + apiHost
	method := "GET"
	var payload []byte = nil
	makeRequest(url, method, payload)
}

// GET https://pentest.ws/api/v1/hosts/{hid}
func getHost(hid string) {
	url := apiBase + apiHost + "/" + hid
	method := "GET"
	var payload []byte = nil
	makeRequest(url, method, payload)
}

// POST https://pentest.ws/api/v1/e/{eid}/hosts
func newHost(eid string, hostIP string, hostOS string, hostType string, hostLabel string, hostname string, hostShell bool, hostOwned bool) {
	url := apiBase + apiEng + "/" + eid + apiHost
	method := "POST"
	var shell string
	if hostShell {
		shell = "true"
	} else {
		shell = "false"
	}
	var owned string
	if hostOwned {
		owned = "true"
	} else {
		owned = "false"
	}

	payload, err := json.Marshal(map[string]string{
		"target":    hostIP,
		"type":      hostType,
		"os":        hostOS,
		"label":     hostLabel,
		"hostnames": hostname,
		"shell":     shell,
		"owned":     owned,
	})

	if err != nil {
		fmt.Println(err)
	}

	makeRequest(url, method, payload)
}

// PUT https://pentest.ws/api/v1/hosts/{hid}
func updateHost(hid string, hostIP string, hostOS string, hostType string, hostLabel string, hostname string, hostShell bool, hostOwned bool) {
	url := apiBase + apiHost + "/" + hid
	method := "PUT"
	var shell string
	if hostShell {
		shell = "true"
	} else {
		shell = "false"
	}
	var owned string
	if hostOwned {
		owned = "true"
	} else {
		owned = "false"
	}

	payload, err := json.Marshal(map[string]string{
		"target":    hostIP,
		"type":      hostType,
		"os":        hostOS,
		"label":     hostLabel,
		"hostnames": hostname,
		"shell":     shell,
		"owned":     owned,
	})

	if err != nil {
		fmt.Println(err)
	}

	makeRequest(url, method, payload)
}

// DELETE https://pentest.ws/api/v1/hosts/{hid}
func deleteHost(hid string) {
	url := apiBase + apiHost + "/" + hid
	method := "DELETE"
	var payload []byte = nil
	makeRequest(url, method, payload)
}

// GET https://pentest.ws/api/v1/hosts/{hid}/ports
func listPorts(hid string) {
	url := apiBase + apiHost + "/" + hid + apiPort
	method := "GET"
	var payload []byte = nil
	makeRequest(url, method, payload)
}

// GET https://pentest.ws/api/v1/ports/{pid}
func getPort(pid string) {
	url := apiBase + apiPort + "/" + pid
	method := "GET"
	var payload []byte = nil
	makeRequest(url, method, payload)
}

// POST https://pentest.ws/api/v1/hosts/{hid}/ports
func newPort(hid string, portNum string, portProto string, portService string, portVersion string, portState string) {
	url := apiBase + apiHost + "/" + hid + apiPort
	method := "POST"

	payload, err := json.Marshal(map[string]string{
		"port":    portNum,
		"proto":   portProto,
		"service": portService,
		"version": portVersion,
		"state":   portState,
	})

	if err != nil {
		fmt.Println(err)
	}

	makeRequest(url, method, payload)
}

// PUT https://pentest.ws/api/v1/ports/{pid}
func updatePort(pid string, portNum string, portProto string, portService string, portVersion string, portState string) {
	url := apiBase + apiPort + "/" + pid
	method := "PUT"

	payload, err := json.Marshal(map[string]string{
		"port":    portNum,
		"proto":   portProto,
		"service": portService,
		"version": portVersion,
		"state":   portState,
	})

	if err != nil {
		fmt.Println(err)
	}

	makeRequest(url, method, payload)
}

// DELETE https://pentest.ws/api/v1/ports/{pid}
func deletePort(pid string) {
	url := apiBase + apiPort + "/" + pid
	method := "DELETE"
	var payload []byte = nil
	makeRequest(url, method, payload)
}

func makeRequest(url string, reqMethod string, payload []byte) {
	client := &http.Client{}
	req, err := http.NewRequest(reqMethod, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("X-API-KEY", apiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	sb := string(body)
	fmt.Println(sb)

}
