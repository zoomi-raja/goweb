import React from "react";
import "./Home.scss";
import Header from "../../components/home/header/Header";
import Intro from "../../components/home/intro/Intro";
import Intrests from "../../components/home/interests/Interests";

const Home = () => {
	return (
		<section className="section">
			<Header />
			<Intro />
			<Intrests />
			<div className="sub-container">motivation</div>
			<div className="sub-container">portfolio</div>
			<div className="sub-container">articles</div>
		</section>
	);
};

export default Home;
