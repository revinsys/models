package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ReadFile (file *os.File) string {
	// получить размер файла
	stat, err := file.Stat()
	if err != nil {
		return ""
	}
	// чтение файла
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return ""
	}

	str := string(bs)
	return str
}

func GetFile(path string) (*os.File, bool) {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		file, err := os.OpenFile(path, os.O_RDWR, 0755)
		if err != nil {
			return nil, false
		}
		return file, false
	}

	file, err := os.Create(path)
	if err != nil {
		return nil, false
	}

	return file, true
}

func CreateModel (modelName string) {
	file, ok := GetFile("./internal/app/models/" + modelName + ".go")
	if !ok {
		fmt.Println("Модель уже существует")
		return
	}
	defer file.Close()

	fileRepository, ok := GetFile("./internal/app/store/" + modelName + "repository.go")
	if !ok {
		fmt.Println("Репозиторий уже существует")
		return
	}
	defer fileRepository.Close()

	fileStore, _ := GetFile("./internal/app/store/store.go")
	if fileStore == nil {
		return
	}
	defer fileStore.Close()

	modNameFile, err := os.Open("./go.mod")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer modNameFile.Close()

	reader := bufio.NewReader(modNameFile)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	modName := strings.Replace(line, "module ", "", 1)
	modName = strings.Replace(modName, "\n", "", 1)

	store := ReadFile(fileStore)

	replaced := make(map[string]string)
	replaced["$MODEL_NAME"] = modelName
	replaced["$UPPER_NAME"] = strings.Title(modelName)
	replaced["$GIT_PATH"] = modName

	model := GenerateFromTemplate(ModelTemplate, replaced)
	file.WriteString(model)

	repository := GenerateFromTemplate(RepositoryTemplate, replaced)
	fileRepository.WriteString(repository)

	storeItem := GenerateFromTemplate(StoreTemplate, replaced)
	re := regexp.MustCompile(`(type Store struct {)([^}]*)(})`)

	storeLine := GenerateFromTemplate(`$MODEL_NAMERepository *$UPPER_NAMERepository`, replaced)
	store = re.ReplaceAllString(store, `$1$2  `+ storeLine + "\n$3")

	_, err = fileStore.WriteAt([]byte(store + storeItem), 0)
	if err != nil {
		fmt.Println(err)
	}
}
