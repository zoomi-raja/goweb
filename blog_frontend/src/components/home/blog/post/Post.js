import React from "react";
import classes from "./Post.module.scss";

const Post = (props) => {
	let limit = 250;
	return (
		<div className="sub-container">
			<h3 className={`text-ac mb-1 ${classes.postHeader}`}>
				<a href="/">{props.title}</a>
			</h3>
			<p>{props.body.substring(0, limit) + "..."}</p>
			<div className={classes.postFooter}>
				<span className={classes.postFooter_date}>{props.date}</span>
				<a className={classes.postFooter_link} href="/">
					&rarr;
				</a>
			</div>
		</div>
	);
};
export default Post;
