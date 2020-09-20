import React from "react";
import Link from "next/link";
import classes from "./Post.module.scss";

const Post = (props) => {
	let limit = 250;
	return (
		<div className={`sub-container ${classes.post}`}>
			<h3 className={`text-ac mb-1 ${classes.postHeader}`}>
				<Link href="/post/[pid]" as={`/post/${props.id}`}>
					<a>{props.title}</a>
				</Link>
			</h3>
			<p>{props.intro.substring(0, limit) + "..."}</p>
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
