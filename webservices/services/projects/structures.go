package main

import "core"

type ProjectRequest struct{
	Project core.Project      `json:"project,omitempty"`
	Session core.Session `json:"session,omitempty"`
}


type ProjectEditRequest struct{
	Project core.ProjectEdit      `json:"project,omitempty"`
	Session core.Session `json:"session,omitempty"`
}

