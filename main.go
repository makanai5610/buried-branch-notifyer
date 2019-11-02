package main

import (
	"branch-purge-list-creator/command"
	"branch-purge-list-creator/model"
	"fmt"
	"os"
	"time"
)

func main() {
	branches, err := command.ExecGitBranch()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, branch := range branches {
		fmt.Println(branch)
	}

	logs, err := command.ExecGitLog(branches)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, log := range logs {
		fmt.Println(log)
	}

	var branchInfos []model.BranchInfo
	for _, log := range logs {
		branchInfos = append(branchInfos, model.New(log))
	}

	for _, branchInfo := range branchInfos {
		fmt.Printf("%+v\n", branchInfo)
	}

	filteredBranchInfo := filterBranchInfo(branchInfos)
	_ = filteredBranchInfo
}

func filterBranchInfo(branchInfos []model.BranchInfo) []model.BranchInfo {
	now := time.Now()
	var filteredBranchInfos []model.BranchInfo
	for _, branchInfo := range branchInfos {
		days := int(now.Sub(branchInfo.LastCommitDate.Time).Hours()) / 24
		if days >= 14 {
			filteredBranchInfos = append(filteredBranchInfos, branchInfo)
		}
	}
	return filteredBranchInfos
}
