import React from "react";
import { BrowserRouter } from "react-router-dom";
import Layout from "../hoc/layout/Layout";
import Home from "../containers/home/Home";
//infuture add redux in this file for auth store
function App() {
	return (
		<BrowserRouter>
			<div className="container">
				<Layout>
					<Home />
				</Layout>
			</div>
		</BrowserRouter>
	);
}

export default App;
