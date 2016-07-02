package dao

import "log"

/*
{
	"Mirella_TCC": [{
		"teste_1": {
			"img1": "/tmp/teste.jpeg",
			"img2": "/tmp/teste.jpeg"
		},
		"teste_2": {
			"img1": "/tmp/teste.jpeg",
			"img2": "/tmp/teste.jpeg"
		}
	}]
}
 */

type form struct  {
	name []questions // nome do teste
}

type questions struct {
	qst image
}

type image struct {
	image1 string
	image2 string
}

func Save(imgName string, path1 string, path2 string, correct bool)  {
	form.id = imgName

	if correct {
		form.name = path1
		form.img.pathMod = path2
	} else {
		form.img.pathOr = path2
		form.img.pathOr = path1
	}

	log.Printf("Name: %s", form.id)
	log.Printf("pathOr: %s; pathMod: %s", form.img.pathOr, form.img.pathMod)
}
