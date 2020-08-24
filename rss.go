package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

func ParseRss(url string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	feeds, err := gofeed.NewParser().ParseURLWithContext(url, ctx)
	if err != nil {
		return err
	}
	fmt.Println(feeds)

	return nil
}
