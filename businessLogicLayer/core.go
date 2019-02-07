package businessLogicLayer

import (
	"bytes"
	"podemos-aprender-daily/dataAccessLayer"
	"strconv"
	"strings"
)

func ProcessMsg(msgText string) string {

	// get project and hours from message text
	msgContent := strings.Split(msgText, ",")

	formatErr := "Para que el registro de horas pueda realizarse correctamente debe ingresarlo de la siguiente forma, ej: @time, banco de tiempo, 15"

	if len(msgContent) != 3 {
		return formatErr
	}

	// extract project name from array
	project := removeSpaces(msgContent[1])

	// remove spaces and convert to integer
	hours, err := strconv.Atoi(removeSpaces(msgContent[2]))
	if err != nil {
		return formatErr
	}

	dataAccessLayer.AddHours(project, hours)
	hoursInvested := dataAccessLayer.GetTimeInvested(project)

	var response bytes.Buffer // efficient way to concatenate strings
	response.WriteString("Actualmente se llevan ")
	response.WriteString(strconv.Itoa(hoursInvested))
	response.WriteString(" horas trabajadas en el proyecto ")
	response.WriteString(project)
	return response.String()
}

func removeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
