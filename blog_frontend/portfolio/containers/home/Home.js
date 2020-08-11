import React, { useEffect, useState } from "react";
import Header from "../../components/home/header/Header";
import Intro from "../../components/home/intro/Intro";
import Intrests from "../../components/home/interests/Interests";
import Motivation from "../../components/home/motivation/Motivation";
import Posts from "../../components/home/blog/Posts";
import axios from "../../../utils/axios";

const Home = () => {
	const [posts, setPost] = useState([]);
	useEffect(() => {
		async function fetchData() {
			const postArr = await axios.get("posts");
			setPost(postArr.data);
		}
		fetchData();
		return () => {
			console.log("This will be logged on unmount");
		};
	}, []);
	return (
		<section className="section">
			<Header />
			<Intro />
			<Intrests />
			<Motivation />
			<Posts data={posts} />
		</section>
	);
};

export default Home;
