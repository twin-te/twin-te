import styles from './Sidebar.module.scss';
import NavLink from '../NavLink';
import Image from 'next/image';
import TwinteLogo from '@/public/images/twinte-sponsor-title.png';
import { ENV_NEXT_PUBLIC_APP_URL } from '@/env';

export const Sidebar: React.FC = () => {
	return (
		<aside className={styles.sidebar}>
			<div className={styles.menuContent}>
				<div className={styles.logo}>
					<a href={ENV_NEXT_PUBLIC_APP_URL}>
						<Image src={TwinteLogo} alt="Twin:te_Logo" />
					</a>
				</div>
				<ul className="menu-list">
					<li>
						<NavLink href="/" activeClassName={styles.activeLink}>
							寄付のお願い
						</NavLink>
					</li>
					<li>
						<NavLink href="/register" activeClassName={styles.activeLink}>
							寄付・サブスク登録
						</NavLink>
					</li>
					<li>
						<NavLink href="/mypage" activeClassName={styles.activeLink}>
							マイページ
						</NavLink>
					</li>
					<li>
						<NavLink href="/gift" activeClassName={styles.activeLink}>
							返礼品一覧
						</NavLink>
					</li>
				</ul>
			</div>
			<div className={`${styles.menuFooter} has-text-white`}>
				<p>
					<span className="has-text-weight-bold">WEB SITE</span>
					<br />
					<a href="https://www.twinte.net">www.twinte.net</a>
				</p>
				<p>
					<span className="has-text-weight-bold">Twitter</span>
					<br />
					<a href="https://twitter.com/te_twin">@te_twin</a>
				</p>
				<p>
					<span className="has-text-weight-bold">E-mail</span>
					<br />
					<a href="mailto:info@twinte.net?subject=お問い合わせ">info@twinte.net</a>
				</p>
			</div>
		</aside>
	);
};
