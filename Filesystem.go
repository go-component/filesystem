package filesystem

import "os"

func Mkdir(path string,  mode os.FileMode){

	if path == ""{
		return
	}

	err := os.MkdirAll(path, mode)

	if err != nil{
		panic(err)
	}
}

func MultiMkdir(paths []string,  mode os.FileMode){

	if len(paths) == 0{
		return
	}

	for _,path := range paths{

		err := os.MkdirAll(path, mode)

		if err != nil{
			panic(err)
		}
	}


}