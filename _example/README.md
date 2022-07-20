# Example

> From: <https://github.com/hajimehoshi/ebiten/blob/main/examples/paint/main.go>

### Patient:
1. When you edit the rpc, you need to use `RPC.Update()` to update.
2. Do not forget filling in your AppID


### Main code:
```go
//Find AppID: https://discord.com/developers/applications 
var RPC = ebiten_rpc.NewGameRPC("Your AppID")

func (g *Game) UpdatePos() {
    x, y := ebiten.CursorPosition()
    RPC.State = fmt.Sprintf("Mouse Pos: %v,%v", x, y)
    err := RPC.Update()
    if err != nil {
        panic(err)
    }
}

func initrpc() {
	RPC.GameName = "Ebiten _example"
	RPC.LargeImageID = "ebiten"
	RPC.LargeImageText = "Ebitengine"
	err := RPC.AddButton("Official Website", "https://ebiten.org/")
	if err != nil {
		return
	}
	err = RPC.Start()
	if err != nil {
		panic(err)
	}
}
```
