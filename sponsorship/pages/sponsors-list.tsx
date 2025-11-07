import type { NextPage } from 'next';
import styles from '../styles/pages/Home.module.scss';
import { NextSeo } from 'next-seo';
import { useEffect, useState } from 'react';
import { useCase } from '@/usecases';
import { Contributor } from '@/domain';

const Home: NextPage = () => {
	const [contributors, setContributors] = useState<Contributor[]>([]);

	useEffect(() => {
		useCase.listContributors().then((contributors) => {
			setContributors(contributors);
		});
	}, []);

	return (
		<>
			<NextSeo />

			<div className={styles.content}>
				<h1 className="title pagetitle">寄付者一覧</h1>

				<p>Twin:teに寄付をしていただき、許可いただいた方を掲載いたします。（敬称略）</p>
				<p>
					なお、収支報告書を
					<a
						href="https://drive.google.com/drive/folders/1nHj8w5LELC5ZTFnWvgF7HXPriMRovFSp?usp=sharing"
						target="_blank"
					>
						こちら
					</a>
					にて公開しております。
				</p>

				<ul>
					{contributors.map((name, index) => (
						<li key={index}>
							{name.link === undefined ? (
								name.displayName
							) : (
								<a href={name.link} target="_blank" rel="nofollow noopener noreferrer" referrerPolicy="no-referrer">
									{name.displayName}
								</a>
							)}
						</li>
					))}
				</ul>
			</div>
		</>
	);
};

export default Home;
