package dao

import (
	"os"
	"log"
	"io/ioutil"
)

func check(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func isFile(path, file string) bool {
	absPath := path + "/" + file
	fileInfo, err := os.Stat(absPath); if err != nil {
		return false
	}
	if fileInfo.IsDir() {
		return false
	} else {
		return true
	}
}

func ReadLocalFile(path, file string, data []byte) ( []byte, int)  {
	absPath := path + "/"  + file
	if isFile(path, file) {
		log.Printf("Reading file:[%s]", absPath)
		data, err := ioutil.ReadFile(absPath)
		if check(err) {
			log.Printf("Error to read file:[%s] \n %s %s", absPath, err, string(data))
			return []byte("Error to read file!"), 500
		}
		return data, 200
	} else {
		log.Printf("File not found: [%s]", absPath)
		return []byte("Not Found"), 404
	}
	//log.Printf("Content file:[%s]", string(data))
}

func RemoveLocalFile(path, file string, data []byte) ( []byte, int)  {
	absPath := path + "/"  + file
	log.Printf("Removing file:[%s]", absPath)
	err := os.Remove(absPath)
	if check(err) {
		log.Printf("Error to remove file:[%s] \n %s", absPath, err)
		return []byte("Error to remove file!"), 500
	}
	data = []byte("File has been removed!")
	log.Printf("File has been removed!:[%s]", absPath)
	return data, 410
}

func SaveLocalFile(path, file string, data []byte) bool {
	
	if CreateLocalDir(path) {
		// create file
		absPath := path + "/"  + file
		if CreateLocalFile(absPath, data) {
			log.Printf("File Created path:[%s] file:[%s]", path, file)
			return true
		}
	}
	return false
}

func CreateLocalDir(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("Creating directory:[%s]", path)
		if err := os.MkdirAll(path, 0755); err != nil {
			log.Printf("Error to create path:[%s] \n %s", path, err)
			return false
		}	
	}
	return true
}

func CreateLocalFile(file string, data []byte) bool {
	if err := ioutil.WriteFile(file, data, 0644); err != nil {
		log.Printf("Error to create file:[%s] \n %s", file, err)
		return false
	}

	return true
}

