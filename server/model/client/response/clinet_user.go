package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/client"

type LoginResponse struct {
	User      client.ClientUser `json:"user"`
	Token     string            `json:"token"`
	ExpiresAt int64             `json:"expiresAt"`
}
