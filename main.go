package main

import (
	"GauGau/config"
	"GauGau/params"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/otiai10/copy"
)

var cfg *config.Config
var currentPath string

func init() {
	currentPath, _ = os.Getwd()
	cfg = config.LoadConfig()
}

func main() {
	fmt.Println("---------------------------BEGIN---------------------------")
	Load()
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("facebook: https://www.facebook.com/minhnguyen29041998/")
	fmt.Println("github: https://github.com/mjnk9xw")
	fmt.Println("---------------------------DONE----------------------------")
}

func Load() {
	//  check layer
	if cfg.Model == "" {
		cfg.Model = string(params.CleanCode)
	}

	if cfg.Type == "" {
		cfg.Type = string(params.APIGIN)
	}

	if cfg.Config == "" {
		cfg.Config = string(params.Json)
	}

	checkModel()
	checkType()
	checkDB()
	checkConfig()
	checkEntities()
	checkUsecases()

	exec.Command("cmd.exe", "/c", "cd", cfg.Path,
		"/c", "go mod init", cfg.ProjectName,
		"/c", "go mod tidy")
}

func checkModel() {

	makeFolder(cfg.Path)

	switch cfg.Model {
	case string(params.CleanCode):
		copy.Copy(filepath.Join(currentPath, "agen", cfg.Model), cfg.Path)
	case string(params.Mvc):
		copy.Copy(filepath.Join(currentPath, "agen", cfg.Model), cfg.Path)
	case string(params.Th3Layer):
		copy.Copy(filepath.Join(currentPath, "agen", cfg.Model), cfg.Path)
	}

}

func checkType() {
	switch cfg.Type {
	case string(params.APIGIN):
	case string(params.APIEcho):
	case string(params.APIMux):
	}
}

func checkDB() {
	switch cfg.Db {
	case string(params.Redis):
	case string(params.Mongo):
	case string(params.Mysql):
	}
}

func checkConfig() {

	cfgs := strings.Split(cfg.Config, ".")
	if len(cfgs) < 2 {
		panic(errors.New("file config error"))
	}

	pathConfig := filepath.Join(cfg.Path, "configs", "config.go")
	fmt.Println(pathConfig)
	code := readFile(pathConfig)
	strings.Replace(code, "{{CONFIG_TYPE}}", cfgs[1], -1)
	strings.Replace(code, "{{CONFIG_NAME}}", cfgs[0], -1)

	fmt.Println(code)

	writeFile(pathConfig, code)

	makeFileWrite(filepath.Join(cfg.Path, cfg.Config), `{"demo": "huhu"}`)
}

func checkEntities() {

	// entities + repository interface + body request
	entities := strings.Split(cfg.Entities, ",")
	for _, v := range entities {

		// model
		makeFileWrite(filepath.Join(cfg.Path, "entities", strings.ToLower(v)+".go"),
			fmt.Sprintf(`package entities
type %s struct {}`, v))

		// repository
		makeFileWrite(filepath.Join(cfg.Path, "repository", strings.ToLower(v)+".go"),
			fmt.Sprintf(`package repository
import "%s/entities"
type %s interface{ 
	Creat(e *entities.%s) (*entities.%s, error)
	Update(id int64,e *entities.%s) (*entities.%s, error)
	Get(id int64) (*entities.%s, error)
	Delete(id int64) (error)
}`, cfg.ProjectName, v, v, v, v, v, v))

		// body
		makeFileWrite(filepath.Join(cfg.Path, "api", "requests", strings.ToLower(v)+".go"),
			fmt.Sprintf(`package requests
type %sReq struct{ 
}`, v))

		// res
		makeFileWrite(filepath.Join(cfg.Path, "api", "responses", strings.ToLower(v)+".go"),
			fmt.Sprintf(`package responses
type %sRes struct{ 
}`, v))
	}

}

func checkUsecases() {
	usecases := strings.Split(cfg.Usecase, ",")
	for _, v := range usecases {
		makeFileWrite(filepath.Join(cfg.Path, "services", "usecases_"+v+".go"),
			fmt.Sprintf("package services\ntype %s struct { \n}", v))
	}
}

func runCommand(command, folder string, arg ...string) {
	cmd := exec.Command(command)
	if folder != "" {
		cmd.Dir = folder
	}
	_, err := cmd.Output()
	if err != nil {
		panic(err)
	}
}

func makeFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, 0777); err != nil {
			panic(err)
		}
	}
}

func makeFile(file string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		if _, err := os.Create(file); err != nil {
			panic(err)
		}
	}
}

func makeFileWrite(file, content string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		f, err := os.Create(file)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		_, err = f.WriteString(content)
		if err != nil {
			panic(err)
		}
	}
}

func writeFile(file, content string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		f, err := os.Open(file)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		_, err = f.WriteString(content)
		if err != nil {
			panic(err)
		}
	}
}

func readFile(file string) string {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		f, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}

		return string(f)
	}
	return ""
}

func copyFolder(old string, new string) {
	// runCommand(fmt.Sprintf("cp -r %s %s", old, new))
}
