import type { NextPage } from 'next';
import { useRouter } from 'next/router';
import { PaymentHistory, PaymentTypeMap } from '../domain/payment_history';
import { NextSeo } from 'next-seo';
import { useEffect, useState } from 'react';
import styles from '../styles/pages/MyPage.module.scss';
import { toast } from 'bulma-toast';
import { MdEdit } from 'react-icons/md';
import { useCase } from '@/usecases';
import { Subscription, User } from '@/domain';
import { isNotFoundError, isUnauthenticatedError } from '@/usecases/error';
import EditUserInfoModal from '@/components/EditUserInfoModal';

const MyPage: NextPage = () => {
	const router = useRouter();

	const handleClick = async (id: string) => {
		try {
			await useCase.cancelSubscription(id);
			toast({
				message: '解約に成功しました',
				type: 'is-success'
			});
			await new Promise((s) => setTimeout(s, 2000));
			router.reload();
		} catch (error) {
			console.error(error);
			toast({
				message: 'エラーが発生しました',
				type: 'is-danger'
			});
		}
	};

	const [isLoading, setIsLoading] = useState<boolean>(true);
	const [user, setUser] = useState<User | undefined>(undefined);
	const [activeSubscription, setActiveSubscription] = useState<Subscription | undefined>(undefined);
	const [paymentHistories, setPaymentHistories] = useState<PaymentHistory[]>([]);
	const [isEditUserModalOpen, setIsEditUserModalOpen] = useState<boolean>(false);

	useEffect(() => {
		useCase
			.getUser()
			.then((user) => {
				setUser(user);
			})
			.then(() => {
				Promise.all([
					useCase
						.getActiveSubscription()
						.then((activeSubscription) => {
							setActiveSubscription(activeSubscription);
						})
						.catch((error) => {
							if (isNotFoundError(error)) return;
							throw error;
						}),
					useCase.getPaymentHistories().then((paymentHistories) => {
						setPaymentHistories(paymentHistories);
					}),
				]);
			})
			.catch((error) => {
				if (isUnauthenticatedError(error)) return;
				throw error;
			})
			.finally(() => setIsLoading(false));
	}, []);

	if (isLoading) {
		return <div>now loading...</div>;
	}

	return (
		<>
			<NextSeo title="マイページ" />
			<div className={styles.content}>
				<h1 className="title pagetitle">マイページ</h1>
				{user ? (
					<>
						<div className="card">
							<h2 className={`title ${styles.title}`}>ユーザー情報</h2>
							<button className={`button is-text ${styles.editButton}`} onClick={() => setIsEditUserModalOpen(true)}>
								<span className={styles.editIcon}>
									<MdEdit size="1.5rem" color="#929292" />
								</span>
								編集
							</button>
							<EditUserInfoModal
								isOpen={isEditUserModalOpen}
								onClose={() => setIsEditUserModalOpen(false)}
								setCurrentUser={setUser}
								prevDisplayName={user.displayName}
								prevLink={user.link}
							/>
							<div className="content">
								<p>
									<a href="https://www.twinte.net/sponsor">寄付者一覧</a>
									に表示するお名前とリンクです。
								</p>
<<<<<<< HEAD
								<p className="has-text-primary has-text-weight-bold is-marginless">
									ID
								</p>
								<p>{user.twinteUserId}</p>
								<p className="has-text-primary has-text-weight-bold is-marginless">
									表示名
								</p>
								<p>{user.displayName || "未設定"}</p>
								<p className="has-text-primary has-text-weight-bold is-marginless">
									リンク
								</p>
								<p>{user.link || "未設定"}</p>
=======
								{user ? (
									<>
										<p className="has-text-primary has-text-weight-bold is-marginless">ID</p>
										<p>{user.twinteUserId}</p>

										<p className="has-text-primary has-text-weight-bold is-marginless">表示名</p>
										<p>{user.displayName || '未設定'}</p>

										<p className="has-text-primary has-text-weight-bold is-marginless">リンク</p>
										<p>{user.link || '未設定'}</p>
									</>
								) : (
									<div>情報の取得に失敗しました。</div>
								)}
>>>>>>> parent of 5744088 ([sponsorship] biome (#148))
							</div>
						</div>
						<div className="card">
							<h2 className={`title ${styles.title}`}>サブスクリプションの登録状況</h2>
							<div className="content">
								<p className="has-text-primary has-text-weight-bold">ご利用中のプラン</p>
								{activeSubscription ? (
									<table>
										<thead>
											<tr>
												<th>プラン</th>
												<th>登録日</th>
												<th>解約</th>
											</tr>
										</thead>
										<tbody>
											<tr key={activeSubscription.id}>
												<td>{activeSubscription.plan.name}</td>
												<td>{activeSubscription.createdAt.format('YYYY.MM.DD')}</td>
												<td>
													<button
														className="button is-danger is-small is-inverted"
														onClick={() => handleClick(activeSubscription.id)}
													>
														解約
													</button>
												</td>
											</tr>
										</tbody>
									</table>
								) : (
									<div>ご利用中のプランはありません。</div>
								)}
							</div>
						</div>
						<div className="card">
							<h2 className={`title ${styles.title}`}>寄付の履歴</h2>
							<div className="content">
								{paymentHistories.length !== 0 ? (
									<table>
										<thead>
											<tr>
												<th>日付</th>
												<th>金額</th>
												<th>種別</th>
											</tr>
										</thead>
										<tbody>
											{paymentHistories
												.filter((payment) => payment.status === 'Succeeded')
												.map((paymentHistory) => (
													<tr key={paymentHistory.id}>
														<td>{paymentHistory.createdAt.format('YYYY-MM-DD')}</td>
														<td>
															<p>{paymentHistory.amount}円</p>
														</td>
														<td>
															<p className="has-text-grey">{PaymentTypeMap[paymentHistory.type]}</p>
														</td>
													</tr>
												))}
										</tbody>
									</table>
								) : (
									<div>寄付の履歴はありません。</div>
								)}
							</div>
						</div>
					</>
				) : (
					<p>
						右上の「<span className="has-text-weight-bold">ログイン</span>」ボタンからログインしてください。
					</p>
				)}
			</div>
		</>
	);
};

export default MyPage;
