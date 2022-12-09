package main

import (
	"bangwon/api/handler"
	"bangwon/config"
	"bangwon/global"
	"bangwon/heartbeat"
	_type "bangwon/type"
	"bangwon/utils"
	"context"
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	setLogger()
	config := LoadConf()
	initStatus()
	go runApi(config)
	go runHeartbeat(config)
	go runCoup(config)
	go runDomestic(config)
	for {
		time.Sleep(time.Second)
	}
}

func setLogger() {
	log.SetReportCaller(true)
	log.SetFormatter(&nested.Formatter{
		CallerFirst: true,
	})
}

func LoadConf() config.Config {
	var config config.Config
	err := utils.LoadFile("./config/config.yaml", &config)
	if err != nil {
		log.Panicf("load config err: %+v", err)
	}
	log.Infof("load config success: %+v", config)
	global.MeId = config.Id
	return config
}

func initStatus() {
	me := _type.Status{}
	me.SetRole(_type.Standby)
	global.SetMe(me)
}

func runApi(config config.Config) {
	api := fiber.New()
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})
	api.Get("/status", handler.GetStatus)
	api.Patch("/status", handler.PatchStatus)
	api.Listen(fmt.Sprintf(":%d", config.Port))
}

func runHeartbeat(config config.Config) {
	interval := time.Second * 5
	apiClient := heartbeat.NewApiClient()
	utils.Tick(interval, func(ctx context.Context) {
		for _, bangwon := range config.Bangwons {
			status, err := apiClient.GetStatus(bangwon.Server)
			if err == nil {
				log.Infof("bangwon(%s): %+v", bangwon.Id, status)
			}
			global.UpdateBangwon(bangwon.Id, status)
		}
	})
}

func runCoup(config config.Config) {
	interval := time.Second * 5
	utils.Tick(interval, func(ctx context.Context) {
		candidate := true
		global.Bangwons().Range(func(id, status any) bool {
			if id == global.MeId {
				return true
			}
			role := status.(_type.Status).Role()
			if role == _type.Active {
				candidate = false
			}
			return true
		})

		me := global.GetMe()
		if me == nil {
			time.Sleep(interval)
			return
		}

		if !candidate {
			log.Infof("to be demoted")
			me.SetRole(_type.Standby)
			exitCode, err := utils.Execute(config.SrcPath + "/kill.bat")
			if err == nil && exitCode == 0 {
				log.Infof("demote success")
			} else {
				log.Infof("demote failed. to be promoted")
				me.SetRole(_type.Active)
			}
			global.SetMe(*me)
		}

		if candidate && me.Role() == _type.Standby {
			log.Infof("to be promoted")
			me.SetRole(_type.Active)
			exitCode, err := utils.Execute(config.SrcPath + "/ps.bat")
			if err == nil && exitCode == 0 {
				log.Warnf("process Already running")
			}
			exitCode, err = utils.Execute(config.SrcPath + "/run.bat")
			if err == nil && exitCode == 0 {
				log.Infof("promote success")
			} else {
				log.Infof("promote failed. to be demoted")
				_, _ = utils.Execute(config.SrcPath + "/kill.bat")
				me.SetRole(_type.Standby)
			}
			global.SetMe(*me)
		}
	})
}

func runDomestic(config config.Config) {
	interval := time.Second * 5
	utils.Tick(interval, func(ctx context.Context) {
		me := global.GetMe()
		if me == nil {
			time.Sleep(interval)
			return
		}

		exitCode, err := utils.Execute(config.SrcPath + "/ps.bat")
		if err != nil || exitCode != 0 {
			log.Warnf("process warning. to be demoted")
			_, _ = utils.Execute(config.SrcPath + "/kill.bat")
			me.SetRole(_type.Standby)
		}
		global.SetMe(*me)
	})
}
