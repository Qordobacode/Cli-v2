package cmd

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"github.com/qordobacode/cli-v2/pkg/general/date"
	"github.com/qordobacode/cli-v2/pkg/general/log"
	"github.com/qordobacode/cli-v2/pkg/types"
	"github.com/spf13/cobra"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	enabled   = "ENABLED"
	disabled  = "DISABLED"
	lineLimit = 50
)

// lsCmd represents the ls command
var (
	lsCmd = &cobra.Command{
		Use:   "ls",
		Short: "Ls files (show 50 only)",
		Run:   printLs,
	}
	lsHeaders = []string{"ID", "NAME", "version", "tag", "#SEGMENTS", "UPDATED_ON", "STATUS"}
)

func init() {
	rootCmd.AddCommand(lsCmd)
}

func printLs(cmd *cobra.Command, args []string) {
	workspace, err := WorkspaceService.LoadWorkspace()
	if err != nil {
		return
	}
	data := make([]*responseRow, 0)
	for _, targetPersona := range workspace.Workspace.TargetPersonas {
		result := handlePersonResult(&targetPersona)
		data = append(data, result...)
		if len(data) > lineLimit {
			data = data[:lineLimit]
			break
		}
	}
	// add sorting for output
	sort.Slice(data, func(i, j int) bool {
		return data[i].Name < data[j].Name
	})

	printFile2Stdin(data)
}

func printFile2Stdin(response []*responseRow) {
	if !IsJSON {
		data := formatResponse2Array(response)
		renderTable2Stdin(lsHeaders, data)
	} else {
		bytes, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			log.Errorf("error occurred on marshalling with JSON: %v", err)
			return
		}
		log.Infof("%v", string(bytes))
	}
}

func handlePersonResult(persona *types.Person) []*responseRow {
	files, e := FileService.WorkspaceFilesWithLimit(persona.ID, false, lineLimit)
	data := make([]*responseRow, 0)
	if e != nil {
		return data
	}
	audiences := Config.Audiences()
	for _, file := range files.Files {
		if _, ok := audiences[persona.Code]; len(audiences) > 0 && !ok {
			continue
		}
		data = append(data, buildDataRowFromFile(&file))
	}
	return data
}

func buildDataRowFromFile(file *types.File) *responseRow {
	tags := make([]string, 0, len(file.Tags))
	for _, tag := range file.Tags {
		tags = append(tags, tag.Name)
	}
	// strconv.Itoa
	row := responseRow{
		ID:          file.FileID,
		Name:        file.Filename,
		Version:     file.Version,
		Tag:         tags,
		UpdatedOn:   date.GetDateFromTimestamp(file.Update),
		SegmentNums: file.Counts.SegmentCount,
		Status:      disabled,
	}
	if file.Enabled {
		row.Status = enabled
	}
	return &row
}

func renderTable2Stdin(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.AppendBulk(data)
	table.Render() // Send output
}

func formatResponse2Array(rows []*responseRow) [][]string {
	data := make([][]string, 0, len(rows))
	for _, responseRow := range rows {
		row := make([]string, len(lsHeaders))
		row[0] = strconv.Itoa(responseRow.ID)
		row[1] = responseRow.Name
		row[2] = responseRow.Version
		row[3] = strings.Join(responseRow.Tag, ", ")
		row[4] = strconv.Itoa(responseRow.SegmentNums)
		row[5] = responseRow.UpdatedOn
		row[6] = responseRow.Status
		data = append(data, row)
	}
	return data
}

// responseRow struct
type responseRow struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Tag         []string `json:"tag"`
	SegmentNums int      `json:"#segments"`
	UpdatedOn   string   `json:"updated_on"`
	Status      string   `json:"status"`
}
