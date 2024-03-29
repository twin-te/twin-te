import { toast } from 'bulma-toast';
import { Dispatch, SetStateAction, useEffect, useState } from 'react';
import Modal from 'react-modal';
import { updateUserInfo } from '../api/updateUserInfo';
import styles from '../styles/components/EditUserInfoModal.module.scss';
import { User } from '../types';

type Props = {
	isOpen: boolean;
	onClose: () => void;
	setCurrentUser: Dispatch<SetStateAction<User | null | undefined>>;
	prevDisplayName: undefined | string;
	prevLink: undefined | string;
};

const EditUserInfoModal: React.FC<Props> = ({ isOpen, onClose, setCurrentUser, prevDisplayName, prevLink }) => {
	const [displayName, setDisplayName] = useState<string>(prevDisplayName || '');
	const [link, setLink] = useState<string>(prevLink || '');
	const [error, setError] = useState({ displayName: '', link: '' });

	useEffect(() => {
		if (!isValidDisplayName(displayName)) {
			setError({ ...error, displayName: '20文字以下で入力してください' });
		} else {
			setError({ ...error, displayName: '' });
		}
	}, [displayName, error]);

	useEffect(() => {
		if (!isValidLink(link)) {
			setError({ ...error, link: 'URLの形式が間違っています' });
		} else {
			setError({ ...error, link: '' });
		}
	}, [link, error]);

	const isValidDisplayName = (name: string) => {
		return name.length < 20;
	};

	const isValidLink = (link: string) => {
		return /^https?:\/\/.+/.test(link) || link === '';
	};

	const handleUpdateClick = async () => {
		try {
			const user = await updateUserInfo(displayName, link);
			toast({
				message: '情報の更新に成功しました',
				type: 'is-success'
			});
			setCurrentUser(user);
			onClose();
		} catch (error) {
			toast({
				message: 'エラーが発生しました',
				type: 'is-danger'
			});
		}
	};

	const handleModalClose = () => {
		setDisplayName(prevDisplayName || '');
		setLink(prevLink || '');
		onClose();
	};

	return (
		<Modal isOpen={isOpen} onRequestClose={handleModalClose} className={styles.modal} overlayClassName={styles.overlay}>
			<h1 className="title">ユーザー情報の編集</h1>
			<button className={`delete ${styles.closeButton}`} onClick={handleModalClose}></button>
			<div className="field">
				<label className="label has-text-primary">表示名</label>
				<div className="control">
					<input
						className={`input is-rounded ${isValidDisplayName(displayName) ? 'is-success' : 'is-danger'}`}
						type="text"
						placeholder="お名前・ユーザーネーム"
						value={displayName}
						onChange={(e) => setDisplayName(e.target.value)}
					/>
				</div>
				<div className="help is-danger">{error.displayName}</div>
			</div>

			<div className="field">
				<label className="label has-text-primary">リンク</label>
				<div className="control">
					<input
						className={`input is-rounded ${isValidLink(link) ? 'is-success' : 'is-danger'}`}
						type="text"
						placeholder="サイトのURL"
						value={link}
						onChange={(e) => setLink(e.target.value)}
					/>
				</div>
				<div className="help is-danger">{error.link}</div>
			</div>

			<div className="has-text-centered">
				<button
					className={`button is-primary ${styles.updateButton}`}
					onClick={handleUpdateClick}
					disabled={!(isValidDisplayName(displayName) && isValidLink(link))}
				>
					更新する
				</button>
			</div>
		</Modal>
	);
};

export default EditUserInfoModal;
