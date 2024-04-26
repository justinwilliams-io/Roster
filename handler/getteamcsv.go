package handler

import (
	"encoding/csv"
	"log"
	"os"
	"rosterize/model"
	"rosterize/view/roster"
	"strconv"

	"github.com/dnlo/struct2csv"
	"github.com/labstack/echo/v4"
)

type GetTeamCsvHandler struct{}

func (h GetTeamCsvHandler) GetCsv(c echo.Context) error {
	c.Request().ParseForm()
	form := c.Request().Form

    players := mapPlayers(form)
    
    file, err := os.Create("test.csv")
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    encoder := struct2csv.New()
    rows, _ := encoder.Marshal(players)

    writer.WriteAll(rows)

	return render(c, roster.Blank())
}

func mapPlayers(formData map[string][]string) []model.Player {
    var players []model.Player

    for i, firstName := range formData["first_name[]"] {
        jerseyNumber, _ := strconv.Atoi(formData["jersey_number[]"][i])

        players = append(players, model.Player{
            FirstName: string(firstName),
            LastName: formData["last_name[]"][i],
            JerseySize: formData["jersey_size[]"][i],
            JerseyNumber: jerseyNumber,
        })   
    }

    return players
}
