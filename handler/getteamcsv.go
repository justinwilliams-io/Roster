package handler

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"rosterize/model"
	"rosterize/util"
	"rosterize/view/roster"
	"strconv"
	"time"

	"github.com/dnlo/struct2csv"
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

	id := uuid.New()

    fileName := createfile(id, team)

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

func createfile(id uuid.UUID, team model.Team) string {
    fileName := team.Name + "-" + id.String() + ".csv"
	file, err := os.Create("files/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	encoder := struct2csv.New()
	rows, err := encoder.Marshal(team.Roster)
	if err != nil {
		log.Fatal(err)
	}

	writer.WriteAll(rows)

    return fileName
}
