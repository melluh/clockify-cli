package project

import (
	"io"
	"os"

	"github.com/lucassabreu/clockify-cli/api/dto"
	"github.com/lucassabreu/clockify-cli/pkg/output/util"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/term"
)

// ProjectPrint will print more details
func ProjectPrint(ps []dto.Project, w io.Writer) error {
	tw := tablewriter.NewWriter(w)
	tw.SetHeader([]string{"Name", "Client"})

	if width, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil {
		tw.SetColWidth(width / 3)
	}

	colors := make([]tablewriter.Colors, 3)
	for i := 0; i < len(ps); i++ {
		w := ps[i]
		client := "(none)"
		if w.ClientID != "" {
			client = w.ClientName
		}
		colors[1] = []int{}
		if w.Color != "" {
			colors[0] = util.ColorToTermColor(w.Color)
		}

		tw.Rich([]string{
			w.Name,
			client,
		}, colors)
	}

	tw.Render()

	return nil
}
