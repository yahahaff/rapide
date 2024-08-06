package initialize

import (
	"github.com/yahahaff/rapide/backend/pkg/cloudflare"
	"github.com/yahahaff/rapide/backend/pkg/config"
)

// SetupR2 initializes Cloudflare R2 client
func SetupR2() {
	Endpoint := config.GetString("ENDPOINT_URL", "https://aa7f603da66f796cbefd1fb7fa990aef.r2.cloudflarestorage.com")
	AccessKey := config.GetString("ACCESS_KEY", "ed68ce2c35d9c992e20e0087cd088c8b")
	SecretKey := config.GetString("SECRET_KEY", "90800662b41e5ca1aaa115205d5cd56bf18f3a11bc918f94795ef2c61726d66d")
	BucketName := config.GetString("BUCKET_NAME", "test-ax")

	var err error
	cloudflare.R2, err = cloudflare.NewR2Client(Endpoint, AccessKey, SecretKey, BucketName)
	if err != nil {
		panic("Failed to initialize R2Client: " + err.Error())
	}
}
