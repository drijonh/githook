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

func Push(w http.ResponseWriter, r *http.Request, url string) {
	var body github.PushEvent
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	if len(body.Commits) == 0 {
		return
	}

	branch := strings.Split(*body.Ref, "/heads/")[1]
	if strings.HasPrefix(branch, "gh-readonly-queue/") || strings.HasPrefix(branch, "pr-") {
		return
	}

	var desc strings.Builder
	for _, c := range body.Commits {
		commit := *c
		fmt.Fprintf(&desc,
			"[`%s`](%s) %s\n",
			(*commit.ID)[:7],
			*commit.URL,
			utils.Truncate(
				strings.Split(*commit.Message, "\n")[0],
				62,
			),
		)
	}

	discord.SendWebhook(
		url,
		discord.WebhookPayload{
			Username:  *body.HeadCommit.Author.Login,
			AvatarURL: fmt.Sprintf("https://github.com/%s.png", *body.HeadCommit.Author.Login),
			Embeds: []discord.Embed{
				{
					Title: fmt.Sprintf(
						"%s%s: %d commit%s",
						*body.Repo.FullName,
						utils.Ternary(
							branch == "" || branch == "master" || branch == "main",
							"",
							"@"+branch,
						),
						len(body.Commits),
						utils.Ternary(len(body.Commits) > 1, "s", ""),
					),
					URL:         *body.Compare,
					Description: utils.Truncate(desc.String(), 4000),
					Color:       utils.GetColors().Default,
				},
			},
		},
	)
}
