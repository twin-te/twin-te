import dayjs from 'dayjs';
import type { NextPage } from 'next';
import { useCurrentUser } from '../hooks/useCurrentUser';
import { useLoginStatus } from '../hooks/useLoginStatus';
import { usePaymentHistory } from '../hooks/usePaymentHistory';
import { useSubscriptions } from '../hooks/useSubscriptions';
import { cancelSubscription } from '../api/stripeApi';
import { useRouter } from 'next/router';
import { PaymentTypeMap } from '../types/Payment';
import { NextSeo } from 'next-seo';
import { useState } from 'react';
import styles from '../styles/pages/MyPage.module.scss';
import EditUserInfoModal from '../components/EditUserInfoModal';
import { toast } from 'bulma-toast';
import { MdEdit } from 'react-icons/md';

const MyPage: NextPage = () => {
	const isLogin = useLoginStatus();
	const [currentUser, setCurrentUser] = useCurrentUser();
	const subscriptions = useSubscriptions();
	const paymentHistory = usePaymentHistory();
	const router = useRouter();

	const [isEditUserModalOpen, setIsEditUserModalOpen] = useState<boolean>(false);

	const handleClick = async (id: string) => {
		try {
			await cancelSubscription(id);
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

	if (isLogin === undefined || currentUser === undefined || subscriptions === undefined || paymentHistory === undefined)
		return <div>loading...</div>;

	return (
		<>
			<NextSeo title="マイページ" />
			<div className={styles.content}>
				<h1 className="title pagetitle">マイページ</h1>
				{isLogin ? (
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
								setCurrentUser={setCurrentUser}
								prevDisplayName={currentUser?.displayName}
								prevLink={currentUser?.link}
							/>
							<div className="content">
								<p>
									<a href="https://www.twinte.net/sponsor">寄付者一覧</a>
									に表示するお名前とリンクです。
								</p>
								{currentUser ? (
									<>
										<p className="has-text-primary has-text-weight-bold is-marginless">ID</p>
										<p>{currentUser.twinteUserId}</p>

										<p className="has-text-primary has-text-weight-bold is-marginless">表示名</p>
										<p>{currentUser.displayName || '未設定'}</p>

										<p className="has-text-primary has-text-weight-bold is-marginless">リンク</p>
										<p>{currentUser.link || '未設定'}</p>
									</>
								) : (
									<div>情報の取得に失敗しました。</div>
								)}
							</div>
						</div>

						<div className="card">
							<h2 className={`title ${styles.title}`}>サブスクリプションの登録状況</h2>
							<div className="content">
								<p className="has-text-primary has-text-weight-bold">ご利用中のプラン</p>
								{subscriptions != null ? (
									subscriptions.length ? (
										<table>
											<thead>
												<tr>
													<th>プラン</th>
													<th>登録日</th>
													<th>解約</th>
												</tr>
											</thead>
											<tbody>
												{subscriptions
													.filter((subscription) => subscription.status === 'Active')
													.map((subscription) => (
														<tr key={subscription.id}>
															<td>{subscription.plans[0].name}</td>
															<td>{dayjs(subscription.created).format('YYYY.MM.DD')}</td>
															<td>
																<button
																	className="button is-danger is-small is-inverted"
																	onClick={() => handleClick(subscription.id)}
																>
																	解約
																</button>
															</td>
														</tr>
													))}
											</tbody>
										</table>
									) : (
										<div>ご利用中のプランはありません。</div>
									)
								) : (
									<div>情報の取得に失敗しました。</div>
								)}
							</div>
						</div>

						<div className="card">
							<h2 className={`title ${styles.title}`}>寄付の履歴</h2>
							<div className="content">
								{paymentHistory != null ? (
									paymentHistory.length ? (
										<table>
											<thead>
												<tr>
													<th>日付</th>
													<th>金額</th>
													<th>種別</th>
												</tr>
											</thead>
											<tbody>
												{paymentHistory
													.filter((payment) => payment.status === 'Succeeded')
													.map((payment) => (
														<tr key={payment.id}>
															<td>{dayjs(payment.created).format('YYYY.MM.DD')}</td>
															<td>
																<p>{payment.amount}円</p>
															</td>
															<td>
																<p className="has-text-grey">{PaymentTypeMap[payment.type]}</p>
															</td>
														</tr>
													))}
											</tbody>
										</table>
									) : (
										<div>寄付の履歴はありません。</div>
									)
								) : (
									<div>情報の取得に失敗しました。</div>
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
