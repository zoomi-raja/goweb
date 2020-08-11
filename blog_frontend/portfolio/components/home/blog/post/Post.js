import React from "react";
import Link from "next/link";
import classes from "./Post.module.scss";

const Post = (props) => {
	let limit = 250;
	return (
		<div className={`sub-container ${classes.post}`}>
			<h3 className={`text-ac mb-1 ${classes.postHeader}`}>
				<a href="/">{props.title}</a>
			</h3>
			<p>{props.body.substring(0, limit) + "..."}</p>
			<div className={classes.postFooter}>
				<span className={classes.postFooter_date}>{props.createdAt}</span>
				<Link href="/post/[pid]" as={`/post/${props.id}`}>
					<a className={classes.postFooter_link}>&rarr;</a>
				</Link>
			</div>
		</div>
	);
};
export default Post;
