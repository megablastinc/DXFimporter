package main

import "fmt"

/* WOL Functions */
func generateWOL(plates map[int]map[string]string, output string) {
	var commands []string

	//commands to do before the recurring process
	commands = append(commands,
		"SET,DESKTOP,OMAXIMPORT")
	commands = append(commands,
		"SAVE,SHEET,ON")
	commands = append(commands,
		"SAVE,REM,ON")

	//commands to do for each plates
	for _, plate := range plates {
		commands = append(commands,
			"SHEET, ADD, REM,"+
				plate["name"]+
				","+plate["group"]+
				","+plate["thickness"]+
				","+"1"+
				","+"DXF"+
				","+plate["filepath"])
	}
	//commands to do after the reccuring process
	commands = append(commands,
		"CLEARWS")
	//fmt.Printf("%#v\n", commands)
	getCommandsTable(commands)
	writeFile(commands, output)
	fmt.Println(output)
	return
}
