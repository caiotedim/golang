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
	content []byte
}


func SetFile(queryString map[string]string, file []byte) ( bool, int ) {
	var object Attr
	object.abs.path = queryString["path"]
	object.abs.file = queryString["file"]
	object.data.content = file
	//log.Printf("%s", object)
	var check bool
	if object.data.content, check = Crypt(object.data.content, "encrypt"); !check {
		return false, 500
	}
	if dao.SaveLocalFile(object.abs.path, object.abs.file, object.data.content) {
		return true, 201
	}
	return false, 500
}

func GetFile(queryString map[string]string) ( []byte, int ) {
	var object Attr
	var respCode int
	object.abs.path = queryString["path"]
	object.abs.file = queryString["file"]
	object.data.content, respCode = dao.ReadLocalFile(object.abs.path, object.abs.file, object.data.content)
	
	if respCode == 200 {
		var check bool
		if object.data.content, check = Crypt(object.data.content, "decrypt"); !check {
			return []byte("Error on decrypt file"), 500
		}
	} else {
		return object.data.content, respCode
	}
	
	return object.data.content, respCode
}

func RemoveFile(queryString map[string]string) ( []byte, int ) {
	var object Attr
	var respCode int
	object.abs.path = queryString["path"]
	object.abs.file = queryString["file"]
	object.data.content, respCode = dao.RemoveLocalFile(object.abs.path, object.abs.file, object.data.content)
	
	return object.data.content, respCode
}