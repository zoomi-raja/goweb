import React from "react";
import classes from "./Posts.module.scss";
import Post from "./post/Post";
const posts = [
	{
		id: 1,
		title: "node",
		body:
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur",
		date: "Aug 08 2020",
	},
	{
		id: 2,
		title: "GoLang",
		body:
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur",
		date: "Aug 04 2020",
	},
	{
		id: 3,
		title: "git",
		body:
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur",
		date: "Aug 01 2020",
	},
];
const Posts = () => {
	let html = posts.map((post, i) => {
		return <Post {...post} key={i} />;
	});
	return (
		<div className={classes.articles}>
			<h1 className="text-ac mb-2">
				<span className="text-ul">Articles</span>
			</h1>
			<div className={classes.posts}>{html}</div>
		</div>
	);
};
export default Posts;
