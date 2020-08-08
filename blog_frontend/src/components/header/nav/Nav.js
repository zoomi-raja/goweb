import React from "react";
import classes from "./Nav.module.scss";
import { Link } from "react-scroll";
const Nav = () => {
	return (
		<ul className={classes.main_nav}>
			<li>
				<Link to="intro" spy={true} smooth={true} duration={300} offset={-60}>
					Introduction
				</Link>
			</li>
			<li>
				<Link
					to="interests"
					spy={true}
					smooth={true}
					duration={300}
					offset={-60}
				>
					Intrests
				</Link>
			</li>
			<li>
				<Link
					to="motivation"
					spy={true}
					smooth={true}
					duration={300}
					offset={-60}
				>
					Motivation
				</Link>
			</li>
			<li>
				<Link
					to="articles"
					spy={true}
					smooth={true}
					duration={300}
					offset={-60}
				>
					Articles
				</Link>
			</li>
		</ul>
	);
};

export default Nav;
