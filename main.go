package main

import (
	"flag"
	"fmt"
	"os"
)

// API URL info for ezpz usage.
// apiBase := "https://pentest.ws/api/v1"
// apiEng := "/e"
// apiHost := "/hosts"
// apiPort := "/ports"

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
	var eid string
	var fName string
	var hid string
	var pid string

	var engs bool
	var neweng bool
	var eng bool
	var ueng bool
	var deleng bool
	var fUpload bool
	var bigfUpload bool
	var hosts bool
	var host bool
	var nhost bool
	var uhost bool
	var dhost bool
	var ports bool
	var port bool
	var nport bool
	var uport bool
	var dport bool

	var help bool

	flag.BoolVar(&help, "h", false, "Print the help.")

	//engagement args
	flag.BoolVar(&engs, "engs", false, "List current engangements.")
	flag.BoolVar(&neweng, "neweng", false, "Create a new engangement.")
	flag.BoolVar(&eng, "eng", false, "Get a current engangement using the eid.")
	flag.BoolVar(&ueng, "ueng", false, "Update an engagement using the eid.")
	flag.BoolVar(&deleng, "deleng", false, "Delete an engagement using the eid.")
	flag.StringVar(&eid, "eid", "", "Engagement ID for commands that require it. ")

	//importing xml args
	flag.BoolVar(&fUpload, "u", false, "Upload nmap or masscan xml file to an engagemnt.")
	flag.BoolVar(&bigfUpload, "bu", false, "File to big to upload all at once? Use this to loop through your xml doc and create a host one by one. This will take time!!")
	flag.StringVar(&fName, "file", "", "File for upload")

	//Single host arguments
	flag.BoolVar(&hosts, "hosts", false, "List all hosts for an engagment.")
	flag.BoolVar(&host, "host", false, "Get info for a host for an engagment.")
	flag.BoolVar(&nhost, "nhost", false, "Create a new host for an engagment.")
	flag.BoolVar(&uhost, "uhost", false, "Update a host for an engagment.")
	flag.BoolVar(&dhost, "dhost", false, "Delete a host for an engagment.")
	flag.StringVar(&hid, "hid", "", "Host ID for commands that require it.")

	//Port arguments
	flag.BoolVar(&ports, "ports", false, "List all ports for a host.")
	flag.BoolVar(&port, "port", false, "Get info for a port on a host.")
	flag.BoolVar(&nport, "nport", false, "Create a new port for a host.")
	flag.BoolVar(&uport, "uport", false, "Update a port for a host.")
	flag.BoolVar(&dport, "dport", false, "Delete a port for a host.")
	flag.StringVar(&pid, "pid", "", "Port ID for commands that require it.")

	//Scratchpad args

	//Note args

	//Credential args

	//Findings args

	//Client args

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
		newEngagement()
		os.Exit(0)
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
			updateEngagement(eid)
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
		} else {
			fmt.Println("You forgot to include the EID. 'pwsapi -hosts -eid EIDString'")
		}
	}

	if host {
		if hid != "" {
			getHost(hid)
		} else {
			fmt.Println("You forgot to include the HID. 'pwsapi -host -hid HIDString'")
		}
	}

	if nhost {
		if hname != "" {

		}
	}

}

func printHelp() {
	fmt.Println("This will print help at some point.")
}

func listEngagements() {
	fmt.Println("This is working as expected" + apiKey)
}

func getEngagement(eid string) {

}

func newEngagement(engName string) {

}

func updateEngagement(eid string) {

}

func deleteEngagement(eid string) {

}

func fileUpload(eid string, fName string) {

}

func bigUpload(eid string, fName string) {

}

func listHosts(eid string) {

}

func getHost(hid string) {

}

func newHost(hname string) {

}

func updateHost(hid string) {

}

func deleteHost(hid string) {

}

func listPorts(hid string) {

}

func getPort(pid string) {

}

func newPort(port string) {

}

func updatePort(pid string) {

}

func deletePort(pid string) {

}
