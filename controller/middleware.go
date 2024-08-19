package controller

import (
	"coding-interview-agustus-1/logic"
)

type (
	middleware struct {
		logic *logic.Logic
	}
	header struct {
		token        string
		xLinkService string
		action       logic.ActionType
	}
)

func (m *middleware) Auth(h *header) error {
	user, err := m.logic.Authentication(h.token)
	if err != nil {
		return err
	}

	return m.logic.Authorize(user, h.xLinkService, h.action)
}
