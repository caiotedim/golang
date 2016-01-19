package bo

import (
	"dao"
)

type Attr struct {
	abs QueryString
	data Object
}

type QueryString struct {
	path string
	file string
}

type Object struct {
	data []byte
}


func SetFile(queryString map[string]string, file []byte) ( bool, int ) {
	var object Attr
	object.abs.path = queryString["path"]
	object.abs.file = queryString["file"]
	object.data.data = file
	//log.Printf("%s", object)
	var check bool
	if object.data.data, check = Crypt(object.data.data, "encrypt"); !check {
		return false, 500
	}
	if dao.SaveLocalFile(object.abs.path, object.abs.file, object.data.data) {
		return true, 201
	}
	return false, 500
}

func GetFile(queryString map[string]string) ( []byte, int ) {
	var object Attr
	var respCode int
	object.abs.path = queryString["path"]
	object.abs.file = queryString["file"]
	object.data.data, respCode = dao.ReadLocalFile(object.abs.path, object.abs.file, object.data.data)
	
	var check bool
	if object.data.data, check = Crypt(object.data.data, "decrypt"); !check {
		return object.data.data, 500
	}
	
	return object.data.data, respCode
}

func RemoveFile(queryString map[string]string) ( []byte, int ) {
	var object Attr
	var respCode int
	object.abs.path = queryString["path"]
	object.abs.file = queryString["file"]
	object.data.data, respCode = dao.RemoveLocalFile(object.abs.path, object.abs.file, object.data.data)
	
	return object.data.data, respCode
}