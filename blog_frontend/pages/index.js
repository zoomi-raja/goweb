import Head from "next/head";
import App from "../portfolio/App";
import { Fragment } from "react";

export default function Home() {
	return (
		<Fragment>
			<Head>
				<link rel="icon" href="/favicon.ico" />
				<meta name="viewport" content="width=device-width, initial-scale=1" />
				<link
					href="https://fonts.googleapis.com/css?family=Lato:100,300,400,700,900"
					rel="stylesheet"
				/>
			</Head>

			<App />
		</Fragment>
	);
}
