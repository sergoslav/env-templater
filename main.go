package main

import (
	"bufio"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type yamlConfig struct {
	Configs struct {
		RemoveExtraEnv bool `yaml:"removeExtraEnv"`
	} `yaml:"configs"`
	Environment []struct {
		Name      string            `yaml:"name"`
		Replacers map[string]string `yaml:"replacers"`
	} `yaml:"environment"`
}

func main() {
	//get flags
	sourceFile := flag.String("s", ".env.template", "the name of the source .env file.")
	destFile := flag.String("d", ".env", "the name of the new file that will be created or overwritten.")
	env := flag.String("e", "dev", "environment name")
	configFile := flag.String("c", ".env-templater.yaml", "templater config .yaml file")
	flag.Parse()

	// fmt.Println("t:", *t)
	// fmt.Println("sourceFile:", *sourceFile)
	// fmt.Println("destFile:", *destFile)
	// fmt.Println("env:", *env)
	// fmt.Println("conf:", *configFile)
	// if len(*configFile) < 1000 {
	// 	os.Exit(0)
	// }

	//load yaml config
	configs, err := getConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%+v\n", configs)

	// Открываем исходный файл
	source, err := os.Open(*sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	defer source.Close()

	// Создаем или перезаписываем новый файл
	dest, err := os.Create(*destFile)
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()

	// Проходим по каждой строке исходного файла и записываем измененные строки в новый файл
	scanner := bufio.NewScanner(source)
	for scanner.Scan() {
		line := scanner.Text()
		newLine, ok := getLine(line, configs, *env)
		if ok {
			fmt.Fprintln(dest, newLine)
		}
	}

	// Проверяем наличие ошибок при сканировании исходного файла
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Env file `" + *destFile + "` generated. Env:`" + *env + "`")
}

func getConfig(configFile string) (yamlConfig, error) {
	configs := yamlConfig{}

	// Открываем исходный файл
	// if len(configFile) == 0 {
	// 	configFile = DefaultConfigFile
	// }
	yfile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
		return configs, err
	}

	err = yaml.Unmarshal([]byte(yfile), &configs)
	if err != nil {
		log.Fatalf("error: %v", err)
		return configs, err
	}

	// fmt.Printf("%+v\n", configs)

	return configs, nil
}

func getLine(line string, configs yamlConfig, env string) (string, bool) {
	if len(line) == 0 {
		return line, true
	}

	//process prefix
	for _, eName := range getEnvNames(configs) {
		if strings.HasPrefix(line, eName) {
			if eName == env {
				line = strings.TrimPrefix(line, eName+"-")
			} else {
				return "", false
			}
		}
	}

	//process replace
	for old, new := range getReplacers(configs, env) {
		old = "{{" + old + "}}"
		line = strings.Replace(line, old, new, -1)
	}

	return line, true
}

func getEnvNames(config yamlConfig) []string {
	var names []string
	for _, env := range config.Environment {
		names = append(names, env.Name)
	}
	return names
}

func getReplacers(config yamlConfig, env string) map[string]string {
	for _, cEnv := range config.Environment {
		if cEnv.Name == env {
			return cEnv.Replacers
		}
	}

	log.Fatal("env not found in yamp config")
	return nil
}
