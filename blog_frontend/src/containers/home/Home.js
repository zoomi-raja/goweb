import React from "react";
import "./Home.scss";
import Header from "../../components/home/header/Header";
import Intro from "../../components/home/intro/Intro";
import Intrests from "../../components/home/interests/Interests";
import Motivation from "../../components/home/motivation/Motivation";
import Posts from "../../components/home/blog/Posts";

const Home = () => {
	return (
		<section className="section">
			<Header />
			<Intro />
			<Intrests />
			<Motivation />
			<Posts />
		</section>
	);
};

export default Home;
