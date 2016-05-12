package common

/*
加载资源配置文件
*/
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var configMap map[string]string = make(map[string]string)

type Config struct {
	files []string
}

func NewConfig(files []string) *Config {
	config := &Config{files}
	return config
}
func init() {

}
func (config *Config) LoadConfig() {
	if len(config.files) == 0 {
		fmt.Println("not found setting files")
		return
	}
	for _, v := range config.files {
		var configPath string = GetBasePath() + v
		file, err := os.Open(configPath)
		if err == nil {
			reader := bufio.NewReader(file)
			for {
				byt, _, readErr := reader.ReadLine()
				if readErr == io.EOF {
					break
				}
				str := string(byt)
				split := strings.Index(str, "=")
				configMap[str[0:split]] = str[split+1:]
			}
		}
		defer func() {
			file.Close()
		}()
	}
	fmt.Println(configMap)
}
func Get(key string) string {
	return configMap[key]
}
