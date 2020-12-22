/*
 * Copyright (C) 2020 The ocluso Authors
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"
)

type builderConfig struct {
	Modules []string `json:"modules"`
}

//TODO: More verbose output?
func main() {
	sourcetreePath := readSourcetreePath()

	clean(sourcetreePath)

	config := readConfig(sourcetreePath)
	validateConfig(config, sourcetreePath)

	createBuildDirectories(sourcetreePath)

	generateGoSource(config, sourcetreePath)
	generateMakefile(config, sourcetreePath)
}

func handleError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func clean(sourcetreePath string) {
	err := os.RemoveAll(path.Join(sourcetreePath, "build"))
	handleError(err)

	err = os.RemoveAll(path.Join(sourcetreePath, "Makefile"))
	handleError(err)
}

func createBuildDirectories(sourcetreePath string) {
	buildPath := path.Join(sourcetreePath, "build")
	permissions := os.FileMode(0775)

	err := os.Mkdir(buildPath, permissions)
	handleError(err)

	subdirectories := []string{"generated"}
	for _, subdirectory := range subdirectories {
		err = os.Mkdir(path.Join(buildPath, subdirectory), os.ModePerm)
		handleError(err)
	}
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func generateGoSource(config builderConfig, sourcetreePath string) {
	outputPath := path.Join(sourcetreePath, "build/generated")

	renderTemplate("templates/go.mod.tmpl", path.Join(outputPath, "go.mod"), config)
	renderTemplate("templates/main.go.tmpl", path.Join(outputPath, "main.go"), config)
}

func generateMakefile(config builderConfig, sourcetreePath string) {
	renderTemplate("templates/Makefile.tmpl", path.Join(sourcetreePath, "Makefile"), config)
}

func readConfig(sourcetreePath string) builderConfig {
	config := builderConfig{}

	configBytes, err := ioutil.ReadFile(path.Join(sourcetreePath, "builder-config.json"))
	handleError(err)

	err = json.Unmarshal(configBytes, &config)
	handleError(err)

	return config
}

func readSourcetreePath() string {
	if len(os.Args) != 2 {
		fmt.Println("\nUsage: ocluso-builder PATH-TO-SOURCETREE")
		os.Exit(1)
	}

	return os.Args[1]
}

func renderTemplate(templatePath string, outputPath string, data interface{}) {
	template, err := template.ParseFiles(templatePath)
	handleError(err)

	f, err := os.Create(outputPath)
	handleError(err)

	err = template.Execute(f, data)
	handleError(err)
}

func validateConfig(config builderConfig, sourcetreePath string) {
	modulesPath := path.Join(sourcetreePath, "modules")

	for _, module := range config.Modules {
		moduleExists := fileExists(path.Join(modulesPath, module, "module.json"))
		if !moduleExists {
			log.Fatalln("Module", module, "does not exist")
		}
	}
}
