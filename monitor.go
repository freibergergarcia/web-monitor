package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const sleepTime = 5

func main() {

	whoAreYou()
	chooseAnOption()

	action := readOption()
	sites := getSites()

	switch action {
	case 1:
		startMonitor(sites)
	case 2:
		printLog()
	case 0:
		fmt.Println("Bye!")
		os.Exit(0)
	default:
		fmt.Println("Couldn't find the option", action)
		os.Exit(-1)
	}
}

func whoAreYou() {
	fmt.Println("Who's monitoring?")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hey", name, "choose an option")
}

func chooseAnOption() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func readOption() int {
	var action int
	/* another option to read the input, + verbose and + ~complicated */
	fmt.Scanf("%d", &action)
	fmt.Println("Option chosen:", action)

	return action
}

func startMonitor(sites []string) {
	fmt.Println("Monitoring...")

	for i := 0; i < monitoring; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(sleepTime * time.Second)
		fmt.Println("Waiting", sleepTime, "seconds...")
	}
}

func getSites() []string {
	sites := readSitesFromFile()

	return sites
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "is up and returned a", resp.StatusCode)
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "returned an unsuccessfull status code", resp.StatusCode)
		registerLog(site, false)
	}
}

func readSitesFromFile() []string {
	var sites []string

	sitesConf, err := os.Open("conf/sites.txt")
	if err != nil {
		fmt.Println("An error has ocurred", err)
	}

	reader := bufio.NewReader(sitesConf)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}
	sitesConf.Close()

	return sites
}

func registerLog(site string, status bool) {
	log, err := os.OpenFile("logs/log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Something went wrong", err)
	}
	log.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + site + " - online: " + strconv.FormatBool(status) + "\n")

	log.Close()
}

func printLog() {
	fmt.Println("Showing logs...")
	fileLog, err := ioutil.ReadFile("logs/log.txt")

	if err != nil {
		fmt.Println("Couldn't open the file", err)
	}
	fmt.Println(string(fileLog))
}
