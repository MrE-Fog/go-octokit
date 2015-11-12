package octokit

import (
	"github.com/jingweno/go-sawyer/hypermedia"
)

// RepoLabelsURL is a URL template for accessing labels for a repository.
//
// https://developer.github.com/v3/issues/labels/
var RepoLabelsURL = Hyperlink("repos/{owner}/{repo}/labels")

// Labels creates a LabelsService
func (c *Client) Labels() (labels *LabelsService) {
	labels = &LabelsService{client: c}
	return
}

// LabelsService is a service providing access to labels for a repository
type LabelsService struct {
	client *Client
}

// All gets a list of all labels for a repository
//
// https://developer.github.com/v3/issues/labels/#list-all-labels-for-this-repository
func (l *LabelsService) All(uri *Hyperlink, uriParams M) (labels []Label, result *Result) {
	url, err := ExpandWithDefault(uri, &RepoLabelsURL, uriParams)
	if err != nil {
		return nil, &Result{Err: err}
	}

	result = l.client.get(url, &labels)
	return
}

// Label represents a label for a GitHub repository
type Label struct {
	*hypermedia.HALResource

	URL   string `json:"url,omitempty"`
	Name  string `json:"name,omitempty"`
  Color string `json:"color,omitempty"`
}
