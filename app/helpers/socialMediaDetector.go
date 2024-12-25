package helpers

import "strings"

// IsSocialMediaCrawler detects if the user agent is a social media crawler
func IsSocialMediaCrawler(userAgent string) bool {
	crawlers := []string{
		// Social Media Platforms
		"facebookexternalhit", // Facebook
		"Facebot",             // Facebook
		"Twitterbot",          // Twitter
		"LinkedInBot",         // LinkedIn
		"Slackbot",            // Slack
		"Pinterest",           // Pinterest
		"redditbot",           // Reddit
		"discordbot",          // Discord
		// Messaging Apps
		"WhatsApp",    // WhatsApp
		"TelegramBot", // Telegram
		"vkShare",     // VK
		"Viber",       // Viber
		"Line",        // LINE
		"Applebot",    // iMessage (Apple)
		"Snapchat",    // Snapchat
		// Search Engines (Optional)
		"Googlebot",    // Google
		"Bingbot",      // Bing
		"Yahoo! Slurp", // Yahoo
		"DuckDuckBot",  // DuckDuckGo
		"YandexBot",    // Yandex
	}

	userAgent = strings.ToLower(userAgent)
	for _, crawler := range crawlers {
		if strings.Contains(userAgent, strings.ToLower(crawler)) {
			return true
		}
	}
	return false
}
