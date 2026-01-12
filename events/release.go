package events

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-github/v61/github"
	"github.com/shi-gg/githook/discord"
	"github.com/shi-gg/githook/utils"
)

func Release(w http.ResponseWriter, r *http.Request, url string) {
	var body github.ReleaseEvent
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	if *body.Action != "released" {
		return
	}

	text := strings.TrimSpace(*body.Release.Body)

	discord.SendWebhook(
		url,
		discord.WebhookPayload{
			Username:  *body.Sender.Login,
			AvatarURL: *body.Sender.AvatarURL,
			Embeds: []discord.Embed{
				{
					Title:       fmt.Sprintf("%s: %s Released", *body.Repo.FullName, *body.Release.TagName),
					Description: utils.Truncate(text, 4000),
					URL:         *body.Release.HTMLURL,
					Color:       utils.GetColors().Default,
				},
			},
		},
	)
}
