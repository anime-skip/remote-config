package operations

import (
	"net/http"
	"sort"

	"anime-skip.com/remote-config/src/backend"
	"anime-skip.com/remote-config/src/backend/utils"
)

func GetAppsHandler(repo backend.Repo) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		apps, err := repo.GetApps()
		if err != nil {
			utils.SendErrorJSON(rw, err)
			return
		}
		sort.Strings(apps)
		utils.SendJSON(rw, http.StatusOK, apps)
	}
}
