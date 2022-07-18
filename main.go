package main

import (
	"GauGau/config"
	"GauGau/params"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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

	os.Chdir(cfg.Path)
	if runtime.GOOS == "windows" {
		exec.Command("cmd.exe", "/c", "go mod init", cfg.ProjectName).Output()
		exec.Command("cmd.exe", "/c", "go mod tidy").Output()
	} else {
		exec.Command("go mod init", cfg.ProjectName).Output()
		exec.Command("go mod tidy").Output()
	}

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

	FRAMEWORK_API_LINK := `"github.com/gin-gonic/gin"`
	PACKAGE_FRAMEWORK_API := "gin"
	PACKAGE_FRAMEWORK_ENGINE := "Engine"
	FRAMEWORK_API_NEW := "g := gin.Default()"
	FRAMEWORK_API_RUN := "g.Run()"
	switch cfg.Type {
	case string(params.APIGIN):
		PACKAGE_FRAMEWORK_API = "gin"
		PACKAGE_FRAMEWORK_ENGINE = "Engine"
		FRAMEWORK_API_LINK = `"github.com/gin-gonic/gin"`
	case string(params.APIEcho):
		FRAMEWORK_API_LINK = `"github.com/labstack/echo/v4"`
		PACKAGE_FRAMEWORK_API = "echo"
		PACKAGE_FRAMEWORK_ENGINE = "Echo"
		FRAMEWORK_API_NEW = "g := echo.New()"
		FRAMEWORK_API_RUN = "g.Start()"
	case string(params.APIMux):
	}

	pathRouters := filepath.Join(cfg.Path, "api", "routers", "routers.go")
	fmt.Println("============", pathRouters)
	code := readFile(pathRouters)

	code = strings.Replace(code, `"{{FRAMEWORK_API_NEW}}"`, FRAMEWORK_API_NEW, -1)
	code = strings.Replace(code, `"{{FRAMEWORK_API_RUN}}"`, FRAMEWORK_API_RUN, -1)
	code = strings.Replace(code, `{{FRAMEWORK_API_LINK}}`, FRAMEWORK_API_LINK, -1)
	code = strings.Replace(code, `{{PROJECT_NAME}}`, cfg.ProjectName, -1)

	writeFile(pathRouters, code)

	pathControllers := filepath.Join(cfg.Path, "api", "controllers", "index.go")
	code = readFile(pathControllers)

	code = strings.Replace(code, `{{PACKAGE_FRAMEWORK_API}}`, PACKAGE_FRAMEWORK_API, -1)
	code = strings.Replace(code, `{{PACKAGE_FRAMEWORK_ENGINE}}`, PACKAGE_FRAMEWORK_ENGINE, -1)
	code = strings.Replace(code, `{{FRAMEWORK_API_LINK}}`, FRAMEWORK_API_LINK, -1)
	code = strings.Replace(code, `{{PROJECT_NAME}}`, cfg.ProjectName, -1)

	writeFile(pathControllers, code)

	pathAPIMain := filepath.Join(cfg.Path, "cmd", "api", "main.go")
	code = readFile(pathAPIMain)
	code = strings.Replace(code, `{{PROJECT_NAME}}`, cfg.ProjectName, -1)

	writeFile(pathAPIMain, code)

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
	code := readFile(pathConfig)

	code = strings.Replace(code, "{{CONFIG_TYPE}}", cfgs[1], -1)
	code = strings.Replace(code, "{{CONFIG_NAME}}", cfgs[0], -1)

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

func writeFile(file, content string) {

	ioutil.WriteFile(file, []byte(content), 0644)

}

func readFile(file string) string {

	f, _ := ioutil.ReadFile(file)

	return string(f)

}
