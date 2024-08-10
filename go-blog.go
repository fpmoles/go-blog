package go_blog

import (
	"fmt"
	"os"

	"github.com/gomarkdown/markdown/md"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Name string `yaml:"name"`
	Site string `yaml:"site"`
}

func (c Config) isValid() bool {
	if c.Name == "" || c.Site == "" {
		return false
	}
	return true
}

type GoBlogError struct {
	cause string
}

func NewGoBlogError(cause string) *GoBlogError {
	return &GoBlogError{
		cause: cause,
	}
}
func (e GoBlogError) Error() string {
	return fmt.Sprintf("internal error caused by %s", e.cause)
}

type Project struct {
	renderer *md.Renderer
}

func main() {
	//TODO: Implement logic https://github.com/fpmoles/go-blog/issues/2
}

func loadConfig(configFile string) (*Config, error) {
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	if !config.isValid() {
		return nil, NewGoBlogError("config is not valid")
	}
	return &config, nil
}

func renderMarkdown(file *os.File, project *Project) {

}
