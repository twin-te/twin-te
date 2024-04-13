import { useEffect, useState } from 'react';
import type { NextPage } from 'next';
import Slider from 'react-input-slider';
import styles from '../styles/pages/Register.module.scss';
import { redirectToCheckout } from '../stripe';
import { NextSeo } from 'next-seo';
import Link from 'next/link';
import { useCase } from '../usecases';
import { SubscriptionPlan } from '../domain';
import SweetModal from '@/components/SweetAlert';
import RadioButton from '@/components/RadioButton';

const ACTUAL_RECEIVED_PERCENTAGE = 0.964;
const MONTHLY_COST = 7052; // ref: /public/images/twinte-cost.png

const Register: NextPage = () => {
	const donationPrices = [500, 1000, 1500, 2000, 3000, 5000, 7000, 10000];
	const [donationPriceIndex, setDonationPriceIndex] = useState(0);

	const [isLoading, setIsLoading] = useState<boolean>(true);
	const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
	const [subscriptionPlans, setSubscriptionPlans] = useState<SubscriptionPlan[]>([]);
	const [subscriptionPlanId, setSubscriptionPlanId] = useState<string>('');

	useEffect(() => {
		Promise.all([
			useCase.checkAuthentication().then((isAuthenticated) => {
				setIsAuthenticated(isAuthenticated);
			}),
			useCase.getSubscriptionPlans().then((subscriptionPlans) => {
				if (subscriptionPlans.length === 0) throw new Error('not found subscription plans');
				subscriptionPlans.sort((planA, planB) => planA.amount - planB.amount);
				setSubscriptionPlans(subscriptionPlans);
				setSubscriptionPlanId(subscriptionPlans[0].id);
			})
		]).finally(() => setIsLoading(false));
	}, []);

	const promptToLoginForOneTimeDonation = async () => {
		await SweetModal.fire({
			title: 'ログインをしてください。',
			text: '寄付をするには、右上のログインボタンよりログインをしてください。',
			showCancelButton: false,
			confirmButtonText: 'はい'
		});
	};
	const promptToLoginForSubscription = async () => {
		await SweetModal.fire({
			title: 'ログインをしてください。',
			text: '継続的な寄付をするには、右上のログインボタンよりログインをしてください。',
			showCancelButton: false,
			confirmButtonText: 'はい'
		});
	};

	const radioButtons = () => {
		return subscriptionPlans.map((plan, index) => {
			return (
				<div key={plan.id} className="field has-text-weight-bold">
					<RadioButton
						defaultChecked={index === 0}
						name="priceChoice"
						id={plan.id}
						value={plan.id}
						onChange={(newValue) => setSubscriptionPlanId(newValue)}
					>
						{plan.amount}円/月
					</RadioButton>
				</div>
			);
		});
	};

	if (isLoading) {
		return <div>loading...</div>;
	}

	return (
		<>
			<NextSeo title="寄付・サブスク登録" />
			<h1 className="title pagetitle">寄付・サブスク登録</h1>
			<div className="card">
				<h1 className={`title ${styles.title}`}>1回きりの決済による寄付</h1>
				<p>
					1回の決済による単発の寄付です。 <br />
					スライダーを動かして金額を設定し、「寄付する」ボタンを押すと、決済ページへ移動します。
				</p>
				<p className={`has-text-centered has-text-primary ${styles.price}`}>¥{donationPrices[donationPriceIndex]}</p>
				<div className={styles.slider}>
					<Slider
						styles={{
							active: {
								backgroundColor: '#97C3C3'
							},
							track: {
								width: '100%'
							},
							thumb: {
								height: '1.5rem',
								width: '1.5rem',
								border: 'solid 0.4rem #00c0c0'
							}
						}}
						axis="x"
						xmin={0}
						xmax={7}
						x={donationPriceIndex}
						onChange={(price) => setDonationPriceIndex(price.x)}
					/>
				</div>
				<p className="has-text-primary">
					ご協力いただく金額で、Twin:teを
					<span style={{ fontWeight: 'bold' }}>
						{Math.round(
							(Math.floor(donationPrices[donationPriceIndex] * ACTUAL_RECEIVED_PERCENTAGE) / MONTHLY_COST) * 100
						) / 100}
						ヶ月
					</span>
					運営することができます。
				</p>
				<p className={styles.priceNotification}>
					※手数料を差し引くとTwin:teには{donationPrices[donationPriceIndex] * ACTUAL_RECEIVED_PERCENTAGE}
					円寄付されます。
				</p>
				<button
					className={`button is-fullwidth is-primary ${styles.buttons}`}
					onClick={() => {
						isAuthenticated
							? useCase.makeOneTimeDonation(donationPrices[donationPriceIndex]).then(redirectToCheckout)
							: promptToLoginForOneTimeDonation();
					}}
				>
					寄付する
				</button>
			</div>

			<div className="card">
				<h1 className={`title ${styles.title}`}>サブスクリプション（毎月のお支払い）の登録</h1>
				<p>
					毎月決済が行われるサブスクリプションです。
					<br />
					月ごとにお支払いいただく金額を下記から選択し、「登録する」ボタンを押すと、決済ページへ移動します。
				</p>
				<p className="has-text-weight-bold">
					このサブスクリプションは
					<Link href="/mypage" passHref>
						マイページ
					</Link>
					よりいつでもご解約いただけます。
				</p>
				<div className={`field ${styles.radioButtonField}`}>{radioButtons()}</div>
				<button
					className={`button is-fullwidth is-primary ${styles.buttons}`}
					onClick={() => {
						isAuthenticated
							? useCase.registerSubscription(subscriptionPlanId).then(redirectToCheckout)
							: promptToLoginForSubscription();
					}}
				>
					登録する
				</button>
			</div>
		</>
	);
};

export default Register;
