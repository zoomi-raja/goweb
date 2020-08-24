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
			// const postArr = await axios.get("posts");
			let postArr = [
				{
					id: 1,
					userId: 0,
					claps: 0,
					intro:
						"A function's this keyword behaves a little differently in JavaScript compared to other languages. It also has some differences between strict mode and non-strict mode.  In most cases, the value of this is determined by how a function is called (runti",
					title: "JavaScript `this`",
					body: "",
					createdAt: "22-Aug-2020 1:24 PM",
					updatedAt: "22-Aug-2020 1:24 PM",
				},
				{
					id: 2,
					userId: 0,
					claps: 0,
					intro:
						"text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five ce",
					title: "golang",
					body: "",
					createdAt: "10-Aug-2020 5:34 PM",
					updatedAt: "10-Aug-2020 5:34 PM",
				},
				{
					id: 3,
					userId: 0,
					claps: 0,
					intro:
						"Just to fill the item as to write whole blog post is a bit challenging and top of that my markup to html generator is not that much effective. It will improve by time but the first priority is to make the portfolio website available.",
					title: "Readme file",
					body: "",
					createdAt: "21-Aug-2020 6:03 PM",
					updatedAt: "21-Aug-2020 6:03 PM",
				},
			];
			setPost(postArr);
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
