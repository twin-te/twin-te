package appenv

var (
	// auth
	SESSION_LIFE_TIME_DAYS int = loadInt("SESSION_LIFE_TIME_DAYS")

	// timetable
	COURSE_CACHE_HOURS int = loadInt("COURSE_CACHE_HOURS")

	// donation
	STRIPE_KEY                  string = loadString("STRIPE_KEY")
	STRIPE_CHECKOUT_SUCCESS_URL string = loadString("STRIPE_CHECKOUT_SUCCESS_URL")
	STRIPE_CHECKOUT_CANCEL_URL  string = loadString("STRIPE_CHECKOUT_CANCEL_URL")

	// db
	DB_URL      string = loadString("DB_URL")
	TEST_DB_URL string = loadString("TEST_DB_URL")

	// handler
	ADDR                 string   = loadString("ADDR")
	AUTH_REDIRECT_URL    string   = loadString("AUTH_REDIRECT_URL")
	CORS_ALLOWED_ORIGINS []string = loadStringSlice("CORS_ALLOWED_ORIGINS")

	OAUTH2_GOOGLE_CLIENT_ID     string = loadString("OAUTH2_GOOGLE_CLIENT_ID")
	OAUTH2_GOOGLE_CLIENT_SECRET string = loadString("OAUTH2_GOOGLE_CLIENT_SECRET")
	OAUTH2_GOOGLE_CALLBACK_URL  string = loadString("OAUTH2_GOOGLE_CALLBACK_URL")

	OAUTH2_APPLE_CLIENT_ID    string = loadString("OAUTH2_APPLE_CLIENT_ID")
	OAUTH2_APPLE_TEAM_ID      string = loadString("OAUTH2_APPLE_TEAM_ID")
	OAUTH2_APPLE_KEY_ID       string = loadString("OAUTH2_APPLE_KEY_ID")
	OAUTH2_APPLE_PRIVATE_KEY  string = loadString("OAUTH2_APPLE_PRIVATE_KEY")
	OAUTH2_APPLE_CALLBACK_URL string = loadString("OAUTH2_APPLE_CALLBACK_URL")

	OAUTH2_TWITTER_CLIENT_ID     string = loadString("OAUTH2_TWITTER_CLIENT_ID")
	OAUTH2_TWITTER_CLIENT_SECRET string = loadString("OAUTH2_TWITTER_CLIENT_SECRET")
	OAUTH2_TWITTER_CALLBACK_URL  string = loadString("OAUTH2_TWITTER_CALLBACK_URL")

	COOKIE_SECURE               bool   = loadBool("COOKIE_SECURE")
	COOKIE_SESSION_NAME         string = loadString("COOKIE_SESSION_NAME")
	COOKIE_OAUTH2_STATE_NAME    string = loadString("COOKIE_OAUTH2_STATE_NAME")
	COOKIE_OAUTH2_STATE_MAX_AGE int    = loadInt("COOKIE_OAUTH2_STATE_MAX_AGE")
)
