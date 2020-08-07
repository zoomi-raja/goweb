import React from "react";
import "./Home.scss";
import Header from "../../components/home/header/Header";

const Home = () => {
	return (
		<section className="section">
			<Header />
			<div className="sub-container">intro</div>
			<div className="sub-container">intrests</div>
			<div className="sub-container">motivation</div>
			<div className="sub-container">portfolio</div>
			<div className="sub-container">articles</div>
		</section>
	);
};

export default Home;
