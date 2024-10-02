import type React from "react";
import { useEffect, useState } from "react";
import { useCase } from "../../usecases";
import { getLogoutUrl } from "../../utils/getAuthUrl";
import LoginModalContent from "../LoginModalContent";
import MobileHeader from "../MobileHeader";
import Sidebar from "../Sidebar";
import SweetModal from "../SweetAlert";
import styles from "./Layout.module.scss";

export const Layout: React.FC<{ children: React.ReactNode }> = ({
	children,
}) => {
	const [isLoading, setIsLoading] = useState<boolean>(true);
	const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);

	useEffect(() => {
		useCase
			.checkAuthentication()
			.then((isAuthenticated) => {
				setIsAuthenticated(isAuthenticated);
			})
			.finally(() => setIsLoading(false));
	}, []);

	const handleLogout = async () => {
		const result = await SweetModal.fire({
			title: "ログアウトしますか？",
			text: "すべてのTwin:teサービスからログアウトします",
			showCancelButton: true,
			confirmButtonText: "はい",
			cancelButtonText: "いいえ",
		});
		if (result.isConfirmed) {
			location.href = getLogoutUrl();
		}
	};

	const handleLogin = async () => {
		await SweetModal.fire({
			title: "どのアカウントでログインしますか?",
			html: LoginModalContent,
			showConfirmButton: false,
			showCancelButton: true,
			cancelButtonText: "閉じる",
		});
	};

	return (
		<>
			<div className="columns is-gapless">
				<div className="column is-hidden-tablet">
					<MobileHeader
						isLogin={isAuthenticated}
						handleLogin={handleLogin}
						handleLogout={handleLogout}
					/>
				</div>
				<div className="column is-narrow is-hidden-mobile">
					<Sidebar />
				</div>
				<div className="column">
					<section className={`section ${styles.section}`}>
						<header className="is-hidden-mobile">
							<div className="has-text-right">
								{isLoading ? (
									<button className="button is-primary is-outlined is-loading" />
								) : (
									<button
										className="button is-primary is-outlined has-text-weight-bold"
										onClick={() =>
											isAuthenticated ? handleLogout() : handleLogin()
										}
									>
										{isAuthenticated ? "ログアウト" : "ログイン"}
									</button>
								)}
							</div>
						</header>
						<main>{children}</main>
					</section>
				</div>
			</div>
		</>
	);
};
