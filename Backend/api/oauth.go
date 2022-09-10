package api

import (
	"context"
	"crud-with-auth/db"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	// discord "discord.com/realTristan/disgoauth"
)

type OAuthAPI struct {
	Api
	Config *oauth2.Config
}

const (
	ScopeIdentify                   = "identify"
	ScopeBot                        = "bot"
	ScopeEmail                      = "email"
	ScopeGuilds                     = "guilds"
	ScopeGuildsJoin                 = "guilds.join"
	ScopeConnections                = "connections"
	ScopeGroupDMJoin                = "gdm.join"
	ScopeMessagesRead               = "messages.read"
	ScopeRPC                        = "rpc"                    // Whitelist only
	ScopeRPCAPI                     = "rpc.api"                // Whitelist only
	ScopeRPCNotificationsRead       = "rpc.notifications.read" // Whitelist only
	ScopeWebhookIncoming            = "webhook.Incoming"
	ScopeApplicationsBuildsUpload   = "applications.builds.upload" // Whitelist only
	ScopeApplicationsBuildsRead     = "applications.builds.read"
	ScopeApplicationsStoreUpdate    = "applications.store.update"
	ScopeApplicationsEntitlements   = "applications.entitlements"
	ScopeRelationshipsRead          = "relationships.read" // Whitelist only
	ScopeActivitiesRead             = "activities.read"    // Whitelist only
	ScopeActivitiesWrite            = "activities.write"   // Whitelist only
	ScopeApplicationsCommands       = "applications.commands"
	ScopeApplicationsCommandsUpdate = "applications.commands.update"
)

// Endpoint is Discord's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:   "https://discord.com/api/oauth2/authorize",
	TokenURL:  "https://discord.com/api/oauth2/token",
	AuthStyle: oauth2.AuthStyleInParams,
}

func (api *Api) OAuth() {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("DISCORD_OAUTH_CLIEND_ID"),
		ClientSecret: os.Getenv("DISCORD_OAUTH_CLIENT_SECRET"),
		// Scopes:       []string{"user:email"},
		Scopes: []string{ScopeIdentify},

		Endpoint:    Endpoint,
		RedirectURL: "http://localhost:3000/login", //fontend url
	}

	oauthApi := OAuthAPI{
		Api:    *api,
		Config: conf,
	}

	oauthApi.BuildRoutes()
}

func (oauth *OAuthAPI) BuildRoutes() {
	group := oauth.Api.r.Group("/oauth/discord")
	group.GET("/login", oauth.DiscordAuthorizeHandler)
	group.GET("/callback", oauth.DiscordCallbackHandler)
}

func (oauth *OAuthAPI) DiscordAuthorizeHandler(c *gin.Context) {
	url := oauth.Config.AuthCodeURL("yooo")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

type DiscordUserResponse struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	PublicFlags   uint32 `json:"public_flags"`
	Flags         uint32 `json:"flags"`
	Locale        string `json:"locale"`
	MfaEnabled    bool   `json:"mfa_enabled"`
	PremiumType   uint8  `json:"premium_type"`
}

func (oauth *OAuthAPI) DiscordCallbackHandler(c *gin.Context) {
	code, ok := c.GetQuery("code")

	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "code is empty",
		})
		return
	}

	ctx := context.Background()
	token, err := oauth.Config.Exchange(ctx, code)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://discord.com/api/users/@me", nil)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	res, err := client.Do(req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer res.Body.Close()

	var discordRespone = DiscordUserResponse{}
	err = json.NewDecoder(res.Body).Decode(&discordRespone)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user db.User
	rows := oauth.db.Storage.Where("provider_id = ? AND provider = 'discord'", discordRespone.Id).First(&user).RowsAffected

	if rows == 0 {
		user.Provider = "discord"
		user.ProviderID = discordRespone.Id
		user.Avatar = discordRespone.Avatar
		user.Username = discordRespone.Username
		if err = oauth.db.Storage.Create(&user).Error; err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	// c.Redirect(http.StatusTemporaryRedirect, "")

	if token, err := user.Token(); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successfully",
			"user":    discordRespone,
			"token":   token,
		})
	}

}
