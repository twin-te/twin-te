import type { NextSeoProps } from "next-seo";

const SEO: NextSeoProps = {
	defaultTitle: "Twin:te スポンサー",
	titleTemplate: "%s | Twin:te スポンサー",
	description: "Twin:te 協賛ページ",
	openGraph: {
		type: "website",
		title: "Twin:te スポンサー",
		description: "Twin:te 協賛ページ",
		url: "https://sponsorship.twinte.net/",
		site_name: "Twin:te Sponsorship",
		images: [
			{
				url: "https://sponsorship.twinte.net/images/sponsor-ogp.jpg",
				alt: "Twin:te 協賛ページ",
			},
		],
	},
	twitter: {
		handle: "@te_twin",
		site: "@te_twin",
		cardType: "summary_large_image",
	},
	additionalLinkTags: [
		{
			rel: "icon",
			href: "/favicon.ico",
		},
	],
};

export default SEO;
