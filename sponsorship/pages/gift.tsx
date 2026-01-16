import type { NextPage } from 'next';
import Image from 'next/image';
import styles from '../styles/pages/Gift.module.scss';
import TwinteSticker from '@/public/images/twinte-sticker.jpg';
import TwinteUSB from '@/public/images/twinte-usb.jpg';
import { NextSeo } from 'next-seo';

const Gift: NextPage = () => {
	return (
		<>
			<NextSeo title="返礼品一覧" />
			<div className={styles.content}>
				<div className={styles.description}>
					<h1 className="title pagetitle">返礼品一覧</h1>
					<p>
						寄付合計金額が以下に記載の値段を超える方はそれぞれの返礼品を申し込むことができます。
						<br />
						返礼品をご希望の場合は、マイページに記載のID・合計金額・希望する返礼品を記載の上
						<a href="mailto:info@twinte.net?subject=返礼品について">info@twinte.net</a>までご連絡ください。
					</p>
				</div>
				<div className="card">
					<div>
						<h1 className="title">3000円 Twin:teオリジナルステッカー</h1>
					</div>
					<div className={styles.imageWrapper}>
						<Image objectFit="contain" src={TwinteSticker} alt="Twin:teステッカー" />
					</div>
					<p>Twin:teのロゴのステッカーです！シールになっているので好きなところに貼れますよ！</p>
				</div>
				<div className="card">
					<div>
						<h1 className="title">10000円 Twin:teオリジナルUSB</h1>
					</div>

					<div className={styles.imageWrapper}>
						<Image objectFit="contain" src={TwinteUSB} alt="Twin:teステッカー" />
					</div>
					<p>
						オリジナルキャラクター「コロン」とロゴが描かれたオリジナルUSBです！
						<br />
						裏面に描かれるコロンは三種類から選択することができます！
						<br />
						メールを送る際にどのコロンか選択してくださいね♪
						<br />
						大容量・高速なので、普段使いにも良しOSインストーラーとしても良し、キーホルダーにしても良しです！
						<br />
						仕様：USB3.0 32GB
						<br />
						※使用したことによって起こった損害に関してTwin:teは一切の責任を負いません。
					</p>
				</div>
			</div>
		</>
	);
};

export default Gift;
