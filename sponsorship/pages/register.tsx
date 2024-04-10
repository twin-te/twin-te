import { useState } from 'react';
import type { NextPage } from 'next';
import { useLoginStatus } from '../hooks/useLoginStatus';
import Slider from 'react-input-slider';
import styles from '../styles/pages/Register.module.scss';
import { redirectToCheckout, subscriptions } from '../stripe';
import { NextSeo } from 'next-seo';
import { SweetModal } from '../components/SweetAlert';
import { RadioButton } from '../components/RadioButton';
import Link from 'next/link';
import { useCase } from '../usecases';
import { usePlans } from '../hooks/usePlans';

const ACTUAL_RECEIVED_PERCENTAGE = 0.964;
const MONTHLY_COST = 7052; // ref: /public/images/twinte-cost.png

const Register: NextPage = () => {
	const isLogin = useLoginStatus();
	const [donationPriceIndex, setDonationPriceIndex] = useState(0);
	const donationPrices = [500, 1000, 1500, 2000, 3000, 5000, 7000, 10000];
	const [subscriptionID, setSubscriptionID] = useState(subscriptions[0].planId);

	const { plans, selectPlanID} = usePlans()

	const confirmmakeOneTimeDonation = async () => {
		await SweetModal.fire({
			title: 'ログインをしてください。',
			text: '寄付をするには、右上のログインボタンよりログインをしてください。',
			showCancelButton: false,
			confirmButtonText: 'はい'
		});
	};
	const confirmregisterSubscription = async () => {
		await SweetModal.fire({
			title: 'ログインをしてください。',
			text: '継続的な寄付をするには、右上のログインボタンよりログインをしてください。',
			showCancelButton: false,
			confirmButtonText: 'はい'
		});
	};

	const radioButtons = () => {
		return subscriptions.map((plan, index) => {
			return (
				<div key={index} className="field has-text-weight-bold">
					<RadioButton
						defaultChecked={index === 0}
						name="priceChoice"
						id={`plan_${index}`}
						value={plan.planId}
						onChange={(inputValue) => {
							setSubscriptionID(inputValue);
						}}
					>
						{plan.amount}円/月
					</RadioButton>
				</div>
			);
		});
	};

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
						isLogin
							? useCase.makeOneTimeDonation(donationPrices[donationPriceIndex]).then(redirectToCheckout)
							: confirmmakeOneTimeDonation();
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
						isLogin
							? useCase.registerSubscription(subscriptionID).then(redirectToCheckout)
							: confirmregisterSubscription();
					}}
				>
					登録する
				</button>
			</div>
		</>
	);
};

export default Register;
