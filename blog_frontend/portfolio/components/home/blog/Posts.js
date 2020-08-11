import React from "react";
import classes from "./Posts.module.scss";
import Post from "./post/Post";
const Posts = (props) => {
	let html;
	if (props.data) {
		html = props.data.map((post, i) => {
			return <Post {...post} key={i} />;
		});
	}
	return (
		<div className={classes.articles} id="articles">
			<h1 className="text-ac mb-2">
				<span className="text-ul">Articles</span>
			</h1>
			<div className={classes.posts}>{html}</div>
		</div>
	);
};
export default Posts;
