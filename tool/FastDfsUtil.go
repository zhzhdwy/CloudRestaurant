package tool

import (
	"bufio"
	"fmt"
	"github.com/tedcy/fdfs_client"
	"os"
	"strings"
)

func UploadFile(fileName string) string {
	client, err := fdfs_client.NewClientWithConfig("./config/fastdfs.conf")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	fileId, err := client.UploadByFilename(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return fileId
}

func FileServerAddr() string {
	file, err := os.Open("./config/fastdfs.conf")
	if err != nil {
		return ""
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		str := strings.SplitN(line, "=", 2)
		switch str[0] {
		case "http_server_port":
			return str[1]
		}
		if err != nil {
			return ""
		}
	}
}
