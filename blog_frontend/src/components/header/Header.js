import React from "react";
import logo from "../../assets/logo.svg";
import classes from "./Header.module.scss";
import Nav from "./nav/Nav";
import Signin from "../../containers/signin/Signin";
const Header = () => {
	return (
		<header className={`sub-container ${classes.header}`}>
			<nav className={classes.nav_holder}>
				<div className={classes.logo}>
					<img src={logo} alt="react baby" />
				</div>
				<Nav />
			</nav>
			<Signin />
		</header>
	);
};

export default Header;
