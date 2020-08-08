import React from "react";
import classes from "./Footer.module.scss";
import Social from "../social/Social";

const Footer = () => {
	return (
		<div className={`sub-container ${classes.footer}`}>
			<span> Built by zoomi, &copy; 2020, All Rights Reserved</span>
			<Social />
		</div>
	);
};

export default Footer;
