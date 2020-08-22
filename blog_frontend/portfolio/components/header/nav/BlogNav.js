import React from "react";
import Link from "next/link";
import classes from "./Nav.module.scss";

const BlogNav = () => {
	return (
		<ul className={classes.main_nav}>
			<li className="navItem">
				<Link href="/">
					<a>Home</a>
				</Link>
			</li>
		</ul>
	);
};

export default BlogNav;
