package main

import (
	"github.com/xuebing1110/gohomekit-aqara/aqara"
	"github.com/xuebing1110/gohomekit/bridge"
)

func main() {

}

func PlatForm() (bridge.PlatForm, error) {
	return &aqara.Aqara{}, nil
}
