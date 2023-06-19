package main

import (
	"context"
	"log"

	"interview/app/boot"
	"interview/app/common"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env, err := common.GetEnv()
	if err != nil {
		log.Fatal(err)
	}

	if err := boot.Init(ctx, *env); err != nil {
		log.Fatal(err)
	}
}
