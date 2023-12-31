package gemini

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
)

func Request(req string) (string, string) {
	parsed, err := url.Parse(req)
	conn, err := tls.Dial("tcp", parsed.Host+":1965", &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println("Falied to connect: " + err.Error())
		return err.Error(), ""
	}
	defer conn.Close()

	conn.Write([]byte(req + "\r\n"))
	reader := bufio.NewReader(conn)

	responseHeader, err := reader.ReadString('\n')
	parts := strings.Fields(responseHeader)
	status, err := strconv.Atoi(parts[0][0:1])

	switch status {
	case 1, 3, 6:
		fmt.Println("Error")
	
	case 2:
		bodyBytes, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Println("Error reading body")
			return err.Error(), ""
		}
		body := string(bodyBytes)

		return body, parts[1]
	}

	return "", ""
}