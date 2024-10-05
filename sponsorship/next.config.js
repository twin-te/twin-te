/** @type {import('next').NextConfig} */
const nextConfig = {
	basePath: '/sponsorship',
	reactStrictMode: true,
	images: {
		domains: ["www.datocms-assets.com"],
	},
	output: "standalone"
};

module.exports = nextConfig;
