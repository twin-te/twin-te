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
	CORS_ALLOWED_ORIGINS []string = loadStringSlice("CORS_ALLOWED_ORIGINS")

	AUTH_DEFAULT_REDIRECT_URL  string   = loadString("AUTH_DEFAULT_REDIRECT_URL")
	AUTH_ALLOWED_REDIRECT_URLS []string = loadStringSlice("AUTH_ALLOWED_REDIRECT_URLS")

	AUTH_GOOGLE_CLIENT_ID     string = loadString("AUTH_GOOGLE_CLIENT_ID")
	AUTH_GOOGLE_CLIENT_SECRET string = loadString("AUTH_GOOGLE_CLIENT_SECRET")
	AUTH_GOOGLE_CALLBACK_URL  string = loadString("AUTH_GOOGLE_CALLBACK_URL")

	AUTH_APPLE_CLIENT_ID    string = loadString("AUTH_APPLE_CLIENT_ID")
	AUTH_APPLE_TEAM_ID      string = loadString("AUTH_APPLE_TEAM_ID")
	AUTH_APPLE_KEY_ID       string = loadString("AUTH_APPLE_KEY_ID")
	AUTH_APPLE_PRIVATE_KEY  string = loadString("AUTH_APPLE_PRIVATE_KEY")
	AUTH_APPLE_CALLBACK_URL string = loadString("AUTH_APPLE_CALLBACK_URL")

	AUTH_TWITTER_CLIENT_ID     string = loadString("AUTH_TWITTER_CLIENT_ID")
	AUTH_TWITTER_CLIENT_SECRET string = loadString("AUTH_TWITTER_CLIENT_SECRET")
	AUTH_TWITTER_CALLBACK_URL  string = loadString("AUTH_TWITTER_CALLBACK_URL")

	COOKIE_SECURE                 bool   = loadBool("COOKIE_SECURE")
	COOKIE_SESSION_NAME           string = loadString("COOKIE_SESSION_NAME")
	COOKIE_AUTH_STATE_NAME        string = loadString("COOKIE_AUTH_STATE_NAME")
	COOKIE_AUTH_REDIRECT_URL_NAME string = loadString("COOKIE_AUTH_REDIRECT_URL_NAME")
)
