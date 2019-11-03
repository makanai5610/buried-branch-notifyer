package slack

import (
	"branch-purge-list-creator/model/git"
	"fmt"
)

type Attachment struct {
	Color  string  `json:"color"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func NewAttachments(branchOwnerMap map[string][]git.BranchInformation) []Attachment {
	var attachments []Attachment
	for key, branchInformations := range branchOwnerMap {
		attachment := Attachment{Color: "danger"}
		var buffer []byte
		for _, branchInformation := range branchInformations {
			message := fmt.Sprintf("%s has stopped from %s\n", branchInformation.BranchName, branchInformation.LastCommitDate.Format("2006-01-02"))
			buffer = append(buffer, message...)
		}
		attachment.Fields = append(attachment.Fields, Field{Title: key, Value: string(buffer), Short: false})
		attachments = append(attachments, attachment)
	}
	return attachments
}