package main

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

/* Table genrator functions */
func getFilesTable(files []string) {
	//create the table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "filepath"})

	//generate the lines
	for id, file := range files {
		slicr := []string{strconv.Itoa(id), file}
		table.Append(slicr)
		//fmt.Printf("%#v\n", slicr)
	}

	table.Render()
}

func getPlatesDetailsTable(plates map[int]map[string]string) {
	//create the table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID.", "Name", "No.", "Section", "Group", "Material", "Finish", "Thickness", "Size", "File Name", "File Path"})

	//generate the lines
	for id, plate := range plates {
		line := []string{
			strconv.Itoa(id),
			plate["name"],
			plate["no"],
			plate["section"],
			plate["group"],
			plate["material"],
			plate["finish"],
			plate["thickness"],
			plate["size"],
			plate["filename"],
			plate["filepath"],
		}
		table.Append(line)
		//fmt.Printf("%#v\n", line)
	}

	table.Render()
}

func getCommandsTable(commands []string) {
	//create the table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "commands"})

	//generate the lines
	for id, command := range commands {
		slicr := []string{strconv.Itoa(id), command}
		table.Append(slicr)
		//fmt.Printf("%#v\n", slicr)
	}

	table.Render()
}
