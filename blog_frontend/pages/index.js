import Head from "next/head";
import App from "../portfolio/App";
import { Fragment } from "react";

export default function Home() {
	return (
		<Fragment>
			<Head>
				<link rel="icon" href="/favicon.ico" />
				<meta name="viewport" content="width=device-width, initial-scale=1" />
				<meta name="author" content="Zamurd Ali"></meta>
				<meta
					name="description"
					content="Portfolio website. Blog will be fully functional soon as 2020 is all about learning, struggle and moving forward but not to give up..!"
				></meta>
				<meta property="og:title" content="Zoomi blog" />
				<meta property="og:image" content="/favicon.ico" />
				<meta
					property="og:description"
					content="Portfolio website. Blog will be fully functional soon as 2020 is all about learning, struggle and moving forward but not to give up..!"
				></meta>
				<link
					href="https://fonts.googleapis.com/css?family=Lato:100,300,400,700,900"
					rel="stylesheet"
				/>
			</Head>

			<App />
		</Fragment>
	);
}
