package handler

import (
	"bangwon/api/request"
	"bangwon/api/response"
	"bangwon/global"
	"github.com/gofiber/fiber/v2"
)

func GetStatus(ctx *fiber.Ctx) error {
	me := global.GetMe()
	if me == nil {
		return nil
	}
	res := response.From(*me)
	return ctx.JSON(res)
}

func PatchStatus(ctx *fiber.Ctx) error {
	req := request.Status{}
	me := global.SetMe(req.To())
	res := response.From(*me)
	return ctx.JSON(res)
}
