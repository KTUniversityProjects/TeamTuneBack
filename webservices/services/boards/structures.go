package main

import (
	"core"
)

type BoardRequest struct {
	Board   core.Board   `json:"board,omitempty"`
	Session core.Session `json:"session,omitempty"`
}

type BoardListRequest struct {
	Project core.Project `json:"project"`
	Session core.Session `json:"session"`
}
