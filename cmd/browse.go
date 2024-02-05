package cmd

import (
	"fmt"

	"github.com/pkg/browser"
)

func Browse(nodeId int) error {
	url := fmt.Sprintf("https://goeverywhere.ernie.org/#nid=%d", nodeId)
	err := browser.OpenURL(url)
	if err != nil {
		return fmt.Errorf("Error opening url: %w", err)
	}
	return nil
}
