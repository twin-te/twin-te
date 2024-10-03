import type { User } from "@/domain";
import { useCase } from "@/usecases";
import { toast } from "bulma-toast";
import { type Dispatch, type SetStateAction, useEffect, useState } from "react";
import Modal from "react-modal";
import styles from "./EditUserInfoModal.module.scss";

type Props = {
	isOpen: boolean;
	onClose: () => void;
	setCurrentUser: Dispatch<SetStateAction<User | undefined>>;
	prevDisplayName: undefined | string;
	prevLink: undefined | string;
};

export const EditUserInfoModal: React.FC<Props> = ({
	isOpen,
	onClose,
	setCurrentUser,
	prevDisplayName,
	prevLink,
}) => {
	const [displayName, setDisplayName] = useState<string>(prevDisplayName || "");
	const [link, setLink] = useState<string>(prevLink || "");
	const [error, setError] = useState({ displayName: "", link: "" });

	useEffect(() => {
		if (!isValidDisplayName(displayName)) {
			setError({ ...error, displayName: "20文字以下で入力してください" });
		} else {
			setError({ ...error, displayName: "" });
		}
	}, [displayName]);

	useEffect(() => {
		if (!isValidLink(link)) {
			setError({ ...error, link: "URLの形式が間違っています" });
		} else {
			setError({ ...error, link: "" });
		}
	}, [link]);

	const isValidDisplayName = (name: string) => {
		return name.length < 20;
	};

	const isValidLink = (link: string) => {
		return /^https?:\/\/.+/.test(link) || link === "";
	};

	const handleUpdateClick = async () => {
		try {
			const user = await useCase.updateUserInfo(
				displayName === "" ? undefined : displayName,
				link === "" ? undefined : link,
			);
			toast({
				message: "情報の更新に成功しました",
				type: "is-success",
			});
			setCurrentUser(user);
			onClose();
		} catch {
			toast({
				message: "エラーが発生しました",
				type: "is-danger",
			});
		}
	};

	const handleModalClose = () => {
		setDisplayName(prevDisplayName || "");
		setLink(prevLink || "");
		onClose();
	};

	return (
		<Modal
			isOpen={isOpen}
			onRequestClose={handleModalClose}
			className={styles.modal}
			overlayClassName={styles.overlay}
		>
			<h1 className="title">ユーザー情報の編集</h1>
			<button
				className={`delete ${styles.closeButton}`}
				onClick={handleModalClose}
			/>
			<div className="field">
				<label className="label has-text-primary">表示名</label>
				<div className="control">
					<input
						className={`input is-rounded ${isValidDisplayName(displayName) ? "is-success" : "is-danger"}`}
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
						className={`input is-rounded ${isValidLink(link) ? "is-success" : "is-danger"}`}
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
