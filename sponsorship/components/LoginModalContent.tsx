import AppleLogin from "@/public/images/auth/login-apple.svg";
import GoogleLogin from "@/public/images/auth/login-google.svg";
import TwitterLogin from "@/public/images/auth/login-twitter.svg";
import Image from "next/image";
import { getLoginUrl } from "../utils/getAuthUrl";

const LoginModalContent = (
	<>
		<p>Twin:teアプリと同じアカウントでログインすることができます。</p>
		<div className="has-text-centered">
			<div className="m-3">
				<a href={getLoginUrl("apple")}>
					<Image src={AppleLogin} alt="sign in with apple" />
				</a>
			</div>
			<div className="m-3">
				<a href={getLoginUrl("twitter")}>
					<Image src={TwitterLogin} alt="sign in with apple" />
				</a>
			</div>
			<div className="m-3">
				<a href={getLoginUrl("google")}>
					<Image src={GoogleLogin} alt="sign in with apple" />
				</a>
			</div>
		</div>
	</>
);

export default LoginModalContent;
