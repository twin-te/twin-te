import Sidebar from './Sidebar';
import MobileHeader from './MobileHeader';
import { useLoginStatus } from '../hooks/useLoginStatus';
import LoginModalContent from './LoginModalContent';
import { SweetModal } from './SweetAlert';
import { getLogoutUrl } from '../utils/getAuthUrl';
import Image from 'next/image';
import styles from '../styles/components/Layout.module.scss';

export const Layout: React.FC = ({ children }) => {
	const isLogin = useLoginStatus();

	const handleLogout = async () => {
		const result = await SweetModal.fire({
			title: 'ログアウトしますか？',
			text: 'すべてのTwin:teサービスからログアウトします',
			showCancelButton: true,
			confirmButtonText: 'はい',
			cancelButtonText: 'いいえ'
		});
		if (result.isConfirmed) {
			window.open(getLogoutUrl());
		}
	};

	const handleLogin = async () => {
		await SweetModal.fire({
			title: 'どのアカウントでログインしますか?',
			html: LoginModalContent,
			showConfirmButton: false,
			showCancelButton: true,
			cancelButtonText: '閉じる'
		});
	};

	return (
		<>
			<div className="columns is-gapless">
				<div className="column is-hidden-tablet">
					<MobileHeader isLogin={isLogin} handleLogin={handleLogin} handleLogout={handleLogout} />
				</div>
				<div className="column is-narrow is-hidden-mobile">
					<Sidebar />
				</div>
				<div className="column">
					<section className={`section ${styles.section}`}>
						<header className="is-hidden-mobile">
							<div className="has-text-right">
								{isLogin == undefined ? (
									<button className="button is-primary is-outlined is-loading" />
								) : (
									<button
										className="button is-primary is-outlined has-text-weight-bold"
										onClick={() => (isLogin ? handleLogout() : handleLogin())}
									>
										{isLogin ? 'ログアウト' : 'ログイン'}
									</button>
								)}
							</div>
						</header>
						<main>{children}</main>
						<footer className="has-text-centered-mobile has-text-right-tablet mt-6">
							<a href="https://vercel.com?utm_source=twin-te&utm_campaign=oss">
								<Image
									src="https://www.datocms-assets.com/31049/1618983297-powered-by-vercel.svg"
									width={212}
									height={44}
									alt="Vercel"
								/>
							</a>
						</footer>
					</section>
				</div>
			</div>
		</>
	);
};

export default Layout;
