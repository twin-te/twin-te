import { ReactNode } from 'react';
import styles from './RadioButton.module.scss';

type Props = {
	defaultChecked: boolean;
	name: string;
	id: string;
	value: string;
	onChange: (inputValue: string) => void;
	children: ReactNode;
};

export const RadioButton: React.FC<Props> = ({ defaultChecked, name, id, value, onChange, children }) => {
	return (
		<div className={styles.radioLabel}>
			<input
				className={styles.radioInput}
				defaultChecked={defaultChecked}
				type="radio"
				name={name}
				id={id}
				value={value}
				onChange={(event) => onChange(event.target.value)}
			/>
			<label htmlFor={id} className={styles.radioText}>
				{children}
			</label>
		</div>
	);
};
