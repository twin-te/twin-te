import type { SubscriptionPlan } from "@/domain";
import { useCase } from "@/usecases";
import type { NextPage } from "next";
import { NextSeo } from "next-seo";
import Link from "next/link";
import { useRouter } from "next/router";
import Script from "next/script";
import { useEffect, useState } from "react";

const Success: NextPage = () => {
	const router = useRouter();
	const { type, amount, plan_id: planId } = router.query;

	const [isLoading, setIsLoading] = useState<boolean>(true);
	const [subscriptionPlans, setSubscriptionPlans] = useState<
		SubscriptionPlan[]
	>([]);

	useEffect(() => {
		useCase
			.getSubscriptionPlans()
			.then((subscriptionPlans) => {
				subscriptionPlans.sort((planA, planB) => planA.amount - planB.amount);
				setSubscriptionPlans(subscriptionPlans);
			})
			.finally(() => setIsLoading(false));
	}, []);

	const getTypeText = () => {
		if (typeof type !== "string") return "";

		if (type === "onetime") return "1回きりの寄付";
		if (type === "subscription") return "サブスクリプション寄付（継続寄付）";
		return "";
	};

	const amountText = () => {
		if (type === "subscription" && planId) {
			const subscriptionPlan = subscriptionPlans.find(
				(plan) => plan.id === planId,
			);
			return subscriptionPlan?.amount ?? "";
		}
		if (typeof amount === "string" && /^[0-9]+(\.[0-9]+)?$/.test(amount)) {
			return amount;
		}
		return "";
	};

	const getTwitterText = () => {
		if (typeof type !== "string") return "Twin:teに寄付しました！";

		if (type === "onetime" && amountText())
			return `Twin:teに${amountText()}円を寄付しました！`;
		if (type === "subscription" && amountText())
			return `Twin:teに月課金として${amountText()}円/月の寄付登録をしました！`;
		return "Twin:teに寄付しました！";
	};

	if (isLoading) {
		return <div>now loading...</div>;
	}

	return (
		<>
			<NextSeo title="寄付完了" />
			<h1 className="title">ありがとうございました！</h1>
			<h2 className="has-text-weight-bold">{getTypeText()}</h2>
			{amountText() ? (
				<>
					<p className="has-text-centered has-text-primary has-text-weight-bold is-size-2">
						¥{amountText()}
					</p>
					<p>以上の金額が寄付されました。</p>
				</>
			) : null}
			<p>
				<Link href="/mypage">マイページ</Link>
				からユーザー情報の編集や寄付の履歴が確認できます。
			</p>
			<p className="mt-4">
				<a
					href="https://twitter.com/share?ref_src=twsrc%5Etfw"
					className="twitter-share-button"
					data-text={getTwitterText()}
					data-size="large"
					data-url="https://sponsorship.twinte.net"
					data-via="te_twin"
					data-hashtags="Twinte"
					data-related="te_twin"
					data-show-count="false"
				>
					Tweet
				</a>
				<Script
					async
					src="https://platform.twitter.com/widgets.js"
					charSet="utf-8"
				/>
			</p>
		</>
	);
};

export default Success;
