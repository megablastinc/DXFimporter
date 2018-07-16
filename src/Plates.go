package main

import (
	"path/filepath"
	"strings"
)

/* Plates Management */
func getPlatesDetails(files []string) (plates map[int]map[string]string) {

	//Initialize the variables
	plates = make(map[int]map[string]string)

	for id, file := range files {
		//Initialize the variables
		plates[id] = make(map[string]string)
		path := strings.ToUpper(file)
		filename := strings.ToUpper(filepath.Base(path))
		extension := strings.ToUpper(filepath.Ext(filename))

		//Decode the name of the file to get the informations about the plate
		decodedInformations := getDecodedDataFromFileName(filename, "-", extension)

		//add the data to the holder
		for key, value := range decodedInformations {
			plates[id][key] = value
			//fmt.Printf("%#v\n", plates[id])
		}
		plates[id]["filename"] = filename
		plates[id]["filepath"] = path
		//fmt.Printf("%#v\n", plates[id])
	}

	return plates
}

func getDecodedDataFromFileName(filename string, seperator string, extension string) (data map[string]string) {
	//define variables
	var input string
	//init the variables
	data = make(map[string]string)

	//remove the file extension if hasExtension is set to yes
	if extension == "" {
		input = filename
	} else {
		input = strings.TrimSuffix(filename, extension)
	}

	//slice the file name based on the seperator provided
	slices := strings.Split(input, seperator)
	//fmt.Printf("%#v\n", slices)
	//plug the slices into the variable
	//print(len(slices))

	data["name"] = input
	data["no"] = slices[5]
	data["group"] = slices[0]
	data["material"] = slices[1]
	data["finish"] = slices[2]
	data["thickness"] = slices[3]
	data["size"] = slices[4]

	if len(slices) >= 7 {
		data["section"] = slices[6]
	}

	return data
}
