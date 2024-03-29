import { useRouter } from 'next/router';
import Link from 'next/link';
import { ReactNode } from 'react';

type Props = {
	href: string;
	activeClassName?: string;
	children: ReactNode;
};

const NavLink = ({ href, activeClassName, children }: Props) => {
	const router = useRouter();

	return (
		<Link href={href}>
			<a className={router.pathname === href ? activeClassName : undefined}>{children}</a>
		</Link>
	);
};

export default NavLink;
