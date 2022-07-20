// Created by EldersJavas(EldersJavas&gmail.com)

package ebiten_rpc

import (
	"github.com/hugolgst/rich-go/client"
	"time"
)

type GameRPC struct {
	AppID          string
	GameName       string
	State          string
	LargeImageID   string
	LargeImageText string
	SmallImageID   string
	SmallImageText string
	Party          *client.Party
	JoinID         string
	StartTime      time.Time
	Button         []*client.Button
	Secret         *client.Secrets
}

func (g *GameRPC) Start() error {
	g.StartTime = time.Now()
	err := client.Login(g.AppID)
	if err != nil {
		return err
	}
	err = client.SetActivity(client.Activity{
		Details:    "Playing " + g.GameName,
		State:      g.State,
		LargeImage: g.LargeImageID,
		LargeText:  g.LargeImageText,
		SmallImage: g.SmallImageID,
		SmallText:  g.SmallImageText,
		Party:      g.Party,
		Timestamps: &client.Timestamps{Start: &g.StartTime},
		Secrets:    g.Secret,
		Buttons:    g.Button,
	})
	if err != nil {
		return err
	}
	return nil
}

func (g *GameRPC) AddButton(Label, Url string) error {
	b := &client.Button{Label: Label, Url: Url}
	g.Button = append(g.Button, b)
	err := g.Update()
	if err != nil {
		return err
	}
	return nil
}
func (g *GameRPC) AddSecret(Match, Join, Spectate string) error {
	b := &client.Secrets{
		Match:    Match,
		Join:     Join,
		Spectate: Spectate,
	}
	g.Secret = b
	err := g.Update()
	if err != nil {
		return err
	}
	return nil
}

func (g *GameRPC) Update() error {
	err := client.SetActivity(client.Activity{
		Details:    "Playing " + g.GameName,
		State:      g.State,
		LargeImage: g.LargeImageID,
		LargeText:  g.LargeImageText,
		SmallImage: g.SmallImageID,
		SmallText:  g.SmallImageText,
		Party:      g.Party,
		Timestamps: &client.Timestamps{Start: &g.StartTime},
		Secrets:    g.Secret,
		Buttons:    g.Button,
	})

	if err != nil {
		return err
	}
	return nil
}

func NewGameRPC(AppID string) *GameRPC {
	return &GameRPC{AppID: AppID}
}
