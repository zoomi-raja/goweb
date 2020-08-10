import React, { Fragment } from "react";
import Layout from "./hoc/layout/Layout";
import Home from "./containers/home/Home";
import Head from "next/head";
//infuture add redux in this file for auth store
function App() {
	return (
		<Fragment>
			<Head>
				<title>zoomi blog</title>
			</Head>
			<div className="container">
				<Layout>
					<Home />
				</Layout>
			</div>
		</Fragment>
	);
}

export default App;
