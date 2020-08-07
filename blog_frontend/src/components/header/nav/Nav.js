import React from "react";
import classes from "./Nav.module.scss";
const Nav = () => {
	return (
		<ul className={classes.main_nav}>
			<li>
				<a href="#">Introduction</a>
			</li>
			<li className={classes.active}>
				<a href="#">Intrests</a>
			</li>
			<li>
				<a href="#">Motivation</a>
			</li>
			<li>
				<a href="#">Portfolio</a>
			</li>
			<li>
				<a href="#">Articles</a>
			</li>
		</ul>
	);
};

export default Nav;
