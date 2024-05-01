package handler

import (
	"encoding/json"
	"rosterize/model"
	"rosterize/util"
	"rosterize/view/roster"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type GetTeamCsvHandler struct{}

type DownloadJsonData struct {
	Id       uuid.UUID
	TeamName string
}

func (h GetTeamCsvHandler) GetCsv(c echo.Context) error {
	c.Request().ParseForm()
	form := c.Request().Form

	team := mapTeam(form)

	fileName := team.CreateCsv()

	data := struct {
		DownloadCsv string `json:"downloadCsv"`
	}{
		DownloadCsv: fileName,
	}

	jsonData, _ := json.Marshal(data)

	c.Response().Header().Add("HX-Trigger", string(jsonData))

	go func(filename string) {
		time.Sleep(5 * time.Minute)

		util.DeleteFile(filename)
	}(fileName)

	return render(c, roster.Blank())
}

func mapTeam(formData map[string][]string) model.Team {
	var team model.Team
	var players []model.Player

	team.Name = formData["team_name"][0]

	for i, firstName := range formData["first_name[]"] {
		jerseyNumber, _ := strconv.Atoi(formData["jersey_number[]"][i])

		players = append(players, model.Player{
			FirstName:    string(firstName),
			LastName:     formData["last_name[]"][i],
			JerseySize:   formData["jersey_size[]"][i],
			JerseyNumber: jerseyNumber,
		})
	}

	team.Roster = players

	return team
}
